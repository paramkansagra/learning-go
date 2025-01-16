package authentication

import (
	"auth-with-db-jwt/models"
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

var privateSigningKey *ecdsa.PrivateKey
var publicSigningKey *ecdsa.PublicKey

func init() {
	var err error

	privateSigningKey, publicSigningKey, err = getPublicPrivateSigningKey("./private-key.pem")

	if err != nil {
		panic(err)
	}
}

func getPublicPrivateSigningKey(filepath string) (*ecdsa.PrivateKey, *ecdsa.PublicKey, error) {
	// first we will try to read the path
	fileBytes, err := os.ReadFile(filepath)

	if err != nil {
		return nil, nil, err
	}

	// now we will decode the data to pem block
	pemData, _ := pem.Decode(fileBytes)

	if pemData == nil {
		return nil, nil, errors.New("failed to load pem block from the file")
	}

	// then we will check the PEM data headers
	if pemData.Type != "PRIVATE KEY" && pemData.Type != "EC PRIVATE KEY" {
		return nil, nil, errors.New("key not private key / ec private key")
	}

	// First try PKCS#8 format
	key, err := x509.ParsePKCS8PrivateKey(pemData.Bytes)
	if err == nil {
		var ok bool
		privateSigningKey, ok = key.(*ecdsa.PrivateKey)
		if !ok {
			return nil, nil, fmt.Errorf("not an ECDSA private key")
		}
		return privateSigningKey, &privateSigningKey.PublicKey, nil
	}

	// If PKCS#8 fails, try SEC1 format
	privateSigningKey, err = x509.ParseECPrivateKey(pemData.Bytes)
	if err != nil {
		// If both parsing attempts fail, return detailed error
		return nil, nil, fmt.Errorf("failed to parse private key (tried PKCS#8 and SEC1 formats). Make sure the key uses a supported curve (P-224, P-256, P-384, P-521)")
	}

	// Verify the curve is supported
	curve := privateSigningKey.Curve
	switch curve {
	case elliptic.P224(), elliptic.P256(), elliptic.P384(), elliptic.P521():
		// These curves are supported
	default:
		return nil, nil, fmt.Errorf("unsupported elliptic curve: %T", curve)
	}

	return privateSigningKey, &privateSigningKey.PublicKey, nil
}

func GetToken(user *models.User) (string, error) {
	// we will check the username and password
	// then generate JWT for the same

	jwtToken, err := generateToken(user.ID.Hex())
	return jwtToken, err
}

func VerifyToken(jwtToken string) (*jwt.Token, error) {
	jwt, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) { return publicSigningKey, nil })

	return jwt, err
}

func generateToken(username string) (string, error) {
	// we will generate

	claims := jwt.NewWithClaims(
		jwt.SigningMethodES256,
		jwt.MapClaims{
			"sub": username,                         // subject
			"iss": "auth-with-db-jwt",               // issuer
			"iat": time.Now().Unix(),                // issued at
			"exp": time.Now().Add(time.Hour).Unix(), // expires at
		},
	)

	jwtString, err := claims.SignedString(privateSigningKey)

	return jwtString, err
}
