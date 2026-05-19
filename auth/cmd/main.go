package main

import (
	"context"
	"log"
	"net/http"
	"os"

	passwordHasher "github.com/Rizabekus/microservices-learning-project/auth/internal/infrastructure/crypto"
	"github.com/Rizabekus/microservices-learning-project/auth/internal/infrastructure/logger"
	"github.com/Rizabekus/microservices-learning-project/auth/internal/infrastructure/storage/postgres"
	"github.com/Rizabekus/microservices-learning-project/auth/internal/infrastructure/token"
	"github.com/Rizabekus/microservices-learning-project/auth/internal/service"
	httpHandler "github.com/Rizabekus/microservices-learning-project/auth/internal/transport/http"
	"github.com/Rizabekus/microservices-learning-project/auth/internal/transport/http/middlewares"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("auth/config/.env")
	if err != nil {
		logger.Log.Error("Error loading .env file", "error", err)
		log.Fatal(err)
	}
	logger.Log.Info("Config loaded successfully")
	dsn := os.Getenv("DATABASE_URL")
	port := os.Getenv("PORT")
	secret := os.Getenv("JWT_SECRET")
	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		logger.Log.Error("Error connecting to database", "error", err)
		log.Fatal(err)
	}
	defer pool.Close()
	repo := postgres.New(pool)
	passwordHasher := passwordHasher.NewBcryptHasher()
	tokenManager := token.NewJWTTokenManager(secret)
	service := service.New(repo, passwordHasher, tokenManager)
	handler := httpHandler.New(service)
	runServer(handler, port)
}

func runServer(handler *httpHandler.Handler, port string) {
	router := chi.NewRouter()
	router.Route("/auth", func(r chi.Router) {
		r.Use(middlewares.LoggingMiddleware)
		r.Post("/register", handler.Register)
		r.Post("/login", handler.Login)
		r.Post("/refresh", handler.Refresh)
	})
	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	logger.Log.Info("Starting server on port " + port)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		logger.Log.Error("Error starting server", "error", err)
		log.Fatal(err)
	}
}
