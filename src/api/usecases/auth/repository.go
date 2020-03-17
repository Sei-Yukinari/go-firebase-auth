package auth

import (
	"context"

	"api/domain"
)

type Repository interface {
	VerifyFirebaseToken(ctx context.Context, token string) (domain.Auth, error)
}
