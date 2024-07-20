package grpc

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrFailedPlaceOrder = status.Errorf(codes.Internal, "failed to place order")
)
