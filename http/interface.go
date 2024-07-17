package http

import (
	"github.com/DangerZombie/golang-gin-boilerplate/service/service_teacher"
	"github.com/DangerZombie/golang-gin-boilerplate/service/service_user"
	"github.com/gin-gonic/gin"
)

type Http interface {
	// Swagger documentation
	SwaggerHttpHandler(r *gin.Engine)

	// API handler
	TeacherHandler(group *gin.RouterGroup, s service_teacher.TeacherService)
	UserHandler(group *gin.RouterGroup, s service_user.UserService)
}
