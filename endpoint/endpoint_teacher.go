package endpoint

import (
	"encoding/json"
	"net/http"

	"github.com/DangerZombie/golang-gin-boilerplate/helper/message"
	"github.com/DangerZombie/golang-gin-boilerplate/helper/static"
	"github.com/DangerZombie/golang-gin-boilerplate/helper/util"
	"github.com/DangerZombie/golang-gin-boilerplate/model/base"
	"github.com/DangerZombie/golang-gin-boilerplate/model/request"
	"github.com/DangerZombie/golang-gin-boilerplate/service/service_teacher"
	"github.com/gin-gonic/gin"
	"github.com/go-faker/faker/v4/pkg/slice"
)

func (e *endpointImpl) CreateTeacherRequest(ctx *gin.Context, s service_teacher.TeacherService) (int, interface{}) {
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

	req := request.TeacherCreateRequestBody{}
	_ = json.NewDecoder(ctx.Request.Body).Decode(&req)
	req.Issuer = claims.Issuer
	result, msg, errMsg := s.TeacherCreate(req)

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

func (e *endpointImpl) ListTeachersRequest(ctx *gin.Context, s service_teacher.TeacherService) (int, interface{}) {
	// Verify JWT token from the request headers
	_, err := e.authHelper.VerifyJWT(ctx.Request.Header)
	if err != nil {
		wrap := base.SetHttpResponse(message.ErrNoAuth.Code, message.ErrNoAuth.Message, nil, nil, map[string]string{"token": err.Error()})
		return http.StatusUnauthorized, wrap
	}

	req := request.TeacherListRequest{
		Page:  util.StringToInt(ctx.Query("page")),
		Limit: util.StringToInt(ctx.Query("limit")),
		Sort:  ctx.Query("sort"),
		Dir:   ctx.Query("dir"),
		Name:  ctx.Query("name"),
	}

	result, pagination, msg, errMsg := s.TeacherList(req)

	var wrap interface{}
	var code int
	if msg.Code == 4000 {
		code = http.StatusBadRequest
		wrap = base.SetHttpResponse(msg.Code, msg.Message, nil, nil, errMsg)
	} else {
		code = http.StatusOK
		wrap = base.SetHttpResponse(msg.Code, msg.Message, result, &pagination, errMsg)
	}

	return code, wrap
}

func (e *endpointImpl) FindTeacherDetailRequest(ctx *gin.Context, s service_teacher.TeacherService) (int, interface{}) {
	// Verify JWT token from the request headers
	_, err := e.authHelper.VerifyJWT(ctx.Request.Header)
	if err != nil {
		wrap := base.SetHttpResponse(message.ErrNoAuth.Code, message.ErrNoAuth.Message, nil, nil, map[string]string{"token": err.Error()})
		return http.StatusUnauthorized, wrap
	}

	req := request.TeacherDetailRequest{}
	req.Id = ctx.Param("id")
	result, msg, errMsg := s.TeacherDetail(req)

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

func (e *endpointImpl) UpdateTeacherRequest(ctx *gin.Context, s service_teacher.TeacherService) (int, interface{}) {
	// Verify JWT token from the request headers
	_, err := e.authHelper.VerifyJWT(ctx.Request.Header)
	if err != nil {
		wrap := base.SetHttpResponse(message.ErrNoAuth.Code, message.ErrNoAuth.Message, nil, nil, map[string]string{"token": err.Error()})
		return http.StatusUnauthorized, wrap
	}

	reqBody := request.TeacherUpdateRequestBody{}
	_ = json.NewDecoder(ctx.Request.Body).Decode(&reqBody)
	req := request.TeacherUpdateRequest{
		Id:   ctx.Param("id"),
		Body: reqBody,
	}

	result, msg, errMsg := s.TeacherUpdate(req)

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

func (e *endpointImpl) DeleteTeacherRequest(ctx *gin.Context, s service_teacher.TeacherService) (int, interface{}) {
	// Verify JWT token from the request headers
	_, err := e.authHelper.VerifyJWT(ctx.Request.Header)
	if err != nil {
		wrap := base.SetHttpResponse(message.ErrNoAuth.Code, message.ErrNoAuth.Message, nil, nil, map[string]string{"token": err.Error()})
		return http.StatusUnauthorized, wrap
	}

	req := request.TeacherDeleteRequest{}
	req.Id = ctx.Param("id")
	result, msg, errMsg := s.TeacherDelete(req)

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
