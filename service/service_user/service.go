package service_user

import (
	"github.com/DangerZombie/golang-gin-boilerplate/helper/auth"
	"github.com/DangerZombie/golang-gin-boilerplate/repository"
	"github.com/DangerZombie/golang-gin-boilerplate/repository/repository_user"
	"go.uber.org/zap"
)

type userServiceImpl struct {
	logger     *zap.Logger
	authHelper auth.AuthHelper
	baseRepo   repository.BaseRepository
	userRepo   repository_user.UserRepository
}

func NewUserService(
	lg *zap.Logger,
	ah auth.AuthHelper,
	br repository.BaseRepository,
	ur repository_user.UserRepository,
) UserService {
	return &userServiceImpl{lg, ah, br, ur}
}
