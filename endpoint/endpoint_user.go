package endpoint

import (
	"encoding/json"
	"net/http"

	"github.com/DangerZombie/golang-gin-boilerplate/helper/message"
	"github.com/DangerZombie/golang-gin-boilerplate/helper/static"
	"github.com/DangerZombie/golang-gin-boilerplate/model/base"
	"github.com/DangerZombie/golang-gin-boilerplate/model/request"
	"github.com/DangerZombie/golang-gin-boilerplate/service/service_user"
	"github.com/gin-gonic/gin"
	"github.com/go-faker/faker/v4/pkg/slice"
)

func (e *endpointImpl) LoginRequest(ctx *gin.Context, s service_user.UserService) (int, interface{}) {
	req := request.LoginRequestBody{}
	_ = json.NewDecoder(ctx.Request.Body).Decode(&req)
	result, msg, errMsg := s.Login(req)

	var wrap interface{}
	var code int
	if msg.Code == 4000 {
		code = http.StatusBadRequest
		wrap = base.SetHttpResponse(msg.Code, msg.Message, nil, nil, errMsg)
	} else {
		code = http.StatusOK
		wrap = base.SetHttpResponse(msg.Code, msg.Message, result, nil, errMsg)
	}

	return code, wrap
}

func (e *endpointImpl) UserProfileRequest(ctx *gin.Context, s service_user.UserService) (int, interface{}) {
	// Verify JWT token from the request headers
	_, err := e.authHelper.VerifyJWT(ctx.Request.Header)
	if err != nil {
		wrap := base.SetHttpResponse(message.ErrNoAuth.Code, message.ErrNoAuth.Message, nil, nil, map[string]string{"token": err.Error()})
		return http.StatusUnauthorized, wrap
	}

	userProfileInput := request.UserProfileRequest{
		Id: ctx.Query("id"),
	}

	result, msg, errMsg := s.UserProfile(userProfileInput)

	var wrap interface{}
	var code int
	if msg.Code == 4000 {
		code = http.StatusBadRequest
		wrap = base.SetHttpResponse(msg.Code, msg.Message, nil, nil, errMsg)
	} else {
		code = http.StatusOK
		wrap = base.SetHttpResponse(msg.Code, msg.Message, result, nil, errMsg)
	}

	return code, wrap
}

func (e *endpointImpl) RegisterUserRequest(ctx *gin.Context, s service_user.UserService) (int, interface{}) {
	// Verify JWT token from the request headers
	claims, err := e.authHelper.VerifyJWT(ctx.Request.Header)
	if err != nil {
		wrap := base.SetHttpResponse(message.ErrNoAuth.Code, message.ErrNoAuth.Message, nil, nil, map[string]string{"token": err.Error()})
		return http.StatusUnauthorized, wrap
	}

	// Validate allowed roles
	if !slice.Contains(claims.Roles, static.RoleADMINISTRATOR) {
		wrap := base.SetHttpResponse(message.ErrForbiddenAccess.Code, message.ErrForbiddenAccess.Message, nil, nil, map[string]string{"role": "you not have correct role"})
		return http.StatusForbidden, wrap
	}

	req := request.RegisterUserRequestBody{}
	_ = json.NewDecoder(ctx.Request.Body).Decode(&req)
	req.Issuer = claims.Issuer
	result, msg, errMsg := s.RegisterUser(req)

	var wrap interface{}
	var code int
	if msg.Code == 4000 {
		code = http.StatusBadRequest
		wrap = base.SetHttpResponse(msg.Code, msg.Message, nil, nil, errMsg)
	} else {
		code = http.StatusOK
		wrap = base.SetHttpResponse(msg.Code, msg.Message, result, nil, errMsg)
	}

	return code, wrap
}
