package auth

import (
	"context"

	"github.com/Sei-Yukinari/go-firebase-auth/src/api/domain"
	"github.com/Sei-Yukinari/go-firebase-auth/src/api/interfaces"
)

type Repository struct {
	AuthHandler interfaces.AuthHandler
}

func (repo *Repository) VerifyFirebaseToken(ctx context.Context, token string) (auth domain.Auth, err error) {
	if _, err = repo.AuthHandler.VerifyIDToken(ctx, token); err != nil {
		return
	}
	return
}
