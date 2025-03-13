package security

import (
	"errors"
	"log"
	"time"

	"github.com/allang-4779/financer/internal/util"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateAccessToken(payload jwt.MapClaims, expirySeconds uint)(string, error) {
	log.Print("Generating token with payload", payload);
	payload["exp"] = time.Now().Add(time.Second * time.Duration(expirySeconds)).Unix()
	payload["iat"] = time.Now().Unix()
	payload["iss"] = "financer"
	payload["sub"] = payload["email"]
	payload["aud"] = "financer"

	token := jwt.NewWithClaims(jwt.SigningMethodPS256.SigningMethodRSA, payload)

	tokenString, err := token.SignedString(util.GetPrivateKey())
	if err != nil {
		log.Println("Error generating token", err)
		return "", errors.New("could not generate access token")
	}
	return tokenString, nil



}

func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return util.GetPublicKey(), nil
	})
	if err != nil {
		log.Println("Error parsing token", err)
		return nil, errors.New("could not parse token")
	}
	if !token.Valid {
		return nil, errors.New("token is invalid")
	}
	return token.Claims.(jwt.MapClaims), nil
}