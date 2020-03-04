package auth

import (
	"api/domain"
	"api/interfaces"
	"api/usecases"
	"context"

	authUsecase "github.com/go-firebase-auth/src/api/usecases/auth"
)

// A UserController belong to the interface layer.
type Controller struct {
	AuthInteractor authUsecase.Interactor
	Logger         usecases.Logger
}

// NewUserController returns the resource of users.
func NewAuthController(logger usecases.Logger) *Controller {
	return &Controller{
		Logger: logger,
	}
}

// Index return response which contain a listing of the resource of users.
func (ac *Controller) Index(c interfaces.Context) {
	ac.AuthInteractor.VerifyFirebaseToken(context.Background(), "123")
	a := domain.Auth{UID: 1, Name: "2222"}
	println("main logic")
	_ = c.Bind(&a)
	//auth := ac.AuthInteractor.VerifyFirebaseToken(c)
	//if err != nil {
	//	c.JSON(500, err)
	//	return
	//}
	c.JSON(200, a)
}
