package endpoint_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DangerZombie/golang-gin-boilerplate/endpoint"
	"github.com/DangerZombie/golang-gin-boilerplate/helper/auth"
	"github.com/DangerZombie/golang-gin-boilerplate/helper/message"
	"github.com/DangerZombie/golang-gin-boilerplate/helper/static"
	"github.com/DangerZombie/golang-gin-boilerplate/model/parameter"
	"github.com/DangerZombie/golang-gin-boilerplate/model/request"
	"github.com/DangerZombie/golang-gin-boilerplate/model/response"
	"github.com/DangerZombie/golang-gin-boilerplate/service/service_teacher"
	"github.com/gin-gonic/gin"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestEndpoint_CreateTeacher(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockTeacherService := service_teacher.NewMockTeacherService(mockCtrl)
	mockAuthHelper := auth.NewMockAuthHelper(mockCtrl)

	endpointModule := endpoint.NewEndpoint(
		mockAuthHelper,
	)

	id := faker.UUIDHyphenated()
	userId := faker.UUIDHyphenated()
	jobTitleId := faker.UUIDHyphenated()
	status := "PERMANENT"
	experience := 10
	degree := "S.Pd"
	issuerId := faker.UUIDHyphenated()
	username := faker.Email()
	nickname := faker.Name()
	claims := parameter.JwtClaims{
		Issuer:  issuerId,
		Subject: username,
		User:    nickname,
		Roles:   []string{static.RoleADMINISTRATOR},
	}

	teacherCreateRequest := request.TeacherCreateRequestBody{
		UserId:     userId,
		JobTitleId: jobTitleId,
		Status:     status,
		Experience: experience,
		Degree:     degree,
		Issuer:     issuerId,
	}

	teacherCreateResponse := response.TeacherCreateResponse{
		Id: id,
	}

	t.Run("Should return OK", func(t *testing.T) {
		reqBody, _ := json.Marshal(teacherCreateRequest)
		req := httptest.NewRequest(http.MethodPost, "/api/v1/teacher", strings.NewReader(string(reqBody)))
		req.Header.Set("Content-Type", gin.MIMEPlain)
		req.Header.Set("Authorization", faker.JWT)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Request = req

		mockAuthHelper.EXPECT().
			VerifyJWT(c.Request.Header).
			Times(1).
			Return(claims, nil)

		mockTeacherService.EXPECT().
			TeacherCreate(teacherCreateRequest).
			Times(1).
			Return(teacherCreateResponse, message.SuccessMsg, nil)

		statusCode, result := endpointModule.CreateTeacherRequest(c, mockTeacherService)

		require.Equal(t, http.StatusOK, statusCode)
		require.NotEmpty(t, result)
	})

	t.Run("Should return Unauthorized", func(t *testing.T) {
		reqBody, _ := json.Marshal(teacherCreateRequest)
		req := httptest.NewRequest(http.MethodPost, "/api/v1/teacher", strings.NewReader(string(reqBody)))
		req.Header.Set("Content-Type", gin.MIMEPlain)
		req.Header.Set("Authorization", faker.JWT)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Request = req

		mockAuthHelper.EXPECT().
			VerifyJWT(c.Request.Header).
			Times(1).
			Return(parameter.JwtClaims{}, errors.New("failed"))

		statusCode, result := endpointModule.CreateTeacherRequest(c, mockTeacherService)

		require.Equal(t, http.StatusUnauthorized, statusCode)
		require.NotEmpty(t, result)
	})

	t.Run("Should return Forbidden", func(t *testing.T) {
		reqBody, _ := json.Marshal(teacherCreateRequest)
		req := httptest.NewRequest(http.MethodPost, "/api/v1/teacher", strings.NewReader(string(reqBody)))
		req.Header.Set("Content-Type", gin.MIMEPlain)
		req.Header.Set("Authorization", faker.JWT)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Request = req

		mockAuthHelper.EXPECT().
			VerifyJWT(c.Request.Header).
			Times(1).
			Return(parameter.JwtClaims{}, nil)

		statusCode, result := endpointModule.CreateTeacherRequest(c, mockTeacherService)

		require.Equal(t, http.StatusForbidden, statusCode)
		require.NotEmpty(t, result)
	})

	t.Run("Should return Bad Request", func(t *testing.T) {
		reqBody, _ := json.Marshal(teacherCreateRequest)
		req := httptest.NewRequest(http.MethodPost, "/api/v1/teacher", strings.NewReader(string(reqBody)))
		req.Header.Set("Content-Type", gin.MIMEPlain)
		req.Header.Set("Authorization", faker.JWT)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Request = req

		mockAuthHelper.EXPECT().
			VerifyJWT(c.Request.Header).
			Times(1).
			Return(claims, nil)

		mockTeacherService.EXPECT().
			TeacherCreate(teacherCreateRequest).
			Times(1).
			Return(response.TeacherCreateResponse{}, message.ErrReqParam, nil)

		statusCode, result := endpointModule.CreateTeacherRequest(c, mockTeacherService)

		require.Equal(t, http.StatusBadRequest, statusCode)
		require.NotEmpty(t, result)
	})
}
