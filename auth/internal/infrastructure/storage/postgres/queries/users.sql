-- name: CreateUser :exec
INSERT INTO users (
    id,
    first_name,
    last_name,
    phone,
    email,
    password_hash,
    created_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
);


-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = $1;


-- name: GetUserByID :one
SELECT *
FROM users
WHERE id = $1;