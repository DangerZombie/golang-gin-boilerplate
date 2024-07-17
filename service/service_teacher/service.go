package service_teacher

import (
	"github.com/DangerZombie/golang-gin-boilerplate/repository"
	"github.com/DangerZombie/golang-gin-boilerplate/repository/repository_teacher"
	"go.uber.org/zap"
)

type teacherServiceImpl struct {
	logger      *zap.Logger
	baseRepo    repository.BaseRepository
	teacherRepo repository_teacher.TeacherRepository
}

func NewTeacherService(
	lg *zap.Logger,
	br repository.BaseRepository,
	tr repository_teacher.TeacherRepository,
) TeacherService {
	return &teacherServiceImpl{lg, br, tr}
}
