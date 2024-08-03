package e2e

import (
	"context"
	"log"
	"strings"
	"testing"

	orderv1 "github.com/Bookil/Bookil-Proto/gen/golang/order/v1"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	tc "github.com/testcontainers/testcontainers-go/modules/compose"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type CreateOrderTestSuite struct {
	suite.Suite
	compose *tc.LocalDockerCompose
}

func TestOrderDatabaseTestSuite(t *testing.T) {
	suite.Run(t, new(CreateOrderTestSuite))
}

func (c *CreateOrderTestSuite) SetupSuite() {
	composeFilePaths := []string{"resources/docker-compose.yml"}

	identifier := strings.ToLower(uuid.New().String())

	compose := tc.NewLocalDockerCompose(composeFilePaths, identifier)

	c.compose = compose

	err := compose.WithCommand([]string{"up", "-d"}).Invoke().Error
	if err != nil {
		log.Fatalf("Could not run compose stack: %v", err)
	}
}

func (c *CreateOrderTestSuite) Test_Should_Create_Order() {
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.NewClient("localhost:8080", opts...)
	if err != nil {
		log.Fatalf("Failed to connect order service. Err: %v", err)
	}

	orderClient := orderv1.NewOrderServiceClient(conn)

	_, errCreate := orderClient.Create(context.Background(), &orderv1.CreateRequest{
		CustomerId: "1",
		Items: []*orderv1.Item{
			{
				ItemId:      "1",
				Name:        "Black Ushanka",
				ProductCode: "13",
				UnitPrice:   2.5,
				Quantity:    1,
			},
		},
	})

	c.Nil(errCreate)
}
