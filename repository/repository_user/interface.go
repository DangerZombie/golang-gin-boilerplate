package repository_user

import (
	"github.com/DangerZombie/golang-gin-boilerplate/model/parameter"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(db *gorm.DB, input parameter.CreateUserInput) (output parameter.CreateUserOutput, err error)
	FindUserById(db *gorm.DB, input parameter.FindUserByIdInput) (output parameter.FindUserByIdOutput, err error)
	FindUserByUsernameAndPassword(db *gorm.DB, input parameter.FindUserByUsernameAndPasswordInput) (output parameter.FindUserByUsernameAndPasswordOutput, err error)
	FindUserRoleByUserId(db *gorm.DB, input parameter.FindUserRoleByUserIdInput) (output parameter.FindUserRoleByUserIdOutput, err error)
}
