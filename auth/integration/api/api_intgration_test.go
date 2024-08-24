package integration

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/Bookil/microservices/auth/config"
	"github.com/Bookil/microservices/auth/internal/adapters/db"
	"github.com/Bookil/microservices/auth/internal/adapters/db/mysql_adapter"
	"github.com/Bookil/microservices/auth/internal/adapters/hash"
	"github.com/Bookil/microservices/auth/internal/application/core/api"
	"github.com/Bookil/microservices/auth/internal/application/core/domain"
	"github.com/docker/go-connections/nat"
	"github.com/stretchr/testify/suite"
	auth_manager "github.com/tahadostifam/go-auth-manager"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type APITestSuite struct {
	suite.Suite
	api                *api.Application
	userID             domain.UserID
	password           string
	verificationCode   string
	refreshToken       string
	accessToken        string
	resetPasswordToken string
}

func TestOAuthDatabaseTestSuite(t *testing.T) {
	suite.Run(t, new(APITestSuite))
}

func (a *APITestSuite) SetupSuite() {
	ctx := context.TODO()

	err := os.Setenv("AUTH_ENV", "test")
	if err != nil {
		log.Fatalf("Could not set the environment variable to test: %s", err)
	}

	config := config.Read()
	redisPort := fmt.Sprintf("%d/tcp", config.Redis.Port)

	redisContainerReq := testcontainers.ContainerRequest{
		Image:        "redis:latest",
		ExposedPorts: []string{redisPort},
		Env:          map[string]string{},
		WaitingFor:   wait.ForListeningPort(nat.Port(redisPort)).WithStartupTimeout(5 * time.Minute),
	}

	redisContainer, redisConnectErr := testcontainers.GenericContainer(ctx,
		testcontainers.GenericContainerRequest{
			ContainerRequest: redisContainerReq,
			Started:          true,
		})

	if redisConnectErr != nil {
		log.Fatal("Failed to start Mysql:", redisConnectErr)
	}

	redisEndpoint, err := redisContainer.Endpoint(ctx, "")
	if err != nil {
		log.Fatal(err)
	}

	redisEndPort, err := strconv.Atoi(strings.Split(redisEndpoint, ":")[1])
	if err != nil {
		log.Fatal(err)
	}

	config.Redis.Port = redisEndPort

	redisClient := db.GetRedisInstance(config.Redis)
	if err != nil {
		log.Fatal(err)
	}

	mysqlPort := fmt.Sprintf("%d/tcp", config.Mysql.Port)

	mysqlContainerReq := testcontainers.ContainerRequest{
		Image:        "mysql:latest",
		ExposedPorts: []string{mysqlPort},
		Env: map[string]string{
			"MYSQL_ROOT_PASSWORD": config.Mysql.Password,
			"MYSQL_DATABASE":      config.Mysql.DBName,
		},
		WaitingFor: wait.ForListeningPort(nat.Port(mysqlPort)).WithStartupTimeout(5 * time.Minute),
	}

	mysqlContainer, mysqlConnectErr := testcontainers.GenericContainer(ctx,
		testcontainers.GenericContainerRequest{
			ContainerRequest: mysqlContainerReq,
			Started:          true,
		})

	if mysqlConnectErr != nil {
		log.Fatal("Failed to start Mysql:", mysqlConnectErr)
	}

	mysqlEndpoint, err := mysqlContainer.Endpoint(ctx, "")
	if err != nil {
		log.Fatal(err)
	}

	mysqlEndPort, err := strconv.Atoi(strings.Split(mysqlEndpoint, ":")[1])
	if err != nil {
		log.Fatal(err)
	}

	config.Mysql.Port = mysqlEndPort

	mysqlAdapter, err := mysql_adapter.NewAdapter(&config.Mysql)
	if err != nil {
		log.Fatal(err)
	}

	authManger := auth_manager.NewAuthManager(redisClient, auth_manager.AuthManagerOpts{
		PrivateKey: config.JWT.SecretKey,
	})

	hashManger := hash.NewHashManager(hash.DefaultHashParams)

	api := api.NewApplication(mysqlAdapter, authManger, hashManger)

	a.api = api
}

