package middlewares

import (
	"context"
	"errors"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type contextKey string

const userIDKey contextKey = "user_id"

func AuthInterceptor(secret string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(codes.Unauthenticated, "missing metadata")
		}

		authHeaders := md.Get("authorization")
		if len(authHeaders) == 0 {
			return nil, status.Error(codes.Unauthenticated, "missing authorization header")
		}

		authHeader := authHeaders[0]

		const prefix = "Bearer "

		if !strings.HasPrefix(authHeader, prefix) {
			return nil, status.Error(codes.Unauthenticated, "invalid authorization header")
		}

		tokenString := strings.TrimPrefix(authHeader, prefix)

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {

			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}

			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			return nil, status.Error(codes.Unauthenticated, "invalid token")
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return nil, status.Error(codes.Unauthenticated, "invalid token claims")
		}

		tokenType, ok := claims["type"].(string)
		if !ok || tokenType != "access" {
			return nil, status.Error(codes.Unauthenticated, "invalid token type")
		}

		userIDStr, ok := claims["user_id"].(string)
		if !ok {
			return nil, status.Error(codes.Unauthenticated, "invalid user id")
		}

		userID, err := uuid.Parse(userIDStr)
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, "invalid user id")
		}

		ctx = context.WithValue(ctx, userIDKey, userID)

		return handler(ctx, req)
	}
}
