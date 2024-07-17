package service_teacher_test

import (
	"errors"
	"testing"

	"github.com/DangerZombie/golang-gin-boilerplate/helper/message"
	"github.com/DangerZombie/golang-gin-boilerplate/initialization"
	"github.com/DangerZombie/golang-gin-boilerplate/model/parameter"
	"github.com/DangerZombie/golang-gin-boilerplate/model/request"
	"github.com/DangerZombie/golang-gin-boilerplate/repository"
	"github.com/DangerZombie/golang-gin-boilerplate/repository/repository_teacher"
	"github.com/DangerZombie/golang-gin-boilerplate/service/service_teacher"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
)

func TestTeacherDelete(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockBaseRepository := repository.NewMockBaseRepository(mockCtrl)
	mockTeacherRepository := repository_teacher.NewMockTeacherRepository(mockCtrl)

	var logger *zap.Logger
	logger, _ = initialization.NewZapLogger("")
	teacherService := service_teacher.NewTeacherService(
		logger,
		mockBaseRepository,
		mockTeacherRepository)

	id := faker.UUIDHyphenated()

	deleteTeacherRequest := request.TeacherDeleteRequest{
		Id: id,
	}

	deleteTeacherEmptyRequest := request.TeacherDeleteRequest{
		Id: "",
	}

	deleteTeacherInput := parameter.DeleteTeacherByIdInput{
		Id: id,
	}

	deleteTeacherOutput := parameter.DeleteTeacherByIdOutput{
		Success: true,
	}

	t.Run("Should return success", func(t *testing.T) {
		mockBaseRepository.EXPECT().
			GetBegin().
			Times(1).
			Return(nil)

		mockTeacherRepository.EXPECT().
			DeleteTeacherById(gomock.Any(), deleteTeacherInput).
			Times(1).
			Return(deleteTeacherOutput, nil)

		mockBaseRepository.EXPECT().
			BeginCommit(gomock.Any()).
			Times(1).
			Return()

		result, msg, errMsg := teacherService.TeacherDelete(deleteTeacherRequest)

		require.Equal(t, true, result.Success)
		require.Equal(t, message.SuccessMsg, msg)
		require.Empty(t, errMsg)
	})

	t.Run("Should return error if id is empty", func(t *testing.T) {
		result, msg, errMsg := teacherService.TeacherDelete(deleteTeacherEmptyRequest)

		require.Equal(t, false, result.Success)
		require.Equal(t, message.FailedMsg, msg)
		require.NotEmpty(t, errMsg)
	})

	t.Run("Should return error if failed to delete teacher", func(t *testing.T) {
		mockBaseRepository.EXPECT().
			GetBegin().
			Times(1).
			Return(nil)

		mockTeacherRepository.EXPECT().
			DeleteTeacherById(gomock.Any(), deleteTeacherInput).
			Times(1).
			Return(parameter.DeleteTeacherByIdOutput{}, errors.New("failed"))

		mockBaseRepository.EXPECT().
			BeginRollback(gomock.Any()).
			Times(1).
			Return()

		result, msg, errMsg := teacherService.TeacherDelete(deleteTeacherRequest)

		require.Equal(t, false, result.Success)
		require.Equal(t, message.FailedMsg, msg)
		require.NotEmpty(t, errMsg)
	})
}
