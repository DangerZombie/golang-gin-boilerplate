package repository_teacher

import (
	"errors"
	"github.com/DangerZombie/golang-gin-boilerplate/model/entity"
	"github.com/DangerZombie/golang-gin-boilerplate/model/parameter"

	"gorm.io/gorm"
)

func (r *teacherRepo) FindTeacherById(db *gorm.DB, input parameter.FindTeacherByIdInput) (output parameter.FindTeacherByIdOutput, err error) {
	var teacher entity.Teacher
	err = db.
		Model(&entity.Teacher{}).
		Preload("JobTitle").
		Preload("User").
		Where("teacher.id = ?", input.Id).
		First(&teacher).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return output, nil
		}

		return output, err
	}

	output = parameter.FindTeacherByIdOutput{
		Id:           teacher.Id,
		Nickname:     teacher.User.Nickname,
		Email:        teacher.User.Username,
		Status:       teacher.Status,
		Experience:   teacher.Experience,
		Degree:       teacher.Degree,
		JobTitleId:   teacher.JobTitleId,
		JobTitleName: teacher.JobTitle.Name,
	}

	return output, nil
}
