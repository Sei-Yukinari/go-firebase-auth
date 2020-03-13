package userUsecase

import (
	"api/domain"
	mockUserUsecase "api/mocks/usecases/user"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

var err error

func mockUserRepository(t *testing.T) *mockUserUsecase.MockUserRepository {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	return mockUserUsecase.NewMockUserRepository(ctrl)
}

func TestMain(m *testing.M) {
	println("before all...")
	m.Run()
	println("after all...")
}

func TestFindAll(t *testing.T) {

	var expected domain.Users

	mockRepository := mockUserRepository(t)
	mockRepository.EXPECT().FindAll().Return(expected, err)

	usecase := UserInteractor{UserRepository: mockRepository}
	actual, err := usecase.FindAll()

	if err != nil {
		t.Error("Actual FindAll() is usecase Error")
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Actual FindAll() is not same as expected")
	}
}

func TestShow(t *testing.T) {

	var expected domain.User
	mockId := 1

	mockRepository := mockUserRepository(t)
	mockRepository.EXPECT().FindByID(mockId).Return(expected, err)

	usecase := UserInteractor{UserRepository: mockRepository}
	actual, err := usecase.Show(mockId)

	if err != nil {
		t.Error("Actual FindByID(id) is usecase Error")
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Actual FindByID(id) is not same as expected")
	}
}
