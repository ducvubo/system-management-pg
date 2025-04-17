-- name: FindUserPatoAccountByEmailAndType :one
SELECT usa_pt_id, usa_pt_email, usa_pt_password, usa_pt_salt, usa_pt_active_time, usa_pt_locked_time, usa_pt_recover_pass_time, usa_pt_verify_time, usa_pt_verify_code, usa_pt_recover_pass_code, usa_pt_active, usa_pt_locked
FROM user_pato_account
WHERE usa_pt_email = ? AND usa_pt_type = ?;

-- name: CreateUserPatoAccount :execresult
INSERT INTO user_pato_account (
    usa_pt_id, usa_pt_email, usa_pt_password, usa_pt_salt, usa_pt_type,
    usa_pt_active, usa_pt_locked,createdAt
) VALUES (
    ?, ?, ?, ?, ?, ?, ?,  NOW()
);

-- name: ActiveUserPatoAccount :exec
UPDATE user_pato_account
SET usa_pt_active_time = NOW(),usa_pt_active = 0, updatedAt = NOW(), updatedBy = ?
WHERE usa_pt_id = ?;