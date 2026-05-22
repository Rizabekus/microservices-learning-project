-- +goose Up
CREATE TABLE orders (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    amount BIGINT NOT NULL,
    currency VARCHAR(10) NOT NULL,
    status VARCHAR(50) NOT NULL CHECK (
        status IN ('pending', 'paid', 'cancelled')
    ),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_orders_user_id ON orders(user_id);

-- +goose Down
DROP INDEX IF EXISTS idx_orders_user_id;

DROP TABLE IF EXISTS orders;