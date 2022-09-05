package store

import (
	"context"

	"student-management-system/models"
)

type Student interface {
	Get(ctx context.Context) ([]models.Student, error)
	GetByID(ctx context.Context, id int) (models.Student, error)
	GetByLastName(ctx context.Context, lastName string) ([]models.Student, error)
	GetByFirstName(ctx context.Context, firstName string) ([]models.Student, error)
	GetByFirstAndLastName(ctx context.Context, firstName, lastName string) ([]models.Student, error)
	Post(ctx context.Context, student *models.Student) (models.Student, error)
}
