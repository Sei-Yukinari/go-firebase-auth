package auth_test

import (
	"api/domain"
	mockAuthUsecase "api/mocks/usecases/auth"
	"api/usecases/auth"
	"context"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

var err error

func mockAuthRepository(t *testing.T) *mockAuthUsecase.MockRepository {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	return mockAuthUsecase.NewMockRepository(ctrl)
}

func TestInteractor_VerifyFirebaseToken(t *testing.T) {
	token := "123"
	ctx := context.Background()
	var expected domain.Auth

	mockRepository := mockAuthRepository(t)
	mockRepository.EXPECT().VerifyFirebaseToken(ctx, token).Return(expected, err)

	usecase := auth.Interactor{AuthRepository: mockRepository}
	actual, err := usecase.VerifyFirebaseToken(ctx, token)

	if err != nil {
		t.Error("Actual VerifyFirebaseToken(ctx,token) is usecase Error")
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Actual VerifyFirebaseToken(ctx,token) is not same as expected")
	}
}
