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

// Auth0KeyGetter is used to retrieve public key from Auth0 service
type Auth0KeyGetter struct {
	Options Auth0KeyGetterOptions
	JwksClient *jwks.Client
}

// Auth0KeyGetterOptions is used to configure Auth0KeyGetter
type Auth0KeyGetterOptions struct {
	ExpectedAudience string
	ExpectedIssuer string
	JwksEndpoint string
}

// NewAuth0KeyGetter creates a new Auth0KeyGetter
func NewAuth0KeyGetter(options Auth0KeyGetterOptions) *Auth0KeyGetter {
	return &Auth0KeyGetter{
		Options: options,
		JwksClient: jwks.NewClient(options.JwksEndpoint, jwks.NewConfig()),
	}
}

// GetValidationKey retrives a public key used to validate a JWT token
func (f *Auth0KeyGetter) GetValidationKey(token *jwt.Token) (interface{}, error) {
	if !f.validateAudience(token) {
		return token, errors.New("Invalid audience")
	}

	if !f.validateIssuer(token) {
		return token, errors.New("Invalid issuer")
	}

	cert, err := f.downloadKeyCert(token)

	if err != nil {
		return token, err
	}

	key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))

	if err != nil {
		return token, err
	}

	return key, nil
}

func (f *Auth0KeyGetter) validateAudience(token *jwt.Token) bool {
	return token.Claims.(jwt.MapClaims).VerifyAudience(f.Options.ExpectedAudience, false)
}

func (f *Auth0KeyGetter) validateIssuer(token *jwt.Token) bool {
	return token.Claims.(jwt.MapClaims).VerifyIssuer(f.Options.ExpectedIssuer, false)
}

func (f *Auth0KeyGetter) downloadKeyCert(token *jwt.Token) (string, error)  {
	kid := token.Header["kid"].(string)
	key, err := f.JwksClient.GetSigningKey(kid)

	if err != nil {
		return "", err
	}

	if key == nil {
		return "", errors.New("Unable to find appropriate key")
	}

	cert := "-----BEGIN CERTIFICATE-----\n" + key.X5c[0] + "\n-----END CERTIFICATE-----"

	return cert, nil
}

