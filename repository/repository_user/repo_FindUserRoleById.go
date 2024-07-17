package repository_user

import (
	"errors"

	"github.com/DangerZombie/golang-gin-boilerplate/model/entity"
	"github.com/DangerZombie/golang-gin-boilerplate/model/parameter"
	"gorm.io/gorm"
)

func (r *userRepo) FindUserRoleByUserId(db *gorm.DB, input parameter.FindUserRoleByUserIdInput) (output parameter.FindUserRoleByUserIdOutput, err error) {
	item := entity.User{}
	err = db.
		Model(&entity.User{}).
		Preload("Roles").
		Where(`"user".id = ?`, input.Id).
		First(&item).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return
		}

		return output, err
	}

	output = parameter.FindUserRoleByUserIdOutput(item)

	return
}
