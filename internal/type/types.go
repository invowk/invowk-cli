package _type

import (
	"fmt"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	translation "github.com/go-playground/validator/v10/translations/en"
)

type PromptText string

func (v PromptText) Check() []error {
	return translateError("", validate.Var(v, "required,min=1,max=80"), &translator)
}

func (v PromptText) Validate() error {
	errs := v.Check()
	if errs != nil {
		return errs[0]
	} else {
		return nil
	}
}

func (v PromptText) Value() PromptText {
	return v
}

type PromptInput string

func (v PromptInput) Check(tag string) error {
	return validate.Var(v, tag)
}

func (v PromptInput) Value() PromptInput {
	return v
}

type HelpText string

func (v HelpText) Check() []error {
	return translateError("", validate.Var(v, "required,min=10,max=80"), &translator)

}

func (v HelpText) Value() HelpText {
	return v
}

type PlaceholderValue string

func (v PlaceholderValue) Check() error {
	return validate.Var(v, "min=1,max=80")
}

func (v PlaceholderValue) Value() PlaceholderValue {
	return v
}

type DefaultValue string

func (v DefaultValue) Check() error {
	return validate.Var(v, "min=1,max=1024")
}

func (v DefaultValue) Value() DefaultValue {
	return v
}

type ValidationRegex string

func (v ValidationRegex) Check() error {
	return validate.Var(v, "max=1024")
}

func (v ValidationRegex) Value() ValidationRegex {
	return v
}

type ValidationJs string

func (v ValidationJs) Check() error {
	return validate.Var(v, "max=4096")
}

func (v ValidationJs) Value() ValidationJs {
	return v
}

type ValidationMinSize uint

func (v ValidationMinSize) Check() error {
	return nil
}

func (v ValidationMinSize) Value() ValidationMinSize {
	return v
}

type ValidationMaxSize uint

func (v ValidationMaxSize) Check() error {
	return nil
}

func (v ValidationMaxSize) Value() ValidationMaxSize {
	return v
}

type ValidationCanBeBlank bool

func (v ValidationCanBeBlank) Check() error {
	return nil
}

func (v ValidationCanBeBlank) Value() ValidationCanBeBlank {
	return v
}

type ValidationIsRequired bool

func (v ValidationIsRequired) Check() error {
	return nil
}

func (v ValidationIsRequired) Value() ValidationIsRequired {
	return v
}

var (
	enLocale          = en.New()
	translatorFactory = ut.New(enLocale, enLocale)
	translator, _     = translatorFactory.GetTranslator(enLocale.Locale())
	validate          = validator.New()
	_                 = translation.RegisterDefaultTranslations(validate, translator)
)

func translateError(fieldName string, err error, trans *ut.Translator) (errs []error) {
	if err == nil {
		return nil
	}
	validatorErrs := err.(validator.ValidationErrors)
	for _, e := range validatorErrs {
		if fieldName == "" {
			fieldName = e.Type().Name()
		}
		translatedErr := fmt.Errorf(fieldName + e.Translate(*trans))
		errs = append(errs, translatedErr)
	}
	return errs
}
