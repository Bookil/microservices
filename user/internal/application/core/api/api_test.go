package api

import (
	"testing"

	"github.com/Bookil/microservices/user/mocks"
	"github.com/stretchr/testify/suite"
)

type ApplicationTestSuit struct {
	suite.Suite
	api *Application
}

func TestApplicationTestSuit(t *testing.T) {
	suite.Run(t, new(ApplicationTestSuit))
}

func (a *ApplicationTestSuit) SetupSuite() {
	mockedAuth := &mocks.MockedAuth{}

	mockedDB := &mocks.MockedDB{}

	mockedEmail := &mocks.MockedEmail{}

	api := NewApplication(mockedAuth, mockedEmail, mockedDB)

	a.api = api
}