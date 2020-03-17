package user

import (
	"api/domain"
	"api/mocks/usecases/user"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

var err error

func mockUserRepository(t *testing.T) *mock_userUsecase.MockUserRepository {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	return mock_userUsecase.NewMockUserRepository(ctrl)
}

func TestFindAll(t *testing.T) {

	expected := domain.Users{
		domain.User{ID: 1, Name: "aaa"},
		domain.User{ID: 2, Name: "bbb"},
		domain.User{ID: 3, Name: "ccc"},
	}

	mockUserRepository := mockUserRepository(t)

	mockUserRepository.EXPECT().FindAll().Return(expected, err)
	actual, err := mockUserRepository.FindAll()

	if err != nil {
		t.Error("Actual FindAll() is usecase Error")
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Actual FindAll() is not same as expected")
	}
}

func TestFindByID(t *testing.T) {

	expected := domain.User{ID: 1, Name: "aaa"}

	mockUserRepository := mockUserRepository(t)

	mockUserRepository.EXPECT().FindByID(1).Return(expected, err)
	actual, err := mockUserRepository.FindByID(1)

	if err != nil {
		t.Error("Actual FindByID(1) is repository Error")
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Actual FindByID(1) is not same as expected")
	}
}
