package helper

import (
	"fmt"
	"strings"

	"github.com/KKGo-Software-engineering/fun-exercise-api/errortype"
	"github.com/go-playground/validator/v10"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(req interface{}) error {
	if err := cv.validator.Struct(req); err != nil {
		var message []string
		for _, e := range err.(validator.ValidationErrors) {
			errorMsg := fmt.Sprintf("Field %s is required", e.Field())
			message = append(message, errorMsg)
		}
		return errortype.ValidationError{Message: strings.Join(message, ", ")}
	}

	return nil
}

func NewValidator() *CustomValidator {
	return &CustomValidator{validator: validator.New()}
}
