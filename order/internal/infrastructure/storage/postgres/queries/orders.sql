-- name: CreateOrder :exec
INSERT INTO orders(
    id,
    user_id,
    amount,
    currency,
    status,
    created_at
) VALUES (
    $1, $2, $3, $4, $5, $6
);

-- name: GetOrderByID :one
SELECT * FROM orders
WHERE id=$1;

-- name: GetOrdersByUserID :many
SELECT * FROM orders
WHERE user_id=$1
ORDER BY created_at DESC;

-- name: UpdateOrderStatus :exec
UPDATE orders
SET status = $2
WHERE id = $1;
