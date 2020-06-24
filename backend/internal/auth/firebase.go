package auth

import (
	"context"
	"log"

	"firebase.google.com/go/v4"
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
		log.Fatalf("error verifying ID token: %v\n", err)
		return nil, err
	}
	return token, nil
}

// Firebase initializes connections to Firebase service
func Firebase(credentialsFileName *string) (*FirebaseAuth, error) {
	ctx := context.Background()

	var opts []option.ClientOption
	if credentialsFileName != nil {
		opts = append(opts, option.WithCredentialsFile(*credentialsFileName))
	}

	// init firebase app
	app, err := firebase.NewApp(ctx, nil, opts...)
	if err != nil {
		return nil, err
	}

	// init firebase auth client
	auth, err := app.Auth(ctx)

	return &FirebaseAuth{
		app:  app,
		auth: auth,
	}, nil
}