func (a *APITestSuite) TestA_Register() {
	ctx := context.TODO()

	testCases := []struct {
		userID        domain.UserID
		password      string
		valid         bool
		expectedError error
	}{
		{
			userID:   "1234567",
			password: "password",
			valid:    true,
		},
		{
			userID:        "1234567",
			password:      "password",
			valid:         false,
			expectedError: api.ErrCreateAuthStore,
		},
	}

	for _, tc := range testCases {
		verificationCode, err := a.api.Register(ctx, tc.userID, tc.password)
		if tc.valid {
			a.NoError(err)
			a.verificationCode = verificationCode
			a.userID = tc.userID
			a.password = tc.password
		} else if !tc.valid {
			a.Error(err)
			if tc.expectedError != nil {
				a.Equal(tc.expectedError, err)
			}
			a.Empty(verificationCode)
		}
	}
}

func (a *APITestSuite) TestB_VerifyEmail() {
	ctx := context.TODO()

	testCases := []struct {
		userID           domain.UserID
		verificationCode string
		valid            bool
		expectedError    error
	}{
		{
			userID:           a.userID,
			verificationCode: "invalid",
			valid:            false,
			expectedError:    api.ErrVerifyEmail,
		},
		{
			userID:           "invalid",
			verificationCode: a.verificationCode,
			valid:            false,
			expectedError:    api.ErrVerifyEmail,
		},
		{
			userID:           a.userID,
			verificationCode: a.verificationCode,
			valid:            true,
		},
	}

	for _, tc := range testCases {
		err := a.api.VerifyEmail(ctx, tc.userID, tc.verificationCode)
		if tc.valid {
			a.NoError(err)
		} else if !tc.valid {
			a.Error(err)
			if tc.expectedError != nil {
				a.Equal(tc.expectedError, err)
			}
		}
	}
}

func (a *APITestSuite) TestC_Login() {
	ctx := context.TODO()

	testCases := []struct {
		userID        domain.UserID
		password      string
		valid         bool
		expectedError error
	}{
		{
			userID:        a.userID,
			password:      "invalid",
			valid:         false,
			expectedError: api.ErrInvalidEmailPassword,
		},
		{
			userID:        "invalid",
			password:      a.password,
			valid:         false,
			expectedError: api.ErrNotFound,
		},
		{
			userID:   a.userID,
			password: a.password,
			valid:    true,
		},
	}

	for _, tc := range testCases {
		accessToken, refreshToken, err := a.api.Login(ctx, tc.userID, tc.password)
		if tc.valid {
			a.NoError(err)
			a.accessToken = accessToken
			a.refreshToken = refreshToken
		} else if !tc.valid {
			a.Error(err)
			if tc.expectedError != nil {
				a.Equal(tc.expectedError, err)
			}
		}
	}
}

func (a *APITestSuite) TestD_Authentication() {
	ctx := context.TODO()

	testCases := []struct {
		accessToken   string
		valid         bool
		expectedError error
	}{
		{
			accessToken:   "invalid",
			valid:         false,
			expectedError: api.ErrAccessDenied,
		},
		{
			accessToken: a.accessToken,
			valid:       true,
		},
	}

	for _, tc := range testCases {
		userID, err := a.api.Authenticate(ctx, tc.accessToken)
		if tc.valid {
			a.NoError(err)
			a.Equal(a.userID, userID)
		} else if !tc.valid {
			a.Error(err)
			if tc.expectedError != nil {
				a.Equal(tc.expectedError, err)
			}
		}
	}
}

func (a *APITestSuite) TestE_ChangePassword() {
	ctx := context.TODO()

	testCases := []struct {
		userID        domain.UserID
		oldPassword   string
		newPassword   string
		valid         bool
		expectedError error
	}{
		{
			userID:        a.userID,
			oldPassword:   "invalid",
			newPassword:   "valid",
			valid:         false,
			expectedError: api.ErrInvalidPassword,
		},
		{
			userID:        "invalid",
			oldPassword:   a.password,
			newPassword:   "valid",
			valid:         false,
			expectedError: api.ErrNotFound,
		},
		{
			userID:      a.userID,
			oldPassword: a.password,
			newPassword: "newPassword",
			valid:       true,
		},
	}

	for _, tc := range testCases {
		err := a.api.ChangePassword(ctx, tc.userID, tc.oldPassword, tc.newPassword)
		if tc.valid {
			a.NoError(err)
			a.password = tc.newPassword
			a.quicLogin(a.userID, a.password)
		} else if !tc.valid {
			a.Error(err)
			if tc.expectedError != nil {
				a.Equal(tc.expectedError, err)
			}
		}
	}
}

