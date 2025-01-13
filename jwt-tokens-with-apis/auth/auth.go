package auth

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

var publicSigningKey *ecdsa.PublicKey
var privateSigningKey *ecdsa.PrivateKey

func init() {
	var filepath string = "./private-key.pem"
	var err error

	publicSigningKey, privateSigningKey, err = getPublicPrivateKeyFromPEM(filepath)
	if err != nil {
		panic(err)
	}
}

func getPublicPrivateKeyFromPEM(filepath string) (*ecdsa.PublicKey, *ecdsa.PrivateKey, error) {
	if filepath == "" {
		return nil, nil, errors.New("no filepath given")
	}

	// read the file from memory
	fileData, err := os.ReadFile(filepath)

	if err != nil {
		return nil, nil, err
	}

	// decode the PEM data from it
	pemData, _ := pem.Decode(fileData)

	// check the block data to be EC type
	if pemData.Type != "EC PRIVATE KEY" && pemData.Type != "PRIVATE KEY" {
		return nil, nil, errors.New("invalid PEM block type")
	}

	// Parse the private key from PEM Data
	var privateKey *ecdsa.PrivateKey

	key, err := x509.ParsePKCS8PrivateKey(pemData.Bytes)

	if err == nil {
		privateKey, ok := key.(*ecdsa.PrivateKey)
		if !ok {
			return nil, nil, errors.New("not an ECDSA private key")
		}

		return &privateKey.PublicKey, privateKey, nil
	}

	// If PKCS#8 fails, try SEC1 format
	privateKey, err = x509.ParseECPrivateKey(pemData.Bytes)
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

	return &privateKey.PublicKey, privateKey, nil
}

func CreateToken(username string, password string) (string, error) {
	// first we will check the username and password

	if username == "Param" && password == "w" {
		// then we will create the claims and sign the claims

		claims := jwt.NewWithClaims(
			jwt.SigningMethodES256,
			jwt.MapClaims{
				"username":   username,
				"issued_at":  time.Now().Unix(),
				"expired_at": time.Now().Add(time.Hour).Unix(),
				"issuer":     "jwt-with-apis",
			},
		)

		jwtToken, err := claims.SignedString(privateSigningKey)

		if err != nil {
			return "", err
		}

		jwtToken = "Bearer " + jwtToken

		return jwtToken, nil
	}

	return "", errors.New("invalid username password")
}

func VerifyToken(jwtToken string) (*jwt.Token, error) {
	if jwtToken == "" {
		return nil, errors.New("empty JWT Token")
	}

	jwtToken = jwtToken[len("Bearer "):]

	// we will then decrypt this token
	token, err := jwt.Parse(jwtToken, func(*jwt.Token) (interface{}, error) { return publicSigningKey, nil })

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid Token")
	}

	return token, nil
}
