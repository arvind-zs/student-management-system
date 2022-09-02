package service

import (
	"context"

	"student-management-system/models"
)

type Student interface {
	Post(ctx context.Context, student models.Student) (models.Student, error)
}
