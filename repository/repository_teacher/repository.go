package repository_teacher

import "github.com/DangerZombie/golang-gin-boilerplate/repository"

type teacherRepo struct {
	base repository.BaseRepository
}

func NewTeacherRepository(br repository.BaseRepository) TeacherRepository {
	return &teacherRepo{br}
}
