package auth

import (
	"api/domain"
	"context"
)

type Repository interface {
	VerifyFirebaseToken(ctx context.Context, token string) (domain.Auth, error)
}
