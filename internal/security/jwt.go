package security

import (
	"errors"
	"log"
	"time"

	"github.com/allang-4779/financer/internal/types"
	"github.com/allang-4779/financer/internal/util"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateAccessToken(payload jwt.MapClaims, expirySeconds int) (types.TokenResponse, error) {
	log.Print("Generating token with payload", payload)
	var iat = time.Now();
	var exp =iat.Add(time.Second * time.Duration(expirySeconds)).Unix()
	payload["exp"] = exp
	payload["iat"] =iat.Unix()
	payload["iss"] = "financer"
	payload["sub"] = payload["username"]
	payload["aud"] = "financer"

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, payload)

	tokenString, err := token.SignedString(util.GetPrivateKey())
	if err != nil {
		log.Println("Error generating token", err)
		return types.TokenResponse{}, errors.New("could not generate access token")
	}
	log.Println("Token generated successfully")
	tokenResponse := types.TokenResponse{
		Token:     tokenString,
		IssuedAt:  iat.Unix(),
		ExpiresAt: exp,
		ValidFor:  expirySeconds,
	}
	return tokenResponse, nil

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
