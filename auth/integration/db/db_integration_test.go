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
	adapter "github.com/Bookil/microservices/auth/internal/adapters/db/mysql_adapter"
	"github.com/Bookil/microservices/auth/internal/application/core/domain"
	"github.com/docker/go-connections/nat"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type AuthDatabaseTestSuite struct {
	suite.Suite
	adapter *adapter.Adapter
	auth    *domain.Auth
}

func TestOAuthDatabaseTestSuite(t *testing.T) {
	suite.Run(t, new(AuthDatabaseTestSuite))
}

func (o *AuthDatabaseTestSuite) SetupSuite() {
	ctx := context.TODO()

	err := os.Setenv("AUTH_ENV", "test")
	if err != nil {
		log.Fatalf("Could not set the environment variable to test: %s", err)
	}

	mysqlConfig := &config.Read().Mysql

	port := fmt.Sprintf("%d/tcp", mysqlConfig.Port)

	req := testcontainers.ContainerRequest{
		Image:        "mysql:latest",
		ExposedPorts: []string{port},
		Env: map[string]string{
			"MYSQL_ROOT_PASSWORD": mysqlConfig.Password,
			"MYSQL_DATABASE":      mysqlConfig.DBName,
		},
		WaitingFor: wait.ForListeningPort(nat.Port(port)).WithStartupTimeout(5 * time.Minute),
	}

	mysqlContainer, connectErr := testcontainers.GenericContainer(ctx,
		testcontainers.GenericContainerRequest{
			ContainerRequest: req,
			Started:          true,
		})

	if connectErr != nil {
		log.Fatal("Failed to start Mysql:", connectErr)
	}

	endpoint, err := mysqlContainer.Endpoint(ctx, "")
	if err != nil {
		log.Fatal(err)
	}

	endPort, err := strconv.Atoi(strings.Split(endpoint, ":")[1])
	if err != nil {
		log.Fatal(err)
	}

	mysqlConfig.Port = endPort

	adapter, err := adapter.NewAdapter(mysqlConfig)
	if err != nil {
		log.Fatal(err)
	}

	o.adapter = adapter
}

func (o *AuthDatabaseTestSuite) TestA_ShouldCreateAuth() {
	ctx := context.TODO()

	testCases := []struct {
		auth  *domain.Auth
		Valid bool
	}{
		{
			auth:  domain.NewAuth("123456", "$^&fullyHashedPassword12$"),
			Valid: true,
		},
	}

	for _, tc := range testCases {
		auth, err := o.adapter.Create(ctx, tc.auth)
		if tc.Valid {
			o.NoError(err)
			o.auth = auth
		} else if !tc.Valid {
			o.Error(err)
			o.Nil(auth)
		}
	}
}

func (o *AuthDatabaseTestSuite) TestB_ShouldGetAuth() {
	ctx := context.TODO()

	testCases := []struct {
		UserID string
		Valid  bool
	}{
		{
			UserID: "invalid",
			Valid:  false,
		},
		{
			UserID: o.auth.UserID,
			Valid:  true,
		},
	}

	for _, tc := range testCases {
		auth, err := o.adapter.GetByID(ctx, tc.UserID)
		if tc.Valid {
			o.NoError(err)
			o.Equal(tc.UserID, auth.UserID)
			o.Equal(o.auth.HashedPassword, auth.HashedPassword)
		} else if !tc.Valid {
			o.Error(err)
			o.Nil(auth)
		}
	}
}

func (o *AuthDatabaseTestSuite) TestC_ShouldVerifyEmail() {
	ctx := context.TODO()

	testCases := []struct {
		UserID string
		Valid  bool
	}{
		{
			UserID: "invalid",
			Valid:  false,
		},
		{
			UserID: o.auth.UserID,
			Valid:  true,
		},
	}

	for _, tc := range testCases {
		auth, err := o.adapter.VerifyEmail(ctx, tc.UserID)
		if tc.Valid {
			o.NoError(err)
			o.Equal(tc.UserID, auth.UserID)
			o.Equal(true, auth.IsEmailVerified)
			o.NotEqual(o.auth.IsEmailVerified, auth.IsEmailVerified)
			o.auth.IsEmailVerified = auth.IsEmailVerified
		} else if !tc.Valid {
			o.Error(err)
			o.Nil(auth)
		}
	}
}

