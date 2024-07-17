package repository_teacher

import (
	"github.com/DangerZombie/golang-gin-boilerplate/model/entity"
	"github.com/DangerZombie/golang-gin-boilerplate/model/parameter"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (r *teacherRepo) UpdateTeacherById(db *gorm.DB, input parameter.UpdateTeacherByIdInput) (output parameter.UpdateTeacherByIdOutput, err error) {
	var teacher entity.Teacher
	err = db.Debug().
		Model(&teacher).
		Where("id = ?", input.Id).
		Clauses(clause.Returning{}).
		Updates(input.Fields).
		Error

	if err != nil {
		return output, err
	}

	output = parameter.UpdateTeacherByIdOutput{
		Teacher: teacher,
	}

	return output, nil
}
