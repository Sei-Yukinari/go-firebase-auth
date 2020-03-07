package auth

import (
	"context"

	"github.com/Sei-Yukinari/go-firebase-auth/src/api/domain"
)

type Repository interface {
	VerifyFirebaseToken(ctx context.Context, token string) (domain.Auth, error)
}
