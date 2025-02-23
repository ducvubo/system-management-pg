-- name: CreateUserProfile :execresult
INSERT INTO user_management_profile (
    us_id, us_name, us_avatar, us_phone, us_gender, us_address, us_birthday, createdBy, createdAt, updatedAt
) VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW()
);

-- name: GetUserProfile :one
SELECT us_id, us_name, us_avatar, us_phone, us_gender, us_address, us_birthday, isDeleted
FROM user_management_profile
WHERE us_id = ?;

-- name: UpdateUserProfile :exec
UPDATE user_management_profile
SET us_name = ?, us_avatar = ?, us_phone = ?, us_gender = ?, us_address = ?, us_birthday = ?, updatedAt = NOW(), updatedBy = ?
WHERE us_id = ?;

-- name: DeleteUserProfile :exec
UPDATE user_management_profile
SET isDeleted = 1, deletedAt = NOW(), deletedBy = ?
WHERE us_id = ?;

-- name: RestoreUserProfile :exec
UPDATE user_management_profile
SET isDeleted = 0, deletedAt = NULL, deletedBy = NULL
WHERE us_id = ?;

-- name: GetListUserProfile :many
WITH total_count AS (
    SELECT COUNT(*) AS total FROM user_management_profile WHERE user_management_profile.isDeleted = ? AND user_management_profile.us_name LIKE ?
)
SELECT 
    us_id, us_name, us_avatar, us_phone, us_gender, us_address, us_birthday,
    (SELECT total FROM total_count) AS total_items,
    CEIL((SELECT total FROM total_count) / CAST(? AS FLOAT)) AS total_pages
FROM user_management_profile
WHERE user_management_profile.isDeleted = ? AND user_management_profile.us_name LIKE ?
LIMIT ? OFFSET ?;


