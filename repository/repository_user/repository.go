package repository_user

import "github.com/DangerZombie/golang-gin-boilerplate/repository"

type userRepo struct {
	base repository.BaseRepository
}

func NewUserRepository(br repository.BaseRepository) UserRepository {
	return &userRepo{br}
}
