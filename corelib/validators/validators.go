package validators

import (
	"errors"
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	entranslations "github.com/go-playground/validator/v10/translations/en"
)

type requestValidator struct {
	validator *validator.Validate
	trans     ut.Translator
}

func NewRequestValidator() *requestValidator {

	v := validator.New()

	en := en.New()
	uni := ut.New(en, en)
	translator, found := uni.GetTranslator("en")
	if !found {
		panic("translator not found")
	}

	if err := entranslations.RegisterDefaultTranslations(v, translator); err != nil {
		panic(err)
	}

	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		tagValue := fld.Tag.Get("pg") // search for pg tag names first
		if len(tagValue) == 0 {
			tagValue = fld.Tag.Get("json")
		}
		name := strings.SplitN(tagValue, ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	return &requestValidator{
		validator: v,
		trans:     translator,
	}
}

func (rv *requestValidator) Validate(i interface{}) error {
	if err := rv.validator.Struct(i); err != nil {
		messages := make([]string, 0)
		for _, e := range err.(validator.ValidationErrors) {
			messages = append(messages, e.Translate(rv.trans))
		}
		errMessage := strings.Join(messages, ", ")
		return errors.New(errMessage)
	}

	return nil
}
