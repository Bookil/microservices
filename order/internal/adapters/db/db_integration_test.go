package db_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/Bookil/microservices/order/config"
	"github.com/Bookil/microservices/order/internal/adapters/db"
	"github.com/Bookil/microservices/order/internal/application/core/domain"
	"github.com/docker/go-connections/nat"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type OrderDatabaseTestSuite struct {
	suite.Suite
	configs *config.Mysql
}

func TestOrderDatabaseTestSuite(t *testing.T) {
	suite.Run(t, new(OrderDatabaseTestSuite))
}

func (o *OrderDatabaseTestSuite) SetupSuite() {
	ctx := context.Background()

	err := os.Setenv("ORDER_ENV", "test")
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
	o.configs = mysqlConfig
	o.configs.Port = endPort
}

func (o *OrderDatabaseTestSuite) Test_Should_Save_Order() {
	adapter, err := db.NewAdapter(o.configs)

	o.Nil(err)

	orderItem := domain.NewOrderItem("item1", "123", 2.5, 1)
	order := domain.NewOrder("1", []*domain.OrderItem{orderItem})

	_, err = adapter.Save(order)

	o.Nil(err)
}

// func (o *OrderDatabaseTestSuite) Test_Should_Get_Order() {
// 	adapter, _ := db.NewAdapter(o.configs)

// 	orderItem := domain.NewOrderItem("item1", "123", 2.5, 1)
// 	order := domain.NewOrder("1", []*domain.OrderItem{orderItem})

// 	_,err:= adapter.Save(order)

// 	o.Nil(err)

// 	ord, _ := adapter.Get(order.ID)

// 	o.Equal("1", ord.CustomerID)
// }
