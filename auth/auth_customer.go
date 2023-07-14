package auth

import (
	"errors"

	"github.com/golang-jwt/jwt/v4"
)

type AuthCustomer interface {
	GenerateToken(userID uint64) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtAuthCustomer struct {
}

func NewCustomerService() Auth {
	return &jwtAuth{}
}

var SECRET_KEYS = []byte("tokokecilkita_secret_key")

func (s *jwtAuth) CustomerGenerateToken(userID uint64) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func (s *jwtAuth) CustomerValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Invalid Token")
		}

		return []byte(SECRET_KEYS), nil
	})

	if err != nil {
		return token, nil
	}

	return token, nil
}
