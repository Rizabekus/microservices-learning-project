package grpc

import (
	"context"

	"github.com/Rizabekus/microservices-learning-project/order/internal/infrastructure/logger"
	"github.com/Rizabekus/microservices-learning-project/order/internal/service"
	pb "github.com/Rizabekus/microservices-learning-project/order/internal/transport/grpc/pb"
	"github.com/Rizabekus/microservices-learning-project/order/internal/transport/http/middlewares"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedOrderServiceServer

	Service service.Service
}

func New(service service.Service) *Server {
	return &Server{
		Service: service,
	}
}
func (s *Server) GetOrderInternal(ctx context.Context, req *pb.GetOrderRequest) (*pb.OrderResponse, error) {
	userID, ok := middlewares.UserIDFromContext(ctx)
	if !ok {
		logger.Log.Error("failed to extract user ID from context")
		return nil, status.Error(codes.Unauthenticated, "missing user id")
	}
	orderID, err := uuid.Parse(req.GetOrderId())
	if err != nil {
		logger.Log.Error("failed to parse order ID", "error", err)
		return nil, status.Error(codes.InvalidArgument, "invalid order id")
	}
	order, err := s.Service.GetOrder(ctx, userID, orderID)
	if err != nil {
		logger.Log.Error("failed to get order", "error", err)
		return nil, MapGRPCError(err)
	}
	return &pb.OrderResponse{
		Id:       order.ID.String(),
		UserId:   order.UserID.String(),
		Amount:   order.Amount,
		Currency: order.Currency,
		Status:   string(order.Status),
	}, nil
}
