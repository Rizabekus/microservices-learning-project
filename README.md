go-commerce-microservices/
├── README.md
├── docker-compose.yml
├── .gitignore
├── go.mod
├── go.sum
│
├── protos/
│   ├── auth.proto
│   ├── order.proto
│   ├── payment.proto
│   └── notification.proto
│
├── auth/
│   ├── cmd/
│   │   └── main.go
│   ├── internal/
│   │   ├── domain/
│   │   ├── service/
│   │   ├── repository/
│   │   └── transport/
│   │       ├── http/
│   │       └── grpc/
│   ├── migrations/
│   └── Dockerfile
│
├── order/
│   ├── cmd/
│   │   └── main.go
│   ├── internal/
│   ├── migrations/
│   └── Dockerfile
│
├── payment/
│   ├── cmd/
│   │   └── main.go
│   ├── internal/
│   ├── migrations/
│   └── Dockerfile
│
└── notification/
    ├── cmd/
    │   └── main.go
    ├── internal/
    └── Dockerfile