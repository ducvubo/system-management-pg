-- name: CreateUserAccount :execresult
INSERT INTO user_management_account (
    usa_id, usa_email, usa_password, usa_salt, usa_active_time, usa_active, usa_locked, createdBy, createdAt
) VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?, NOW()
);

-- name: FindUserAccountByEmail :one
SELECT usa_id, usa_email, usa_password, usa_salt, usa_active_time, usa_locked_time, usa_recover_pass_time, usa_verify_time, usa_verify_code, usa_recover_pass_code, usa_active, usa_locked
FROM user_management_account
WHERE usa_email = ?;

-- name: FindUserAccountById :one
SELECT usa_id, usa_email, usa_password, usa_salt, usa_active_time, usa_locked_time, usa_recover_pass_time, usa_verify_time, usa_verify_code, usa_recover_pass_code, usa_active, usa_locked
FROM user_management_account
WHERE usa_id = ?;


