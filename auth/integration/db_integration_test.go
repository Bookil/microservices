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
	"github.com/Bookil/microservices/auth/internal/application/core/domain"
	"github.com/docker/go-connections/nat"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type AuthDatabaseTestSuite struct {
	suite.Suite
	adapter *db.Adapter
	auth    *domain.Auth
}

func TestOrderDatabaseTestSuite(t *testing.T) {
	suite.Run(t, new(AuthDatabaseTestSuite))
}

func (o *AuthDatabaseTestSuite) SetupSuite() {
	ctx := context.Background()

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

	adapter, err := db.NewAdapter(mysqlConfig)
	if err != nil {
		log.Fatal(err)
	}

	o.adapter = adapter
}

func (o *AuthDatabaseTestSuite) TestShouldCreateAuth() {
	ctx := context.Background()

	auth := domain.NewAuth("123456", "$^&fullyHashedPassword12$")

	err := o.adapter.Create(ctx, auth)

	o.Nil(err)

	o.auth = auth
}

func (o *AuthDatabaseTestSuite) TestShouldGetAuth() {
	ctx := context.Background()

	auth, err := o.adapter.GetByID(ctx, o.auth.UserID)

	o.Nil(err)
	o.Equal(o.auth.UserID, auth.UserID)
	o.Equal(o.auth.HashedPassword, auth.HashedPassword)
}

func (o *AuthDatabaseTestSuite) TestShouldVerifyEmail() {
	ctx := context.Background()

	savedAuth, err := o.adapter.VerifyEmail(ctx, o.auth.UserID)

	o.Nil(err)
	o.Equal(savedAuth.UserID, o.auth.UserID)
	o.Equal(savedAuth.IsEmailVerified, true)
	o.NotEqual(savedAuth.IsEmailVerified, o.auth.IsEmailVerified)

	o.auth.IsEmailVerified = savedAuth.IsEmailVerified
}

func (o *AuthDatabaseTestSuite) TestShouldChangePassword() {
	ctx := context.Background()

	savedAuth, err := o.adapter.ChangePassword(ctx, o.auth.UserID, "newFullyHashedPassword@w!##@#")

	o.Nil(err)
	o.Equal(o.auth.UserID, savedAuth.UserID)
	o.NotEqual(o.auth.HashedPassword, savedAuth.HashedPassword)

	o.auth.HashedPassword = savedAuth.HashedPassword
}

func (o *AuthDatabaseTestSuite) TestShouldIncrementFailedLoginAttempts() {
	ctx := context.Background()

	savedAuth, err := o.adapter.IncrementFailedLoginAttempts(ctx, o.auth.UserID)

	o.Nil(err)
	o.Equal(o.auth.UserID, savedAuth.UserID)
	o.NotEqual(o.auth.FailedLoginAttempts, savedAuth.FailedLoginAttempts)

	o.auth.FailedLoginAttempts = savedAuth.FailedLoginAttempts
}

func (o *AuthDatabaseTestSuite) TestShouldClearFailedLoginAttempts() {
	ctx := context.Background()

	savedAuth, err := o.adapter.ClearFailedLoginAttempts(ctx, o.auth.UserID)

	o.Nil(err)
	o.Equal(o.auth.UserID, savedAuth.UserID)
	o.NotEqual(o.auth.FailedLoginAttempts, savedAuth.FailedLoginAttempts)

	o.auth.FailedLoginAttempts = savedAuth.FailedLoginAttempts
}

func (o *AuthDatabaseTestSuite) TestShouldLockAccount() {
	ctx := context.Background()

	savedAuth, err := o.adapter.LockAccount(ctx, o.auth.UserID, 5*time.Minute)

	o.Nil(err)
	o.Equal(o.auth.UserID, savedAuth.UserID)
	o.NotEqual(o.auth.AccountLockedUntil, savedAuth.AccountLockedUntil)

	o.auth.AccountLockedUntil = savedAuth.AccountLockedUntil
}

func (o *AuthDatabaseTestSuite) TestShouldUnLockAccount() {
	ctx := context.Background()

	savedAuth, err := o.adapter.UnlockAccount(ctx, o.auth.UserID)

	o.Nil(err)
	o.Equal(o.auth.UserID, savedAuth.UserID)
	o.NotEqual(o.auth.AccountLockedUntil, savedAuth.AccountLockedUntil)

	o.auth.AccountLockedUntil = savedAuth.AccountLockedUntil
}

func (o *AuthDatabaseTestSuite) TestShouldDeleteAuth() {
	ctx := context.Background()

	err := o.adapter.DeleteByID(ctx, o.auth.UserID)

	o.Nil(err)
}
