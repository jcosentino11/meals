package main

import (
	"errors"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	jwks "github.com/hkra/go-jwks"
)

// Jwt provides token verification capabilities
type Jwt struct {
	signingMethod jwt.SigningMethod
	audience      string
	issuer        string
	jwksClient    *jwks.Client
}

// JwtConfig holds config options for Jwt
type JwtConfig struct {
	Audience     string
	Issuer       string
	JwksEndpoint string
}

// NewJwt creates and initializes a new jwt verifier
func NewJwt(config JwtConfig) *Jwt {
	return &Jwt{
		signingMethod: jwt.SigningMethodRS256,
		audience:      config.Audience,
		issuer:        config.Issuer,
		jwksClient:    jwks.NewClient(config.JwksEndpoint, jwks.NewConfig()),
	}
}

// ParseTokenFromRequest parses jwt token from an http request
// The returned token is fully verified and valid when error is nil.
func (j *Jwt) ParseTokenFromRequest(request *http.Request) (*jwt.Token, error) {
	token, err := j.getAuthHeaderToken(request)
	if err != nil || token == "" {
		return nil, errors.New("Unable to get token from header")
	}

	parsedToken, err := j.parseToken(token)
	if err != nil {
		return nil, err
	}

	if err := j.validateSigningMethod(parsedToken); err != nil {
		return nil, err
	}

	if !parsedToken.Valid {
		return nil, errors.New("auth token is invalid")
	}

	return parsedToken, nil
}

func (j *Jwt) parseToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, j.downloadValidationKey)
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

func (j *Jwt) downloadValidationKey(token *jwt.Token) (interface{}, error) {
	if !j.validateAudience(token) {
		return token, errors.New("Invalid audience")
	}

	if !j.validateIssuer(token) {
		return token, errors.New("Invalid issuer")
	}

	cert, err := j.downloadKeyCert(token)

	if err != nil {
		return token, err
	}

	key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))

	if err != nil {
		return token, err
	}

	return key, nil
}

func (j *Jwt) downloadKeyCert(token *jwt.Token) (string, error) {
	kid := token.Header["kid"].(string)
	key, err := j.jwksClient.GetSigningKey(kid)

	if err != nil {
		return "", err
	}

	if key == nil {
		return "", errors.New("Unable to find appropriate key")
	}

	cert := "-----BEGIN CERTIFICATE-----\n" + key.X5c[0] + "\n-----END CERTIFICATE-----"

	return cert, nil
}

func (j *Jwt) validateSigningMethod(token *jwt.Token) error {
	if j.signingMethod != nil && j.signingMethod.Alg() != token.Header["alg"] {
		return errors.New("Invalid signing method")
	}
	return nil
}

func (j *Jwt) validateAudience(token *jwt.Token) bool {
	return token.Claims.(jwt.MapClaims).VerifyAudience(j.audience, false)
}

func (j *Jwt) validateIssuer(token *jwt.Token) bool {
	return token.Claims.(jwt.MapClaims).VerifyIssuer(j.issuer, false)
}
