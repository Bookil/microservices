package validation_test

import (
	"log"
	"testing"

	"product/internal/adapters/validation"
	"product/internal/application/core/domain"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type VerificationTestSuite struct {
	suite.Suite
	validator *validation.Validator
}

func TestVerificationTestSuite(t *testing.T) {
	suite.Run(t, new(VerificationTestSuite))
}

func (v *VerificationTestSuite) SetupSuite() {
	v.validator = validation.NewValidator()
}

func (v *VerificationTestSuite) TestValidateGetBooksByTitle() {
	testCases := []struct {
		title string
		valid bool
	}{
		{
			title: " ",
			valid: false,
		},
		{
			title: "valid",
			valid: true,
		},
	}

	validator := validation.NewValidator()

	for n, tc := range testCases {
		n++
		log.Printf("number %d,case:%v\n", n, tc)
		err := validator.ValidateGetBooksByTitle(tc.title)

		if tc.valid {
			require.NoError(v.T(), err)
		} else {
			require.Error(v.T(), err)
		}

	}
}

func (v *VerificationTestSuite) TestValidateGetBooksByAuthor() {
	testCases := []struct {
		Genre string
		valid bool
	}{
		{
			Genre: " ",
			valid: false,
		},
		{
			Genre: "valid",
			valid: true,
		},
	}

	validator := validation.NewValidator()

	for n, tc := range testCases {
		n++
		log.Printf("number %d,case:%v\n", n, tc)
		err := validator.ValidateGetBooksByAuthor(tc.Genre)

		if tc.valid {
			require.NoError(v.T(), err)
		} else {
			require.Error(v.T(), err)
		}

	}
}

func (v *VerificationTestSuite) TestValidateGetBooksByGenre() {
	testCases := []struct {
		genre string
		valid bool
	}{
		{
			genre: " ",
			valid: false,
		},
		{
			genre: "valid",
			valid: true,
		},
	}

	validator := validation.NewValidator()

	for n, tc := range testCases {
		n++
		log.Printf("number %d,case:%v\n", n, tc)
		err := validator.ValidateGetBooksByGenre(tc.genre)

		if tc.valid {
			require.NoError(v.T(), err)
		} else {
			require.Error(v.T(), err)
		}

	}
}

func (v *VerificationTestSuite) TestValidateAddBook() {
	testCases := []struct {
		book  *domain.Book
		valid bool
	}{
		{
			book:  domain.NewBook("  ", "valid", nil, 12, 1900, nil, 25),
			valid: false,
		},
		{
			book:  domain.NewBook("valid", " ", nil, 12, 1900, nil, 25),
			valid: false,
		},
		{
			book:  domain.NewBook("valid", "valid", nil, 0, 1900, nil, 25),
			valid: false,
		},
		{
			book:  domain.NewBook("valid", "valid", nil, 12, 0, nil, 25),
			valid: false,
		},
		{
			book:  domain.NewBook("valid", "valid", nil, 12, 1900, nil, 0),
			valid: false,
		},
		{
			book:  domain.NewBook("valid", "valid", nil, 12, 1900, nil, 25),
			valid: true,
		},
	}

	validator := validation.NewValidator()

	for n, tc := range testCases {
		n++
		log.Printf("number %d,case:%v\n", n, tc)
		err := validator.ValidateAddBook(tc.book)

		if tc.valid {
			require.NoError(v.T(), err)
		} else {
			require.Error(v.T(), err)
		}

	}
}

func (v *VerificationTestSuite) TestValidateModifyBook() {
	testCases := []struct {
		book  *domain.Book
		valid bool
	}{
		{
			book:  domain.NewBook("  ", "valid", nil, 12, 1900, nil, 25),
			valid: false,
		},
		{
			book:  domain.NewBook("valid", " ", nil, 12, 1900, nil, 25),
			valid: false,
		},
		{
			book:  domain.NewBook("valid", "valid", nil, 0, 1900, nil, 25),
			valid: false,
		},
		{
			book:  domain.NewBook("valid", "valid", nil, 12, 0, nil, 25),
			valid: false,
		},
		{
			book:  domain.NewBook("valid", "valid", nil, 12, 1900, nil, 0),
			valid: false,
		},
		{
			book:  domain.NewBook("valid", "valid", nil, 12, 1900, nil, 25),
			valid: true,
		},
	}

	validator := validation.NewValidator()

	for n, tc := range testCases {
		n++
		log.Printf("number %d,case:%v\n", n, tc)
		err := validator.ValidateModifyBookByID(tc.book)

		if tc.valid {
			require.NoError(v.T(), err)
		} else {
			require.Error(v.T(), err)
		}

	}
}

func (v *VerificationTestSuite) TestValidateAddAuthor() {
	testCases := []struct {
		name  string
		about string
		valid bool
	}{
		{
			name:  " ",
			about: "valid",
			valid: false,
		},
		{
			name:  "valid",
			about: " ",
			valid: false,
		},
		{
			name:  "valid",
			about: "valid",
			valid: true,
		},
	}

	validator := validation.NewValidator()

	for n, tc := range testCases {
		n++
		log.Printf("number %d,case:%v\n", n, tc)
		err := validator.ValidateAddAuthor(tc.name, tc.about)

		if tc.valid {
			require.NoError(v.T(), err)
		} else {
			require.Error(v.T(), err)
		}

	}
}

func (v *VerificationTestSuite) TestValidateAddGenre() {
	testCases := []struct {
		name  string
		valid bool
	}{
		{
			name:  " ",
			valid: false,
		},
		{
			name:  "valid",
			valid: true,
		},
	}

	validator := validation.NewValidator()

	for n, tc := range testCases {
		n++
		log.Printf("number %d,case:%v\n", n, tc)
		err := validator.ValidateAddGenre(tc.name)

		if tc.valid {
			require.NoError(v.T(), err)
		} else {
			require.Error(v.T(), err)
		}

	}
}
