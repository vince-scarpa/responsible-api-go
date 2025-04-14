package auth

import (
	"responsible-api-go/concerns"
	"responsible-api-go/internal"

	"github.com/golang-jwt/jwt/v5"
)

func NewAuth(options concerns.Options) *AuthWrapper {
	return &AuthWrapper{
		Auth: &concerns.Auth{
			Options: options,
		},
	}
}

type AuthWrapper struct {
	Auth *concerns.Auth
}

func (a *AuthWrapper) GrantToken(userID string, Hash string) (string, error) {
	tokenString, err := internal.Grant(userID, Hash, a.Auth.Options)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (a *AuthWrapper) ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := internal.Validate(tokenString, a.Auth.Options)
	if err != nil {
		return nil, err
	}
	return token, nil
}
