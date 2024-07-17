package endpoint

import (
	"github.com/DangerZombie/golang-gin-boilerplate/service/service_teacher"
	"github.com/DangerZombie/golang-gin-boilerplate/service/service_user"
	"github.com/gin-gonic/gin"
)

type Endpoint interface {
	// Endpoint User
	LoginRequest(ctx *gin.Context, s service_user.UserService) (int, interface{})
	RegisterUserRequest(ctx *gin.Context, s service_user.UserService) (int, interface{})
	UserProfileRequest(ctx *gin.Context, s service_user.UserService) (int, interface{})

	// Endpoint Teacher
	CreateTeacherRequest(ctx *gin.Context, s service_teacher.TeacherService) (int, interface{})
	DeleteTeacherRequest(ctx *gin.Context, s service_teacher.TeacherService) (int, interface{})
	FindTeacherDetailRequest(ctx *gin.Context, s service_teacher.TeacherService) (int, interface{})
	ListTeachersRequest(ctx *gin.Context, s service_teacher.TeacherService) (int, interface{})
	UpdateTeacherRequest(ctx *gin.Context, s service_teacher.TeacherService) (int, interface{})
}