func (o *AuthDatabaseTestSuite) TestD_ShouldChangePassword() {
	ctx := context.TODO()
	testCases := []struct {
		UserID   string
		Valid    bool
		Password string
	}{
		{
			UserID:   "invalid",
			Valid:    false,
			Password: "passwordChange",
		},
		{
			UserID:   o.auth.UserID,
			Valid:    true,
			Password: "passwordChange",
		},
	}

	for _, tc := range testCases {
		auth, err := o.adapter.ChangePassword(ctx, tc.UserID, tc.Password)
		if tc.Valid {
			o.NoError(err)
			o.Equal(tc.UserID, auth.UserID)
			o.NotEqual(o.auth.HashedPassword, auth.HashedPassword)
			o.auth.HashedPassword = auth.HashedPassword
		} else if !tc.Valid {
			o.Error(err)
			o.Nil(auth)
		}
	}
}

func (o *AuthDatabaseTestSuite) TestE_ShouldIncrementFailedLoginAttempts() {
	ctx := context.TODO()

	testCases := []struct {
		UserID string
		Valid  bool
	}{
		{
			UserID: "invalid",
			Valid:  false,
		},
		{
			UserID: o.auth.UserID,
			Valid:  true,
		},
	}

	for _, tc := range testCases {
		auth, err := o.adapter.IncrementFailedLoginAttempts(ctx, tc.UserID)
		if tc.Valid {
			o.NoError(err)
			o.Equal(tc.UserID, auth.UserID)
			o.NotEqual(o.auth.FailedLoginAttempts, auth.FailedLoginAttempts)

			o.auth.FailedLoginAttempts = auth.FailedLoginAttempts
		} else if !tc.Valid {
			o.Error(err)
			o.Nil(auth)
		}
	}
}

func (o *AuthDatabaseTestSuite) TestF_ShouldClearFailedLoginAttempts() {
	ctx := context.TODO()

	testCases := []struct {
		UserID string
		Valid  bool
	}{
		{
			UserID: "invalid",
			Valid:  false,
		},
		{
			UserID: o.auth.UserID,
			Valid:  true,
		},
	}

	for _, tc := range testCases {
		auth, err := o.adapter.ClearFailedLoginAttempts(ctx, tc.UserID)
		if tc.Valid {
			o.NoError(err)
			o.Equal(tc.UserID, auth.UserID)
			o.NotEqual(o.auth.FailedLoginAttempts, auth.FailedLoginAttempts)

			o.auth.FailedLoginAttempts = auth.FailedLoginAttempts
		} else if !tc.Valid {
			o.Error(err)
			o.Nil(auth)
		}
	}
}

func (o *AuthDatabaseTestSuite) TestG_ShouldLockAccount() {
	ctx := context.TODO()

	testCases := []struct {
		UserID string
		Valid  bool
	}{
		{
			UserID: "invalid",
			Valid:  false,
		},
		{
			UserID: o.auth.UserID,
			Valid:  true,
		},
	}

	for _, tc := range testCases {
		auth, err := o.adapter.LockAccount(ctx, tc.UserID, 5*time.Minute)
		if tc.Valid {
			o.NoError(err)
			o.Equal(tc.UserID, auth.UserID)
			o.NotEqual(o.auth.AccountLockedUntil, auth.AccountLockedUntil)

			o.auth.AccountLockedUntil = auth.AccountLockedUntil
		} else if !tc.Valid {
			o.Error(err)
			o.Nil(auth)
		}
	}
}

func (o *AuthDatabaseTestSuite) TestH_ShouldUnLockAccount() {
	ctx := context.TODO()

	testCases := []struct {
		UserID string
		Valid  bool
	}{
		{
			UserID: "invalid",
			Valid:  false,
		},
		{
			UserID: o.auth.UserID,
			Valid:  true,
		},
	}

	for _, tc := range testCases {
		auth, err := o.adapter.UnlockAccount(ctx, tc.UserID)
		if tc.Valid {
			o.NoError(err)
			o.Equal(tc.UserID, auth.UserID)
			o.NotEqual(o.auth.AccountLockedUntil, auth.AccountLockedUntil)

			o.auth.AccountLockedUntil = auth.AccountLockedUntil
		} else if !tc.Valid {
			o.Error(err)
			o.Nil(auth)
		}
	}
}

func (o *AuthDatabaseTestSuite) TestI_ShouldDeleteAuth() {
	ctx := context.TODO()

	testCases := []struct {
		UserID string
		Valid  bool
	}{
		{
			UserID: "invalid",
			Valid:  false,
		},
		{
			UserID: o.auth.UserID,
			Valid:  true,
		},
	}

	for _, tc := range testCases {
		err := o.adapter.DeleteByID(ctx, tc.UserID)
		if tc.Valid {
			o.NoError(err)
		} else if !tc.Valid {
			o.Error(err)
		}
	}
}
