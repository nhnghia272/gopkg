package gopkg

import (
	"errors"
	"fmt"
	"reflect"
	"slices"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validate *validator.Validate
	tags     map[string][]string
}

func NewValidator(tags map[string][]string) *Validator {
	validate := validator.New(validator.WithRequiredStructEnabled())
	validate.RegisterTagNameFunc(func(e reflect.StructField) string {
		return strings.Split(e.Tag.Get("json"), ",")[0]
	})
	for tag, values := range tags {
		validate.RegisterValidation(tag, func(e validator.FieldLevel) bool {
			return slices.Contains(values, e.Field().String())
		})
	}
	return &Validator{validate, tags}
}

func (s *Validator) RegisterValidation(tag string, fn func(val string) bool) {
	s.validate.RegisterValidation(tag, func(e validator.FieldLevel) bool {
		return fn(e.Field().String())
	})
}

func (s *Validator) Validate(data any) error {
	if err := s.validate.Struct(data); err != nil {
		results := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			msgs := []string{fmt.Sprintf("(%v)", e.Kind().String()), e.Tag()}
			if values, ok := s.tags[e.Tag()]; ok {
				msgs[1] = fmt.Sprintf("must be one of [%v]", strings.Join(values, ","))
			}
			if e.Param() != "" {
				msgs = append(msgs, e.Param())
			}
			results = append(results, fmt.Sprintf("[%v]:", e.Field())+strings.Join(msgs, " "))
		}
		return errors.New(strings.Join(results, "; "))
	}
	return nil
}
