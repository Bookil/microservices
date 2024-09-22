package validation

import (
	"strings"
	"sync"

	"github.com/Bookil/microservices/user/internal/application/core/domain"
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

func (v *Validator) ValidateRegisterInputs(firstName, lastName, email string) error {
	trimSpacedFirstName := strings.TrimSpace(firstName)
	trimSpacedLastName := strings.TrimSpace(lastName)
	trimSpacedEmail := strings.TrimSpace(email)

	registerInputs := &registerInputs{
		FirstName: trimSpacedFirstName,
		LastName:  trimSpacedLastName,
		Email:     trimSpacedEmail,
	}

	err := v.validate(registerInputs)

	return err
}

func (v *Validator) ValidateGetUserIDByEmailInputs(email string) error {
	trimSpacedEmail := strings.TrimSpace(email)

	getUserIDByEmailInputs := &getUserIDByEmailInputs{
		Email: trimSpacedEmail,
	}

	err := v.validate(getUserIDByEmailInputs)

	return err
}

func (v *Validator) ValidateChangePasswordInputs(UserID domain.UserID, oldPassword string, newPassword string) error {
	changePasswordInputs := &changePasswordInputs{
		UserID:      strings.TrimSpace(UserID),
		OldPassword: strings.TrimSpace(oldPassword),
		NewPassword: strings.TrimSpace(newPassword),
	}

	err := v.validate(changePasswordInputs)

	return err
}

func (v *Validator) ValidateUpdateInputs(UserID domain.UserID, firstName string, lastName string) error {
	changePasswordInputs := &updateInputs{
		UserID:    strings.TrimSpace(UserID),
		FirstName: strings.TrimSpace(firstName),
		LastName:  strings.TrimSpace(lastName),
	}

	err := v.validate(changePasswordInputs)

	return err
}

func (v *Validator) ValidateDeleteAccountInputs(UserID domain.UserID, password string) error {
	deleteAccountInputs := &deleteAccountInputs{
		UserID:   strings.TrimSpace(UserID),
		Password: strings.TrimSpace(password),
	}

	err := v.validate(deleteAccountInputs)

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
