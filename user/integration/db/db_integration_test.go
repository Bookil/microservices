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

	"github.com/Bookil/microservices/user/config"
	"github.com/Bookil/microservices/user/internal/adapters/db"
	adapter "github.com/Bookil/microservices/user/internal/adapters/db/mysql_adapter"
	"github.com/Bookil/microservices/user/internal/application/core/domain"
	"github.com/docker/go-connections/nat"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type UserDatabaseTestSuite struct {
	suite.Suite
	adapter *adapter.Adapter
	user    *domain.User
}

func TestUserDatabaseTestSuite(t *testing.T) {
	suite.Run(t, new(UserDatabaseTestSuite))
}

func (o *UserDatabaseTestSuite) SetupSuite() {
	ctx := context.TODO()

	err := os.Setenv("USER_ENV", "test")
	if err != nil {
		log.Fatalf("Could not set the environment variable to test: %s", err)
	}

	mysqlConfig := config.Read().Mysql

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

	db, err := db.NewDB(&mysqlConfig)
	if err != nil {
		log.Fatalf("error getting DB instance:%v", err)
	}

	adapter, err := adapter.NewAdapter(db)
	if err != nil {
		log.Fatal(err)
	}

	o.adapter = adapter
}

func (o *UserDatabaseTestSuite) TestA_ShouldCreateUser() {
	ctx := context.TODO()

	testCases := []struct {
		user  *domain.User
		Valid bool
	}{
		{
			user:  domain.NewUser("john", "doe", "johndoe@email.com"),
			Valid: true,
		},
	}

	for _, tc := range testCases {
		user, err := o.adapter.Create(ctx, tc.user)
		if tc.Valid {
			o.NoError(err)
			o.user = user
		} else if !tc.Valid {
			o.Error(err)
			o.Nil(user)
		}
	}
}

func (o *UserDatabaseTestSuite) TestB_ShouldGetUserByID() {
	ctx := context.TODO()

	testCases := []struct {
		UserID domain.UserID
		Valid  bool
	}{
		{
			UserID: "invalid",
			Valid:  false,
		},
		{
			UserID: o.user.UserID,
			Valid:  true,
		},
	}

	for _, tc := range testCases {
		user, err := o.adapter.GetUserByID(ctx, tc.UserID)
		if tc.Valid {
			o.NoError(err)
			o.Equal(tc.UserID, user.UserID)
			o.Equal(o.user.FirstName, user.FirstName)
			o.Equal(o.user.LastName, user.LastName)
			o.Equal(o.user.Email, user.Email)
		} else if !tc.Valid {
			o.Error(err)
			o.Nil(user)
		}
	}
}

func (o *UserDatabaseTestSuite) TestC_ShouldGetUserByEmail() {
	ctx := context.TODO()

	testCases := []struct {
		Email string
		Valid bool
	}{
		{
			Email: "invalid",
			Valid: false,
		},
		{
			Email: o.user.Email,
			Valid: true,
		},
	}

	for _, tc := range testCases {
		user, err := o.adapter.GetUserByEmail(ctx, tc.Email)
		if tc.Valid {
			o.NoError(err)
			o.Equal(o.user.UserID, user.UserID)
			o.Equal(o.user.FirstName, user.FirstName)
			o.Equal(o.user.LastName, user.LastName)
			o.Equal(o.user.Email, user.Email)
		} else if !tc.Valid {
			o.Error(err)
			o.Nil(user)
		}
	}
}

func (o *UserDatabaseTestSuite) TestD_ShouldUpdateUpdate() {
	ctx := context.TODO()
	testCases := []struct {
		UserID    domain.UserID
		FirstName string
		LastName  string
		Password  string
		Valid     bool
	}{
		{
			UserID:    "invalid",
			FirstName: "valid",
			LastName:  "valid",
			Valid:     false,
		},
		{
			UserID:    o.user.UserID,
			FirstName: "valid",
			LastName:  "valid",
			Valid:     true,
		},
	}

	for _, tc := range testCases {
		user, err := o.adapter.Update(ctx, tc.UserID, tc.FirstName, tc.LastName)
		if tc.Valid {
			o.NoError(err)
			o.Equal(tc.UserID, user.UserID)
			o.user = user
		} else if !tc.Valid {
			o.Error(err)
			o.Nil(user)
		}
	}
}

func (o *UserDatabaseTestSuite) TestE_ShouldDeleteUser() {
	ctx := context.TODO()
	testCases := []struct {
		UserID domain.UserID
		Valid  bool
	}{
		{
			UserID: "invalid",
			Valid:  false,
		},
		{
			UserID: o.user.UserID,
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
