package user_test

import (
	"api/domain"
	"reflect"
	"testing"
)

var err error

type fakeUserController struct {
	Interactor fakeUserInteractor
}

type fakeUserInteractor struct {
	FindAll func() (domain.Users, error)
	Show    func(int) (domain.User, error)
}

func TestController_Index(t *testing.T) {
	expected := domain.Users{
		domain.User{ID: 1, Name: "aaa"},
		domain.User{ID: 2, Name: "bbb"},
		domain.User{ID: 3, Name: "ccc"},
	}

	controller := &fakeUserController{
		Interactor: fakeUserInteractor{
			FindAll: func() (domain.Users, error) {
				return expected, nil
			},
		},
	}
	actual, err := controller.Interactor.FindAll()
	if err != nil {
		t.Error("Actual Index() is controller Error")
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Error("Actual Index() is not same as expected")
	}
}

func TestController_Show(t *testing.T) {
	expected := domain.User{ID: 1, Name: "aaa"}
	userID := 1
	controller := &fakeUserController{
		Interactor: fakeUserInteractor{
			Show: func(id int) (domain.User, error) {
				return expected, nil
			},
		},
	}
	actual, err := controller.Interactor.Show(userID)
	if err != nil {
		t.Error("Actual Show(id) is controller Error")
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Error("Actual Show(id) is not same as expected")
	}
}
