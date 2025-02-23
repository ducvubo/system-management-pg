-- name: SaveSystemParameter :execresult
INSERT INTO system_parameters (sys_para_id, sys_para_description, sys_para_value, createdAt, updatedAt, createdBy, updatedBy)
VALUES (?, ?, ?, NOW(), NOW(), ?, ?)
ON DUPLICATE KEY UPDATE sys_para_value = ?, updatedAt = NOW(), updatedBy = ?, sys_para_description = ?;

-- name: GetSystemParameter :one
SELECT sys_para_id, sys_para_description, sys_para_value, createdAt, updatedAt, createdBy, updatedBy
FROM system_parameters
WHERE sys_para_id = ?;

-- name: GetAllSystemParameters :many
SELECT sys_para_id, sys_para_description, sys_para_value, createdAt, updatedAt, createdBy, updatedBy
FROM system_parameters;