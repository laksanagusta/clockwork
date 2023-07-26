package auth

import (
	"clockwork-server/config"
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

type Auth interface {
	GenerateToken(entityID int, email string, userType string) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtAuth struct {
	config config.Config
}

func NewService(config config.Config) Auth {
	return &jwtAuth{config}
}

func (s *jwtAuth) GenerateToken(entityID int, email string, userType string) (string, error) {
	claim := jwt.MapClaims{}

	if userType == "user" {
		claim["user_id"] = entityID
	} else {
		claim["customer_id"] = entityID
	}

	claim["email"] = email

	fmt.Println(claim)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString(s.config.JWT.SecretKey)
	if err != nil {
		return signedToken, err
	}

	fmt.Println(signedToken)

	return signedToken, nil
}

func (s *jwtAuth) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Invalid Token")
		}

		return []byte(s.config.JWT.SecretKey), nil
	})

	if err != nil {
		return token, nil
	}

	return token, nil
}
