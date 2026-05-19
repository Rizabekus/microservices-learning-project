-- name: CreateSession :exec
INSERT INTO sessions (
    id,
    user_id,
    token,
    expires_at,
    created_at,
    revoked_at
) VALUES (
    $1, $2, $3, $4, $5, $6
);


-- name: GetSessionByToken :one
SELECT *
FROM sessions
WHERE token = $1;


-- name: DeleteSessionByToken :exec
DELETE FROM sessions
WHERE token = $1;