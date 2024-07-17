package http

import (
	"github.com/DangerZombie/golang-gin-boilerplate/endpoint"
	"github.com/DangerZombie/golang-gin-boilerplate/service/service_teacher"
	"github.com/gin-gonic/gin"
)

func (h *httpImpl) TeacherHandler(group *gin.RouterGroup, s service_teacher.TeacherService) {
	group.POST("", func(ctx *gin.Context) {
		statusCode, result := endpoint.NewEndpoint(h.authHelper).CreateTeacherRequest(ctx, s)
		ctx.JSON(statusCode, result)
	})

	group.GET("", func(ctx *gin.Context) {
		statusCode, result := endpoint.NewEndpoint(h.authHelper).ListTeachersRequest(ctx, s)
		ctx.JSON(statusCode, result)
	})

	group.GET("/:id", func(ctx *gin.Context) {
		statusCode, result := endpoint.NewEndpoint(h.authHelper).FindTeacherDetailRequest(ctx, s)
		ctx.JSON(statusCode, result)
	})

	group.PUT("/:id", func(ctx *gin.Context) {
		statusCode, result := endpoint.NewEndpoint(h.authHelper).UpdateTeacherRequest(ctx, s)
		ctx.JSON(statusCode, result)
	})

	group.DELETE("/:id", func(ctx *gin.Context) {
		statusCode, result := endpoint.NewEndpoint(h.authHelper).DeleteTeacherRequest(ctx, s)
		ctx.JSON(statusCode, result)
	})
}
