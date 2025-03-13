package util

import (
	"crypto/rsa"
	"fmt"
	"os"
	"log"
	"github.com/golang-jwt/jwt/v5"
)

// Global variables for keys
var privateKey *rsa.PrivateKey
var publicKey *rsa.PublicKey

// Load RSA keys from files
func LoadKeys(privatePath string, publicPath string) {
	privateBytes, err := os.ReadFile(privatePath)
	if err != nil {
		log.Fatal("Could not read private key:", err)
	}
	publicBytes, err := os.ReadFile(publicPath)
	if err != nil {
		log.Fatal("Could not read public key:", err)
	}

	// Parse private key
	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil {
		log.Fatal("Could not parse private key:", err)
	}

	// Parse public key
	publicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil {
		log.Fatal("Could not parse public key:", err)
	}

	fmt.Println("RSA keys loaded successfully!")
}
func GetPrivateKey() *rsa.PrivateKey {
    return privateKey
}

// GetPublicKey returns the public key
func GetPublicKey() *rsa.PublicKey {
    return publicKey
}