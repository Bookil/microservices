package validation

import (
	"strings"
	"sync"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	locale_EN         = en.New()
	uni               = ut.New(locale_EN, locale_EN)
	ValiTranslator, _ = uni.GetTranslator("en")
)

var (
	valiLock     = &sync.Mutex{}
	valiInstance *Validator
)

type Validator struct {
	vi *validator.Validate
}

func NewValidator() *Validator {
	if valiInstance == nil {
		valiLock.Lock()
		defer valiLock.Unlock()

		vi := validator.New()

		valiInstance = &Validator{vi}
		err := valiInstance.translateOverride()
		if err != nil {
			panic(err)
		}
	}
	return valiInstance
}

func (v *Validator) ValidateGetBooksByTitle(title string) error {
	trimSpacedTitle := strings.TrimSpace(title)

	registerInputs := &getBooksByTitle{
		Title: trimSpacedTitle,
	}

	err := v.validate(registerInputs)

	return err
}

func (v *Validator) ValidateGetBooksByAuthor(author string) error {
	trimSpacedAuthor := strings.TrimSpace(author)

	registerInputs := &getBooksByAuthor{
		Author: trimSpacedAuthor,
	}

	err := v.validate(registerInputs)

	return err
}

func (v *Validator) ValidateGetBooksByGenre(genre string) error {
	trimSpacedGenre := strings.TrimSpace(genre)

	registerInputs := &getBooksByGenre{
		Genre: trimSpacedGenre,
	}

	err := v.validate(registerInputs)

	return err
}

func (v *Validator) ValidateAddBook(title string, description string, price float32, quantity uint, year uint) error {
	return v.validateBookData(title, description, price, quantity, year)
}

func (v *Validator) ValidateModifyBookByID(title string, description string, price float32, quantity uint, year uint) error {
	return v.validateBookData(title, description, price, quantity, year)
}

func (v *Validator) validateBookData(title string, description string, price float32, quantity uint, year uint) error {
	trimSpacedTitle := strings.TrimSpace(title)
	trimSpacedDescription := strings.TrimSpace(description)

	registerInputs := &bookInput{
		Title:       trimSpacedTitle,
		Description: trimSpacedDescription,
		Price:       float32(price),
		Year:        year,
		Quantity:    quantity,
	}

	err := v.validate(registerInputs)

	return err
}

func (v *Validator) ValidateAddAuthor(name, about string) error {
	trimSpacedName := strings.TrimSpace(name)
	trimSpacedAbout := strings.TrimSpace(about)

	registerInputs := &addAuthor{
		Name:  trimSpacedName,
		About: trimSpacedAbout,
	}

	err := v.validate(registerInputs)

	return err
}

func (v *Validator) ValidateAddGenre(name string) error {
	trimSpacedName := strings.TrimSpace(name)

	registerInputs := &addGenre{
		Name: trimSpacedName,
	}

	err := v.validate(registerInputs)

	return err
}

func (v *Validator) validate(s interface{}) error {
	err := v.vi.Struct(s)
	return err
}

func (v *Validator) translateOverride() error {
	err := en_translations.RegisterDefaultTranslations(v.vi, ValiTranslator)
	return err
}
