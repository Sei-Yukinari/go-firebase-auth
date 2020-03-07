package userUsecase

import "github.com/Sei-Yukinari/go-firebase-auth/src/api/domain"

// A UserInteractor belong to the usecases layer.
type UserInteractor struct {
	UserRepository UserRepository
}

// Index is display a listing of the resource.
func (ui *UserInteractor) Index() (users domain.Users, err error) {
	users, err = ui.UserRepository.FindAll()
	return
}

// Show is display the specified resource.
func (ui *UserInteractor) Show(id int) (user domain.User, err error) {
	user, err = ui.UserRepository.FindByID(id)
	return
}
