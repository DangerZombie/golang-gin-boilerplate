package service_user_test

import (
	"errors"
	"testing"

	"github.com/DangerZombie/golang-gin-boilerplate/helper/auth"
	"github.com/DangerZombie/golang-gin-boilerplate/helper/message"
	"github.com/DangerZombie/golang-gin-boilerplate/initialization"
	"github.com/DangerZombie/golang-gin-boilerplate/model/parameter"
	"github.com/DangerZombie/golang-gin-boilerplate/model/request"
	"github.com/DangerZombie/golang-gin-boilerplate/repository"
	"github.com/DangerZombie/golang-gin-boilerplate/repository/repository_user"
	"github.com/DangerZombie/golang-gin-boilerplate/service/service_user"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
)

func TestUserProfile(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockAuthHelper := auth.NewMockAuthHelper(mockCtrl)
	mockUserRepository := repository_user.NewMockUserRepository(mockCtrl)
	mockBaseRepository := repository.NewMockBaseRepository(mockCtrl)

	var logger *zap.Logger
	logger, _ = initialization.NewZapLogger("")
	userService := service_user.NewUserService(
		logger,
		mockAuthHelper,
		mockBaseRepository,
		mockUserRepository)

	id := faker.UUIDHyphenated()
	userProfileRequest := request.UserProfileRequest{
		Id: id,
	}

	userProfileEmptyRequest := request.UserProfileRequest{
		Id: "",
	}

	findUserByIdInput := parameter.FindUserByIdInput{
		Id: id,
	}

	findUserByIdOutput := parameter.FindUserByIdOutput{
		Id:       id,
		Nickname: faker.Name(),
	}

	t.Run("Should return user profile", func(t *testing.T) {
		mockBaseRepository.EXPECT().
			GetBegin().
			Times(1).
			Return(nil)

		mockUserRepository.EXPECT().
			FindUserById(gomock.Any(), findUserByIdInput).
			Times(1).
			Return(findUserByIdOutput, nil)

		mockBaseRepository.EXPECT().
			BeginCommit(gomock.Any()).
			Times(1).
			Return()

		result, msg, err := userService.UserProfile(userProfileRequest)

		require.NotEmpty(t, result)
		require.Equal(t, message.SuccessMsg, msg)
		require.Nil(t, err)
	})

	t.Run("Should return error if id is empty", func(t *testing.T) {
		result, msg, err := userService.UserProfile(userProfileEmptyRequest)

		require.Empty(t, result)
		require.Equal(t, message.FailedMsg, msg)
		require.NotEmpty(t, err)
	})

	t.Run("Should return error if failed to fetch user", func(t *testing.T) {
		mockBaseRepository.EXPECT().
			GetBegin().
			Times(1).
			Return(nil)

		mockUserRepository.EXPECT().
			FindUserById(gomock.Any(), findUserByIdInput).
			Times(1).
			Return(parameter.FindUserByIdOutput{}, errors.New("failed"))

		mockBaseRepository.EXPECT().
			BeginRollback(gomock.Any()).
			Times(1).
			Return()

		result, msg, err := userService.UserProfile(userProfileRequest)

		require.Empty(t, result)
		require.Equal(t, message.FailedMsg, msg)
		require.NotEmpty(t, err)
	})
}
