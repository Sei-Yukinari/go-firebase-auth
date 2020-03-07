package auth

import (
	"context"

	"github.com/Sei-Yukinari/go-firebase-auth/src/api/domain"
	"github.com/Sei-Yukinari/go-firebase-auth/src/api/interfaces"
	"github.com/Sei-Yukinari/go-firebase-auth/src/api/usecases"

	authUsecase "github.com/Sei-Yukinari/go-firebase-auth/src/api/usecases/auth"
)

// A UserController belong to the interface layer.
type Controller struct {
	AuthInteractor authUsecase.Interactor
	Logger         usecases.Logger
}

// NewUserController returns the resource of users.
func NewAuthController(authHandler interfaces.AuthHandler, logger usecases.Logger) *Controller {
	return &Controller{
		AuthInteractor: authUsecase.Interactor{
			AuthRepository: &Repository{
				AuthHandler: authHandler,
			},
		},
		Logger: logger,
	}
}

// Index return response which contain a listing of the resource of users.
func (ac *Controller) Index(c interfaces.Context) {
	ac.AuthInteractor.VerifyFirebaseToken(context.Background(), "123")
	a := domain.Auth{UID: 1, Name: "2222"}
	_ = c.Bind(&a)
	auth, err := ac.AuthInteractor.VerifyFirebaseToken(context.Background(), "123")
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, auth)
}
