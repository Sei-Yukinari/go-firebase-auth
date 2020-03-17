package auth_test

import (
	"api/domain"
	mockauth "api/mocks/usecases/auth"
	"context"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

var err error

func mockAuthRepository(t *testing.T) *mockauth.MockRepository {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	return mockauth.NewMockRepository(ctrl)
}

func TestRepository_VerifyFirebaseToken(t *testing.T) {
	token := "123"
	ctx := context.Background()
	expected := domain.Auth{UID: 1, Name: "aaa"}

	mockAuthRepository := mockAuthRepository(t)

	mockAuthRepository.EXPECT().VerifyFirebaseToken(ctx, token).Return(expected, err)
	actual, err := mockAuthRepository.VerifyFirebaseToken(ctx, token)

	if err != nil {
		t.Error("Actual VerifyFirebaseToken(ctx,token) is usecase Error")
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Actual VerifyFirebaseToken(ctx,token) is not same as expected")
	}
}
