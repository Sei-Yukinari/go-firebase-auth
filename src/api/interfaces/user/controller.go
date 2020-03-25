package user

import (
	"api/domain"
	"api/infrastructure/util"
	"net/http"
	"strconv"

	"api/interfaces"
	"api/usecases"

	userUsecase "api/usecases/user"
)

// A UserController belong to the interface layer.
type Controller struct {
	UserInteractor userUsecase.UserInteractor
	Logger         usecases.Logger
}

// NewUserController returns the resource of users.
func NewUserController(sqlHandler interfaces.SQLHandler, logger usecases.Logger) *Controller {
	return &Controller{
		UserInteractor: userUsecase.UserInteractor{
			UserRepository: &Repository{
				SQLHandler: sqlHandler,
			},
		},
		Logger: logger,
	}
}

// Index godoc
// @Summary Show users
// @Description get user all
// @Tags user
// @Accept  json
// @Produce  json
// @Success 200 {object} domain.Users
// @Failure 500 {object} util.HTTPError
// @Router /user [get]
func (uc *Controller) Index(c interfaces.Context) {
	u := domain.User{}
	c.Bind(&u)
	users, err := uc.UserInteractor.FindAll()
	if err != nil {
		util.NewError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, users)
}

// Show godoc
// @Summary Show user
// @Description get user by ID
// @Tags user
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} domain.User
// @Failure 500 {object} util.HTTPError
// @Router /user/{id} [get]
func (uc *Controller) Show(c interfaces.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := uc.UserInteractor.Show(id)
	if err != nil {
		util.NewError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, user)
}
