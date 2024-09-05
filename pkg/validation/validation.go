package validation

import (
	"encoding/json"
	"errors"
	"reflect"
	"strings"
	"tel/product/pkg/exception"

	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validator *validator.Validate
}

func NewValidator() *Validator {
	v := validator.New()

	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	return &Validator{
		validator: v,
	}
}

func (v *Validator) Validate(i any) error {
	if err := v.validator.Struct(i); err != nil {
		var validationErrors validator.ValidationErrors
		var messages []map[string]interface{}

		if errors.As(err, &validationErrors) {
			for _, ve := range validationErrors {
				fieldName := ve.StructField()
				messages = append(messages, map[string]any{
					"field":   fieldName,
					"message": msgForTag(ve.Tag()),
				})
			}

			msg, err := json.Marshal(messages)
			if err != nil {
				return err
			}

			return exception.NewValidatonError(string(msg), validationErrors)
		}

		return err
	}

	return nil
}

func msgForTag(tag string) string {
	switch tag {
	case "required":
		return "This field is required"
	}
	return ""
}
