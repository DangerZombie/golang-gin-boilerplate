package repository_user

import (
	"errors"

	"github.com/DangerZombie/golang-gin-boilerplate/model/entity"
	"github.com/DangerZombie/golang-gin-boilerplate/model/parameter"
	"gorm.io/gorm"
)

func (r *userRepo) FindUserByUsernameAndPassword(db *gorm.DB, input parameter.FindUserByUsernameAndPasswordInput) (output parameter.FindUserByUsernameAndPasswordOutput, err error) {
	err = db.
		Model(&entity.User{}).
		Where("username = ? AND password = ?", input.Username, input.Password).
		First(&output).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return
		}

		return output, err
	}

	return
}
