package auth_test

import (
	"api/domain"
	"context"
	"reflect"
	"testing"
)

type fakeAuthController struct {
	Interactor fakeAuthInteractor
}

type fakeAuthInteractor struct {
	VerifyFirebaseToken func(context.Context, string) (domain.Auth, error)
}

func TestController_Index(t *testing.T) {
	token := "123"
	expected := domain.Auth{UID: 1, Name: "aaa"}

	controller := &fakeAuthController{
		Interactor: fakeAuthInteractor{
			VerifyFirebaseToken: func(ctx context.Context, token string) (domain.Auth, error) {
				return expected, nil
			},
		},
	}
	actual, err := controller.Interactor.VerifyFirebaseToken(context.Background(), token)
	if err != nil {
		t.Error("Actual Index() is controller Error")
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Error("Actual Index() is not same as expected")
	}
}
