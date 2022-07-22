package services

import (
	"github.com/alekseyklimenko/go-proj-bootstrap/models"
	"github.com/alekseyklimenko/go-proj-bootstrap/models/requests"
	"github.com/alekseyklimenko/go-proj-bootstrap/services/validation"
)

var (
	Some       SomeService
	Validation ValidationService
	Processing ProcessingService
)

type SomeService interface {
	CreateNew(item *models.Some, formData requests.Some) error
	GetSomeToProcess(count uint) *[]models.Some
	LockItem(id uint) error
	FreeItem(id uint)
}

type ValidationService interface {
	GetErrors(err error, model validation.HasErrorsMessage) []validation.ApiError
	ValidateIdParam(value string, paramName string) (int, error)
}

type ProcessingService interface {
	QueueItem(some models.Some)
	Shutdown()
}
