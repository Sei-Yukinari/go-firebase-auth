package middleware_test

import (
	"api/middleware"
	mock_interfaces "api/mocks/interfaces"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func mockAuthHandler(t *testing.T) *mock_interfaces.MockAuthHandler {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	return mock_interfaces.NewMockAuthHandler(ctrl)
}

func testHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if !f(w) {
		t.Fail()
	}
}

func testMiddlewareRequest(t *testing.T, r *gin.Engine, expectedHTTPCode int, token string) {
	req, _ := http.NewRequest("GET", "/", nil)
	if len(token) != 0 {
		req.Header.Set("Authorization", token)

	}
	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		return w.Code == expectedHTTPCode
	})
}

func TestMiddleware_Authorize_TokenNotExist(t *testing.T) {
	r := gin.Default()
	mockHandler := mockAuthHandler(t)
	r.GET("/", middleware.Authorize(mockHandler), func(c *gin.Context) {
		t.Fail()
	})
	testMiddlewareRequest(t, r, http.StatusBadRequest, "")
}

func TestMiddleware_Authorize_ValidToken(t *testing.T) {
	r := gin.Default()
	mockHandler := mockAuthHandler(t)
	mockToken := "123"
	mockHandler.EXPECT().VerifyIDToken(context.Background(), mockToken).Return(nil, nil)
	r.GET("/", middleware.Authorize(mockHandler), func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	testMiddlewareRequest(t, r, http.StatusOK, mockToken)
}
