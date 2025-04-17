-- name: CreateUserPatoSession :execresult
INSERT INTO user_pato_session (
    uss_pt_id, usa_pt_id, uss_pt_rf, uss_pt_key_at, uss_pt_key_rf, uss_pt_client_id, uss_pt_login_time, createdAt
) VALUES (
    ?, ?, ?, ?, ?, ?, NOW(), NOW()
);

-- name: FindUserPatoSessionBySessionIdAndRefreshToken :one
SELECT uss_pt_id, usa_pt_id, uss_pt_rf, uss_pt_key_at, uss_pt_key_rf, uss_pt_client_id, uss_pt_login_time, updatedAt
FROM user_pato_session
WHERE uss_pt_client_id = ? AND uss_pt_rf = ?;

-- name: DeleteUserPatoSessionByClientIdAndUsaId :exec
DELETE FROM user_pato_session WHERE uss_pt_client_id = ? AND usa_pt_id = ?;