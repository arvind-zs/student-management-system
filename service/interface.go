package service

import (
	"context"

	"student-management-system/models"
)

type Student interface {
	Get(ctx context.Context, firstName, lastName string) ([]models.Student, error)
	GetByID(ctx context.Context, id int) (models.Student, error)
	Post(ctx context.Context, student *models.Student) (models.Student, error)
}
