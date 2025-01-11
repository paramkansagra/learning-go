package authentication

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

/*
	We have a EC private key generated using this command -> openssl ecparam -name prime256v1 -genkey -noout -out private-key.pem

	Now we will load this Private key using the loadECDSAPrivateKeyFromPEM where we are following these steps:-
		1. Load the file using os.ReadFile()
			1.1. Then if there is any error we will return the error
		2. Once we get the byte[] data we will try to decode it using pem.Decode()
			2.1. If there is any error in decoding the block we will throw that error
		3. Then we will check weather it is a private key or not
			3.1. If it is not a private key we wll return Invalid PEM block type
		4. Now we will try to decode using the PKCS#8 & SEC1 format
			4.1 If it is none of these formats we will throw an error saying the pem key is not supported

	After all of this we will get the ecdsa.PrivateKey for signing our JWT tokens
		We need to sign our JWT tokens so that we can verify that it is coming from us
*/

var privateSigningKey *ecdsa.PrivateKey
var publicSigningKey *ecdsa.PublicKey

/*
In this init function we want to initialize some global variables used in jwt process like
 1. Public Signing Key
 2. Private Signing Key

Now we are using the Private Signing key for creating new JWT tokens
And we are using the Public Signing key for veriftying the incoming JWT tokens weather they are valid or not
*/
func init() {
	var err error
	privateSigningKey, publicSigningKey, err = loadECDSAPrivateKeyFromPEM()

	if err != nil {
		panic(err)
	}
}

func loadECDSAPrivateKeyFromPEM() (*ecdsa.PrivateKey, *ecdsa.PublicKey, error) {

	var filepath string = "./private-key.pem"

	// Read the PEM file
	pemData, err := os.ReadFile(filepath)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read private key file: %w", err)
	}

	// Decode PEM block
	block, _ := pem.Decode(pemData)
	if block == nil {
		return nil, nil, fmt.Errorf("failed to decode PEM block")
	}

	// Check if it's a private key
	if block.Type != "EC PRIVATE KEY" && block.Type != "PRIVATE KEY" {
		return nil, nil, fmt.Errorf("invalid PEM block type: %s", block.Type)
	}

	// Parse the private key
	var privateKey *ecdsa.PrivateKey

	// First try PKCS#8 format
	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err == nil {
		var ok bool
		privateKey, ok = key.(*ecdsa.PrivateKey)
		if !ok {
			return nil, nil, fmt.Errorf("not an ECDSA private key")
		}
		return privateKey, &privateKey.PublicKey, nil
	}

	// If PKCS#8 fails, try SEC1 format
	privateKey, err = x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		// If both parsing attempts fail, return detailed error
		return nil, nil, fmt.Errorf("failed to parse private key (tried PKCS#8 and SEC1 formats). Make sure the key uses a supported curve (P-224, P-256, P-384, P-521)")
	}

	// Verify the curve is supported
	curve := privateKey.Curve
	switch curve {
	case elliptic.P224(), elliptic.P256(), elliptic.P384(), elliptic.P521():
		// These curves are supported
	default:
		return nil, nil, fmt.Errorf("unsupported elliptic curve: %T", curve)
	}

	return privateKey, &privateKey.PublicKey, nil
}

/*
To create a JWT token we need to have the username for which we are going to generate the token

Now before creating the token we can do checks like:-
 1. Check the username password
 2. If there is already a jwt token present then revoke it and all

# After checking we can create this token for them

Here we are creating a new token with claims where we have
 1. Username of the person
 2. Issuer of the jwt token
 3. Expiration time of the jwt token
 4. Issuing time of the jwt token
*/
func CreateToken(username string) (string, error) {
	claims := jwt.NewWithClaims(
		jwt.SigningMethodES256, // issuing method
		jwt.MapClaims{
			"user":            username,                         // user
			"issuer":          "basic-jwt-app",                  // Issuer
			"expiration_time": time.Now().Add(time.Hour).Unix(), // expiration time
			"issuing_time":    time.Now().Unix(),                // Issuing time
		},
	)

	// after getting the signed jwtTokenString we will check for the error and return
	jwtTokenString, err := claims.SignedString(privateSigningKey)

	// checking if there is any error in signing the jwtToken
	if err != nil {
		return "", err
	}

	// return the jwtTokenString
	return jwtTokenString, nil
}

func VerifyToken(jwtToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(jwtToken, func(*jwt.Token) (interface{}, error) { return publicSigningKey, nil })

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return token, nil
}