func (a *APITestSuite) TestF_RefreshToken() {
	ctx := context.TODO()

	testCases := []struct {
		userID        domain.UserID
		refreshToken  string
		valid         bool
		expectedError error
	}{
		{
			userID:        a.userID,
			refreshToken:  "invalid",
			valid:         false,
			expectedError: api.ErrAccessDenied,
		},
		{
			userID:        "invalid",
			refreshToken:  a.refreshToken,
			valid:         false,
			expectedError: api.ErrAccessDenied,
		},
		{
			userID:       a.userID,
			refreshToken: a.refreshToken,
			valid:        true,
		},
	}

	for _, tc := range testCases {
		accessToekn, err := a.api.RefreshToken(ctx, tc.userID, tc.refreshToken)
		if tc.valid {
			a.NoError(err)
			a.accessToken = accessToekn
		} else if !tc.valid {
			a.Error(err)
			if tc.expectedError != nil {
				a.Equal(tc.expectedError, err)
			}
		}
	}
}

func (a *APITestSuite) TestG_ResetPassword() {
	ctx := context.TODO()

	testCases := []struct {
		userID        domain.UserID
		valid         bool
		expectedError error
	}{
		{
			userID:        "invalid",
			valid:         false,
			expectedError: api.ErrNotFound,
		},
		{
			userID: a.userID,
			valid:  true,
		},
	}

	for _, tc := range testCases {
		resetPasswordToken, _, err := a.api.ResetPassword(ctx, tc.userID)
		if tc.valid {
			a.NoError(err)
			a.resetPasswordToken = resetPasswordToken
		} else if !tc.valid {
			a.Error(err)
			if tc.expectedError != nil {
				a.Equal(tc.expectedError, err)
			}
		}
	}
}

func (a *APITestSuite) TestH_SubmitResetPassword() {
	ctx := context.TODO()

	testCases := []struct {
		resetPasswordToken string
		newPassword        string
		valid              bool
		expectedError      error
	}{
		{
			resetPasswordToken: "invalid",
			newPassword:        "valid",
			valid:              false,
			expectedError:      api.ErrAccessDenied,
		},
		{
			resetPasswordToken: a.resetPasswordToken,
			newPassword:        "valid",
			valid:              true,
		},
	}

	for _, tc := range testCases {
		err := a.api.SubmitResetPassword(ctx, tc.resetPasswordToken, tc.newPassword)
		if tc.valid {
			a.NoError(err)
			a.password = tc.newPassword
			a.quicLogin(a.userID, a.password)
		} else if !tc.valid {
			a.Error(err)
			if tc.expectedError != nil {
				a.Equal(tc.expectedError, err)
			}
		}
	}
}

func (a *APITestSuite) TestI_DeleteAccount() {
	ctx := context.TODO()

	testCases := []struct {
		userID        domain.UserID
		password      string
		valid         bool
		expectedError error
	}{
		{
			userID:        "invalid",
			password:      a.password,
			valid:         false,
			expectedError: api.ErrNotFound,
		},
		{
			userID:        a.userID,
			password:      "invalid",
			valid:         false,
			expectedError: api.ErrInvalidPassword,
		},
		{
			userID:   a.userID,
			password: a.password,
			valid:    true,
		},
	}

	for _, tc := range testCases {
		err := a.api.DeleteAccount(ctx, tc.userID, tc.password)
		if tc.valid {
			a.NoError(err)
		} else if !tc.valid {
			a.Error(err)
			if tc.expectedError != nil {
				a.Equal(tc.expectedError, err)
			}
		}
	}
}

func (a *APITestSuite) quicLogin(userID domain.UserID, password string) {
	ctx := context.TODO()

	_, _, err := a.api.Login(ctx, userID, password)
	a.NoError(err)
}
