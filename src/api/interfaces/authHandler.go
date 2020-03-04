package interfaces

import (
	"context"

	"firebase.google.com/go/auth"
)

type AuthHandler interface {
	VerifyIDToken(context.Context, string) (*auth.Token, error)
}
