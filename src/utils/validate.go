package utils

import (
	"github.com/go-playground/validator/v10"
)

func ValidateData[T any](data T) error {
	v := validator.New()
	err := v.Struct(data)
	if err != nil {
		return err
	}
	return nil
}
