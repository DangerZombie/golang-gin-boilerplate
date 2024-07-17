package auth

import (
	"github.com/DangerZombie/golang-gin-boilerplate/repository"
	"github.com/DangerZombie/golang-gin-boilerplate/repository/repository_user"
)

type authHelperImpl struct {
	baseRepo repository.BaseRepository
	userRepo repository_user.UserRepository
}

func NewAuthHelper(br repository.BaseRepository, ur repository_user.UserRepository) AuthHelper {
	return &authHelperImpl{br, ur}
}
