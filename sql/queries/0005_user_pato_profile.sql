-- name: CreateUserPatoProfile :execresult
INSERT INTO user_pato_profile (
    us_pt_id, us_pt_name, us_pt_avatar, us_pt_phone, us_pt_gender, us_pt_address, us_pt_birthday, createdBy, createdAt, updatedAt
) VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW()
);

-- name: GetUserPatoProfile :one
SELECT us_pt_id, us_pt_name, us_pt_avatar, us_pt_phone, us_pt_gender, us_pt_address, us_pt_birthday, isDeleted
FROM user_pato_profile
WHERE us_pt_id = ?;

-- name: UpdateUserPatoProfile :exec
UPDATE user_pato_profile
SET us_pt_name = ?, us_pt_avatar = ?, us_pt_phone = ?, us_pt_gender = ?, us_pt_address = ?, us_pt_birthday = ?, updatedAt = NOW(), updatedBy = ?
WHERE us_pt_id = ?;

-- name: DeleteUserPatoProfile :exec
UPDATE user_pato_profile
SET isDeleted = 1, deletedAt = NOW(), deletedBy = ?
WHERE us_pt_id = ?;

-- name: RestoreUserPatoProfile :exec
UPDATE user_pato_profile
SET isDeleted = 0, deletedAt = NULL, deletedBy = NULL
WHERE us_pt_id = ?;

-- name: GetListUserPatoProfile :many
WITH total_count AS (
    SELECT COUNT(*) AS total FROM user_pato_profile WHERE user_pato_profile.isDeleted = ? AND user_pato_profile.us_pt_name LIKE ?
)
SELECT 
    us_pt_id, us_pt_name, us_pt_avatar, us_pt_phone, us_pt_gender, us_pt_address, us_pt_birthday,
    (SELECT total FROM total_count) AS total_items,
    CEIL((SELECT total FROM total_count) / CAST(? AS FLOAT)) AS total_pages
FROM user_pato_profile
WHERE user_pato_profile.isDeleted = ? AND user_pato_profile.us_pt_name LIKE ?
LIMIT ? OFFSET ?;


