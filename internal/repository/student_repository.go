package repository

import (
	"context"
	"gorm.io/gorm"
	"student-service/internal/model"
)

type StudentRepository interface {
	Create(ctx context.Context, student *model.Student) error
	GetByID(ctx context.Context, id uint) (*model.Student, error)
	Update(ctx context.Context, student *model.Student) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context) ([]model.Student, error)
}

type studentRepository struct {
	db *gorm.DB
}

func NewStudentRepository(db *gorm.DB) StudentRepository {
	return &studentRepository{db: db}
}

func (r *studentRepository) Create(ctx context.Context, student *model.Student) error {
	return r.db.WithContext(ctx).Create(student).Error
}

func (r *studentRepository) GetByID(ctx context.Context, id uint) (*model.Student, error) {
	var student model.Student
	if err := r.db.WithContext(ctx).First(&student, id).Error; err != nil {
		return nil, err
	}
	return &student, nil
}

func (r *studentRepository) Update(ctx context.Context, student *model.Student) error {
	return r.db.WithContext(ctx).Save(student).Error
}

func (r *studentRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.Student{}, id).Error
}

func (r *studentRepository) List(ctx context.Context) ([]model.Student, error) {
	var students []model.Student
	if err := r.db.WithContext(ctx).Find(&students).Error; err != nil {
		return nil, err
	}
	return students, nil
}
