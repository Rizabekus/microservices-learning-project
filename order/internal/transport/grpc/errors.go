package grpc

import (
	"errors"

	"github.com/Rizabekus/microservices-learning-project/order/internal/domain"
	"github.com/Rizabekus/microservices-learning-project/order/internal/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrUnauthenticated = status.Error(codes.Unauthenticated, "unauthenticated")
)

func MapGRPCError(err error) error {
	switch {
	// validation (domain)
	case errors.Is(err, domain.ErrInvalidUserID):
		return status.Error(codes.InvalidArgument, "invalid user id")

	case errors.Is(err, domain.ErrInvalidAmount):
		return status.Error(codes.InvalidArgument, "invalid amount")

	case errors.Is(err, domain.ErrInvalidCurrency):
		return status.Error(codes.InvalidArgument, "invalid currency")

	// business logic (service)

	case errors.Is(err, service.ErrOrderNotFound):
		return status.Error(codes.NotFound, "order not found")

	case errors.Is(err, service.ErrForbidden):
		return status.Error(codes.PermissionDenied, "order does not belong to user")

	case errors.Is(err, service.ErrOrderAlreadyCancelled):
		return status.Error(codes.FailedPrecondition, "order already cancelled")

	case errors.Is(err, service.ErrCannotCancelPaidOrder):
		return status.Error(codes.FailedPrecondition, "cannot cancel paid order")

	default:
		return status.Error(codes.Internal, "internal server error")
	}
}
