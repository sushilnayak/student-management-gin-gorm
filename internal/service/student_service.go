package service

import (
	"context"
	"student-service/internal/model"
	"student-service/internal/repository"
)

type StudentService interface {
	CreateStudent(ctx context.Context, student *model.Student) error
	GetStudent(ctx context.Context, id uint) (*model.Student, error)
	UpdateStudent(ctx context.Context, student *model.Student) error
	DeleteStudent(ctx context.Context, id uint) error
	ListStudents(ctx context.Context) ([]model.Student, error)
}

type studentService struct {
	repo repository.StudentRepository
}

func NewStudentService(repo repository.StudentRepository) StudentService {
	return &studentService{repo: repo}
}

func (s *studentService) CreateStudent(ctx context.Context, student *model.Student) error {
	return s.repo.Create(ctx, student)
}

func (s *studentService) GetStudent(ctx context.Context, id uint) (*model.Student, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *studentService) UpdateStudent(ctx context.Context, student *model.Student) error {
	return s.repo.Update(ctx, student)
}

func (s *studentService) DeleteStudent(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}

func (s *studentService) ListStudents(ctx context.Context) ([]model.Student, error) {
	return s.repo.List(ctx)
}
