package validators

import (
	"errors"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"strings"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
)

func New(data interface{}) error {

	en := en.New()
	uni = ut.New(en, en)

	trans, _ := uni.GetTranslator("en")

	validate = validator.New()
	en_translations.RegisterDefaultTranslations(validate, trans)

	err := Translate(trans, data)
	if err != nil {
		return err
	}
	return nil
}

func Translate(trans ut.Translator, data interface{}) error {

	validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} must have a value!", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())

		return t
	})
	validate.RegisterTranslation("email", trans, func(ut ut.Translator) error {
		return ut.Add("email", "{0} must be a valid email", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("email", fe.Field())
		return t
	})
	err := validate.Struct(data)
	if err != nil {

		// translate all error at once
		errs := err.(validator.ValidationErrors)
		results := errs.Translate(trans)
		var str strings.Builder

		for _, val := range results {
			str.WriteString(strings.ToLower(val) + " ")
		}
		err = errors.New(str.String())
		return err
	}

	return nil

}
