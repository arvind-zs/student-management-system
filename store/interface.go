package store

import (
	"context"

	"student-management-system/models"
)

type Student interface {
	Get(ctx context.Context) ([]models.Student, error)
	Post(ctx context.Context, student models.Student) (models.Student, error)
}
