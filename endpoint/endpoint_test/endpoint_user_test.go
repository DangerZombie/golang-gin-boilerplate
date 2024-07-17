package endpoint_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DangerZombie/golang-gin-boilerplate/endpoint"
	"github.com/DangerZombie/golang-gin-boilerplate/helper/auth"
	"github.com/DangerZombie/golang-gin-boilerplate/helper/message"
	"github.com/DangerZombie/golang-gin-boilerplate/model/request"
	"github.com/DangerZombie/golang-gin-boilerplate/model/response"
	"github.com/DangerZombie/golang-gin-boilerplate/service/service_user"
	"github.com/gin-gonic/gin"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestEndpointUser_Login(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockUserService := service_user.NewMockUserService(mockCtrl)
	mockAuthHelper := auth.NewMockAuthHelper(mockCtrl)

	endpointModule := endpoint.NewEndpoint(
		mockAuthHelper,
	)

	loginRequest := request.LoginRequestBody{
		Username: faker.Username(),
		Password: faker.Name(),
	}

	loginResponse := response.LoginResponse{
		Token: faker.Jwt(),
	}

	t.Run("Should return OK", func(t *testing.T) {
		reqBody, _ := json.Marshal(loginRequest)
		req := httptest.NewRequest(http.MethodPost, "/api/v1/user/login", strings.NewReader(string(reqBody)))
		req.Header.Set("Content-Type", gin.MIMEPlain)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Request = req

		mockUserService.EXPECT().
			Login(loginRequest).
			Times(1).
			Return(loginResponse, message.SuccessMsg, nil)

		statusCode, result := endpointModule.LoginRequest(c, mockUserService)

		require.Equal(t, http.StatusOK, statusCode)
		require.NotEmpty(t, result)
	})

	t.Run("Should return Bad Request if something went wrong", func(t *testing.T) {
		reqBody, _ := json.Marshal(loginRequest)
		req := httptest.NewRequest(http.MethodPost, "/api/v1/user/login", strings.NewReader(string(reqBody)))
		req.Header.Set("Content-Type", gin.MIMEPlain)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Request = req

		mockUserService.EXPECT().
			Login(loginRequest).
			Times(1).
			Return(response.LoginResponse{}, message.ErrReqParam, nil)

		statusCode, result := endpointModule.LoginRequest(c, mockUserService)

		require.Equal(t, http.StatusBadRequest, statusCode)
		require.NotEmpty(t, result)
	})
}
