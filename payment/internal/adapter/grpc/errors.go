package grpc

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrFailedChargeCreditCard = status.Errorf(codes.Internal, "failed to change credit card")
)
