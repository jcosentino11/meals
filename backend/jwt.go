package main

import (
	"errors"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

// Jwt provides token verification capabilities
type Jwt struct {
	KeyFunc       jwt.Keyfunc
	SigningMethod jwt.SigningMethod
}

// ParseTokenFromRequest parses jwt token from an http request
func (j *Jwt) ParseTokenFromRequest(request *http.Request) (*jwt.Token, error) {
	token, err := j.getAuthHeaderToken(request)
	if err != nil || token == "" {
		return nil, errors.New("Unable to get token from header")
	}
	return j.ParseToken(token)
}

// ParseToken parses jwt token from string format
func (j *Jwt) ParseToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, j.KeyFunc)
}

func (j *Jwt) getAuthHeaderToken(request *http.Request) (string, error) {
	authHeader := request.Header.Get("Authorization")

	if authHeader == "" {
		return "", nil
	}

	authHeaderFields := strings.Fields(authHeader)

	if len(authHeaderFields) != 2 || authHeaderFields[0] != "Bearer" {
		return "", errors.New("Invalid auth header format")
	}

	return authHeaderFields[1], nil
}

// VerifySigningMethod checks that the token's signing method is expected
func (j *Jwt) VerifySigningMethod(token *jwt.Token) error {
	if j.SigningMethod != nil && j.SigningMethod.Alg() != token.Header["alg"] {
		return errors.New("Invalid signing method")
	}
	return nil
}
