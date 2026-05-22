package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Rizabekus/microservices-learning-project/order/internal/infrastructure/logger"
	"github.com/Rizabekus/microservices-learning-project/order/internal/infrastructure/storage/postgres"
	"github.com/Rizabekus/microservices-learning-project/order/internal/service"
	grpcHandler "github.com/Rizabekus/microservices-learning-project/order/internal/transport/grpc"
	pb "github.com/Rizabekus/microservices-learning-project/order/internal/transport/grpc/pb"
	httpHandler "github.com/Rizabekus/microservices-learning-project/order/internal/transport/http"
	"github.com/Rizabekus/microservices-learning-project/order/internal/transport/http/middlewares"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	err := godotenv.Load("order/config/.env")
	if err != nil {
		logger.Log.Error("Error loading .env file", "error", err)
		log.Fatal(err)
	}
	logger.Log.Info("Config loaded successfully")
	dsn := os.Getenv("DATABASE_URL")
	httpPort := os.Getenv("HTTP_PORT")
	grpcPort := os.Getenv("GRPC_PORT")
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		logger.Log.Error("Error connecting to database", "error", err)
		log.Fatal(err)
	}
	defer pool.Close()
	repo := postgres.New(pool)
	// tokenManager := token.NewJWTTokenManager(secret)
	service := service.New(repo)
	httpHandler := httpHandler.New(service)

	httpServer := runHTTPServer(httpHandler, httpPort)
	grpcServer := runGRPCServer(grpcPort, service)

	<-ctx.Done()
	logger.Log.Info("shutdown signal received")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(shutdownCtx); err != nil {
		logger.Log.Error("http shutdown error", "error", err)
	}

	grpcServer.GracefulStop()

	logger.Log.Info("servers stopped")
}

func runHTTPServer(handler *httpHandler.Handler, port string) *http.Server {
	jwtSecret := os.Getenv("JWT_SECRET")

	router := chi.NewRouter()
	router.Route("/orders", func(r chi.Router) {
		r.Use(middlewares.LoggingMiddleware)
		r.Use(middlewares.AuthMiddleware(jwtSecret))
		r.Post("/", handler.CreateOrder)
		r.Get("/{id}", handler.GetOrder)
		r.Get("/my", handler.GetMyOrders)
		r.Patch("/{id}", handler.CancelOrder)
	})
	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	httpServer := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	go func() {
		logger.Log.Info("Starting HTTP server on port " + port)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Log.Error("Error starting HTTP server", "error", err)
			log.Fatal(err)
		}
	}()
	return httpServer
}

func runGRPCServer(port string, service service.Service) *grpc.Server {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()

	orderServer := grpcHandler.New(service)

	pb.RegisterOrderServiceServer(grpcServer, orderServer)
	go func() {
		logger.Log.Info("starting grpc server", "port", port)

		if err := grpcServer.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()
	return grpcServer
}
