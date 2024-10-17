package validation

import (
	"strings"
	"sync"

	"product/internal/application/core/domain"

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

func (v *Validator) ValidateAddBook(book *domain.Book) error {
	return v.validateBookData(book)
}

func (v *Validator) ValidateModifyBookByID(book *domain.Book) error {
	return v.validateBookData(book)
}

func (v *Validator) validateBookData(book *domain.Book) error {
	trimSpacedTitle := strings.TrimSpace(book.Title)
	trimSpacedDescription := strings.TrimSpace(book.Description)

	registerInputs := &bookInput{
		Title:       trimSpacedTitle,
		Description: trimSpacedDescription,
		Price:       float32(book.Price),
		Year:        book.Year,
		Quantity:    book.Quantity,
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
