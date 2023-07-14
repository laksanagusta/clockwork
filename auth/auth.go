package auth

import (
	"clockwork-server/config"
	"errors"

	"github.com/golang-jwt/jwt/v4"
)

type Auth interface {
	GenerateToken(userID uint64) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtAuth struct {
	config config.Config
}

func NewService(config config.Config) Auth {
	return &jwtAuth{config}
}

func (s *jwtAuth) GenerateToken(userID uint64) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString(s.config.JWT.SecretKey)
	if err != nil {
		return signedToken, err
	}

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
