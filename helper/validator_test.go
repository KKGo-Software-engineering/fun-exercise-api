package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/go-playground/validator/v10"
)

type InvalidStruct struct {
	InvalidField int `validate:"required"`
}

func TestCustomValidator(t *testing.T) {
	t.Run("InvalidStruct: Failed", func(t *testing.T) {
		customValidator := CustomValidator{validator: validator.New()}

		err := customValidator.Validate(struct{}{})
		assert.NoError(t, err)

		invalidStruct := InvalidStruct{}
		err = customValidator.Validate(invalidStruct)
		assert.Error(t, err)
	})

	t.Run("InvalidStruct: Success", func(t *testing.T) {
		customValidator := CustomValidator{validator: validator.New()}

		err := customValidator.Validate(struct{}{})
		assert.NoError(t, err)

		invalidStruct := InvalidStruct{
			InvalidField: 1,
		}

		err = customValidator.Validate(invalidStruct)
		assert.Nil(t, err)
	})
}
