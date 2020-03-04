package auth

import (
	"api/infrastructure/env"
	"api/interfaces"
	"context"
	"path/filepath"

	fire "firebase.google.com/go"
	"google.golang.org/api/option"

	"firebase.google.com/go/auth"
)

type Handler struct {
	Client *auth.Client
}

func NewAuthHandler() (interfaces.AuthHandler, error) {
	var env = env.Load()
	filename, err := filepath.Abs(env.Credentials)
	if err != nil {
		panic(err.Error)
	}
	opt := option.WithCredentialsFile(filename)
	app, err := fire.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic(err.Error)
	}
	client, err := app.Auth(context.Background())
	if err != nil {
		panic(err.Error)
	}
	authHandler := new(Handler)
	authHandler.Client = client
	return authHandler, nil
}

func (handler Handler) VerifyIDToken(ctx context.Context, token string) (*auth.Token, error) {
	return handler.Client.VerifyIDToken(ctx, token)
}
