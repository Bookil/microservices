package cart

import (
	"context"
	"fmt"
	"log"
	"time"

	"product/config"
	"product/internal/adapters/grpc/interceptor"
	"product/internal/application/core/domain"

	retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"

	cartv1 "github.com/Bookil/Bookil-Proto/gen/golang/cart/v1"
	"github.com/sony/gobreaker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
)

type Adapter struct {
	cart cartv1.CartServiceClient
}

func generateURL(url *config.CartService) string {
	return fmt.Sprintf("%s:%d", url.Host, url.Port)
}

func NewAdapter(url *config.CartService) (*Adapter, error) {
	var opts []grpc.DialOption

	if config.CurrentEnv == config.Production {
		cb := gobreaker.NewCircuitBreaker(
			gobreaker.Settings{
				Name:        "cart",
				MaxRequests: 5,
				Timeout:     5 * time.Second,
				ReadyToTrip: func(counts gobreaker.Counts) bool {
					failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
					return failureRatio > 0.5
				},
				OnStateChange: func(name string, from gobreaker.State, to gobreaker.State) {
					log.Println("Circcuit breaker state changed:", name, from, to)
				},
			},
		)

		opts = append(opts, grpc.WithUnaryInterceptor(interceptor.CircuitBreakerInterceptor(cb)))

		opts = append(opts, grpc.WithUnaryInterceptor(
			retry.UnaryClientInterceptor(
				retry.WithCodes(codes.ResourceExhausted, codes.Unavailable),
				retry.WithMax(3),
				retry.WithPerRetryTimeout(time.Second),
				retry.WithBackoff(retry.BackoffLinear(2*time.Second)),
			),
		))
	}

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.NewClient(generateURL(url), opts...)
	if err != nil {
		return nil, err
	}

	client := cartv1.NewCartServiceClient(conn)

	return &Adapter{cart: client}, nil
}

func (a *Adapter) AddBookToCart(ctx context.Context, bookID domain.BookID, userID string) error {
	_, err := a.cart.AddBookToCart(ctx, &cartv1.AddBookToCartRequest{
		BookId:     uint32(bookID),
		CustomerId: userID,
	})
	if err != nil {
		return err
	}

	return nil
}

func (a *Adapter) DeleteBookFromCartByID(ctx context.Context, bookID domain.BookID, cartID uint) error {
	_, err := a.cart.DeleteBookFromCart(ctx, &cartv1.DeleteBookFromCartRequest{
		BookId: uint32(bookID),
		CartId: uint32(cartID),
	})
	if err != nil {
		return err
	}

	return nil
}
