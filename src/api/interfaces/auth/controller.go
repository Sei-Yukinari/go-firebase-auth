package auth

import (
	"api/infrastructure/util"
	"context"
	"net/http"

	"api/domain"
	"api/interfaces"
	"api/usecases"

	authUsecase "api/usecases/auth"
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

// Show godoc
// @Summary Show auth
// @Description get Auth
// @Tags auth
// @Accept  json
// @Produce  json
// @Success 200 {object} domain.Auth
// @Failure 500 {object} util.HTTPError
// @Router /auth [get]
func (ac *Controller) Index(c interfaces.Context) {
	ac.AuthInteractor.VerifyFirebaseToken(context.Background(), "123")
	a := domain.Auth{UID: 1, Name: "2222"}
	_ = c.Bind(&a)
	auth, err := ac.AuthInteractor.VerifyFirebaseToken(context.Background(), "123")
	if err != nil {
		util.NewError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, auth)
}
