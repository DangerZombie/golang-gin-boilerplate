package http

import (
	"github.com/DangerZombie/golang-gin-boilerplate/endpoint"
	"github.com/DangerZombie/golang-gin-boilerplate/service/service_user"
	"github.com/gin-gonic/gin"
)

func (h *httpImpl) UserHandler(group *gin.RouterGroup, s service_user.UserService) {
	group.POST("/login", func(ctx *gin.Context) {
		statusCode, result := endpoint.NewEndpoint(h.authHelper).LoginRequest(ctx, s)
		ctx.JSON(statusCode, result)
	})

	group.GET("/profile", func(ctx *gin.Context) {
		statusCode, result := endpoint.NewEndpoint(h.authHelper).UserProfileRequest(ctx, s)
		ctx.JSON(statusCode, result)
	})

	group.POST("/register", func(ctx *gin.Context) {
		statusCode, result := endpoint.NewEndpoint(h.authHelper).RegisterUserRequest(ctx, s)
		ctx.JSON(statusCode, result)
	})
}
