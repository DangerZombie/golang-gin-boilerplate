package http_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DangerZombie/golang-gin-boilerplate/helper/auth"
	transport "github.com/DangerZombie/golang-gin-boilerplate/http"
	"github.com/DangerZombie/golang-gin-boilerplate/service/service_user"
	"github.com/gin-gonic/gin"
	"go.uber.org/mock/gomock"
)

func TestHttpUser(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockAuthHelper := auth.NewMockAuthHelper(mockCtrl)
	mockUserService := service_user.NewMockUserService(mockCtrl)

	httpModule := transport.NewHttp(
		mockAuthHelper,
	)

	g := gin.New()
	gr := g.Group("/api/v1/user")

	t.Run("Should access user http handler", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/login", nil)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Request = req

		httpModule.UserHandler(gr, mockUserService)
	})
}
