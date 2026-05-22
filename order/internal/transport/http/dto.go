package http

type CreateOrderDTO struct {
	Amount   int64  `json:"amount"`
	Currency string `json:"currency"`
}

type UpdateOrderStatusDTO struct {
	Status string `json:"status"`
}
