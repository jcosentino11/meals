package auth

import (
	"context"
	"log"
	"strings"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

// FirebaseAuth holds the firebase authentication client
type FirebaseAuth struct {
	app  *firebase.App
	auth *auth.Client
}

// VerifyAuthToken checks that a given authorization token is valid
func (f *FirebaseAuth) VerifyAuthToken(ctx context.Context, authToken string) (*auth.Token, error) {
	token, err := f.auth.VerifyIDToken(ctx, authToken)
	if err != nil {
		log.Printf("error verifying ID token: %v\n", err)
		return nil, err
	}
	return token, nil
}

// Firebase initializes connections to Firebase service
func Firebase(credsPath *string) (*FirebaseAuth, error) {
	ctx := context.Background()

	var opts []option.ClientOption
	if credsPath != nil {
		log.Printf("Initializing Firebase client with credentials from %s", *credsPath)
		opts = append(opts, option.WithCredentialsFile(*credsPath))
	} else {
		log.Printf("Initialize Firebase with default google credentials")
	}

	// init firebase app
	app, err := firebase.NewApp(ctx, nil, opts...)
	if err != nil {
		return nil, err
	}

	// init firebase auth client
	auth, err := app.Auth(ctx)

	if err != nil {
		return nil, err
	}

	return &FirebaseAuth{
		app:  app,
		auth: auth,
	}, nil
}
