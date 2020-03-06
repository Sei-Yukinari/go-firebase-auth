package auth

import (
	"api/domain"
	"api/interfaces"
	"context"
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
