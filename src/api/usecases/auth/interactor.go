package auth

import (
	"api/domain"
	"context"
)

type Interactor struct {
	AuthRepository Repository
}

func (ai *Interactor) VerifyFirebaseToken(ctx context.Context, token string) (auth domain.Auth, err error) {
	auth, err = ai.AuthRepository.VerifyFirebaseToken(ctx, token)
	return
}
