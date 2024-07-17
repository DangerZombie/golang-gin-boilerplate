package http

import (
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/runtime/middleware"
)

func (h *httpImpl) SwaggerHttpHandler(r *gin.Engine) {
	// serve swagger spec file
	r.GET("/swagger.yaml", func(c *gin.Context) {
		c.File("swagger.yaml")
	})

	// serve swagger UI
	opts := middleware.SwaggerUIOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.SwaggerUI(opts, nil)
	r.GET("/docs/", gin.WrapH(sh))

	// serve Redoc
	opts1 := middleware.RedocOpts{SpecURL: "/swagger.yaml", Path: "doc"}
	sh1 := middleware.Redoc(opts1, nil)
	r.GET("/doc/", gin.WrapH(sh1))
}
