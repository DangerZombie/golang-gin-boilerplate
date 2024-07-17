package service_user_test

import (
	"errors"
	"testing"

	"github.com/DangerZombie/golang-gin-boilerplate/helper/auth"
	"github.com/DangerZombie/golang-gin-boilerplate/initialization"
	"github.com/DangerZombie/golang-gin-boilerplate/model/base"
	"github.com/DangerZombie/golang-gin-boilerplate/model/entity"
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

func TestRegisterUser(t *testing.T) {
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
	issuerId := faker.UUIDHyphenated()
	username := faker.Username()
	password := faker.Password()
	nickname := faker.Name()
	registerUserRequest := request.RegisterUserRequestBody{
		Username: username,
		Password: password,
		Nickname: nickname,
		Issuer:   issuerId,
	}

	createUserInput := parameter.CreateUserInput{
		User: entity.User{
			Username: registerUserRequest.Username,
			Password: registerUserRequest.Password,
			Nickname: registerUserRequest.Nickname,
			Status:   "ACTIVE",
			BaseModel: base.BaseModel{
				CreatedBy: issuerId,
				UpdatedBy: issuerId,
			},
		},
	}

	createUserOutput := parameter.CreateUserOutput{
		Id: id,
	}

	t.Run("Should return id", func(t *testing.T) {
		mockBaseRepository.EXPECT().
			GetBegin().
			Times(1).
			Return(nil)

		mockUserRepository.EXPECT().
			CreateUser(gomock.Any(), createUserInput).
			Times(1).
			Return(createUserOutput, nil)

		mockBaseRepository.EXPECT().
			BeginCommit(gomock.Any()).
			Times(1).
			Return()

		result, message, err := userService.RegisterUser(registerUserRequest)

		require.Equal(t, id, result.Id)
		require.NotEmpty(t, message)
		require.Nil(t, err)
	})

	t.Run("Should return error", func(t *testing.T) {
		mockBaseRepository.EXPECT().
			GetBegin().
			Times(1).
			Return(nil)

		mockUserRepository.EXPECT().
			CreateUser(gomock.Any(), createUserInput).
			Times(1).
			Return(parameter.CreateUserOutput{}, errors.New("failed"))

		mockBaseRepository.EXPECT().
			BeginRollback(gomock.Any()).
			Times(1).
			Return()

		result, message, err := userService.RegisterUser(registerUserRequest)

		require.Empty(t, result)
		require.NotEmpty(t, message)
		require.NotEmpty(t, err)
	})
}
