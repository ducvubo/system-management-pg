-- name: CreateUserSession :execresult
INSERT INTO user_management_session (
    uss_id, usa_id, uss_rf, uss_key_at, uss_key_rf, uss_client_id, uss_login_time, createdAt
) VALUES (
    ?, ?, ?, ?, ?, ?, NOW(), NOW()
);

-- name: FindUserSessionBySessionIdAndRefreshToken :one
SELECT uss_id, usa_id, uss_rf, uss_key_at, uss_key_rf, uss_client_id, uss_login_time, updatedAt
FROM user_management_session
WHERE uss_client_id = ? AND uss_rf = ?;

-- name: DeleteUserSessionByClientIdAndUsaId :exec
DELETE FROM user_management_session WHERE uss_client_id = ? AND usa_id = ?;