package validation

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"strconv"
)

type Service struct{}

func NewService() *Service {
	service := &Service{}
	service.initSomeValidation()
	return service
}

type ApiError struct {
	Field string
	Msg   string
}

type HasErrorsMessage interface {
	ValidationErrorMessage(tag string) string
}

func (s *Service) GetErrors(err error, model HasErrorsMessage) []ApiError {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]ApiError, len(ve))
		for i, fe := range ve {
			out[i] = ApiError{fe.Field(), model.ValidationErrorMessage(fe.Tag())}
		}
		return out
	}
	return nil
}

func (s *Service) ValidateIdParam(value string, paramName string) (int, error) {
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("%s parameter must be an integer", paramName)
	}
	if intValue <= 0 {
		return 0, fmt.Errorf("%s parameter must be an integer greater than 0", paramName)
	}
	return intValue, nil
}

var isCorrectUrl validator.Func = func(fl validator.FieldLevel) bool {
	//todo validate for correct url,
	return true
}

func (s *Service) initSomeValidation() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("correctUrl", isCorrectUrl)
	}
}
