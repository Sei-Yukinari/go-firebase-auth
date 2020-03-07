package userUsecase

import (
	"github.com/Sei-Yukinari/go-firebase-auth/src/api/domain"
)

// UserRepository belong to the usecases layer.
type UserRepository interface {
	FindAll() (domain.Users, error)
	FindByID(int) (domain.User, error)
}
