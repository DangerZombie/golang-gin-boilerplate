package http_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DangerZombie/golang-gin-boilerplate/helper/auth"
	transport "github.com/DangerZombie/golang-gin-boilerplate/http"
	"github.com/DangerZombie/golang-gin-boilerplate/service/service_teacher"
	"github.com/gin-gonic/gin"
	"go.uber.org/mock/gomock"
)

func TestHttpTeacher(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockAuthHelper := auth.NewMockAuthHelper(mockCtrl)
	mockTeacherService := service_teacher.NewMockTeacherService(mockCtrl)

	httpModule := transport.NewHttp(
		mockAuthHelper,
	)

	g := gin.New()
	gr := g.Group("/api/v1/teacher")

	t.Run("Should access teacher http handler", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/login", nil)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Request = req

		httpModule.TeacherHandler(gr, mockTeacherService)
	})
}
