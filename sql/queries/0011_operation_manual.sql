-- name: CreateOperationManual :execresult
INSERT INTO operation_manual (
    opera_manual_id,opera_manua_res_id, opera_manual_title, opera_manual_content, opera_manual_type, opera_manual_status, note,
    createdBy, createdAt, updatedAt
) VALUES (
    ?, ?,?, ?, ?, ?, ?, ?, NOW(), NOW()
);

-- name: GetOperationManual :one
SELECT opera_manual_id, opera_manual_title, opera_manual_content, opera_manual_type,
       opera_manual_status, note, isDeleted
FROM operation_manual
WHERE opera_manual_id = ? AND opera_manua_res_id = ?;

-- name: UpdateOperationManual :exec
UPDATE operation_manual
SET opera_manual_title = ?, opera_manual_content = ?, opera_manual_type = ?, note = ?,
    updatedAt = NOW(), updatedBy = ?
WHERE opera_manual_id = ? AND opera_manua_res_id = ?;

-- name: DeleteOperationManual :exec
UPDATE operation_manual
SET isDeleted = 1, deletedAt = NOW(), deletedBy = ?
WHERE opera_manual_id = ? AND opera_manua_res_id = ?;

-- name: RestoreOperationManual :exec
UPDATE operation_manual
SET isDeleted = 0, deletedAt = NULL, deletedBy = NULL
WHERE opera_manual_id = ? AND opera_manua_res_id = ?;

-- name: UpdateOperationManualStatus :exec
UPDATE operation_manual
SET opera_manual_status = ?, updatedAt = NOW(), updatedBy = ?
WHERE opera_manual_id = ? AND opera_manua_res_id = ?;

-- name: GetListOperationManual :many
WITH total_count AS (
    SELECT COUNT(*) AS total FROM operation_manual
    WHERE operation_manual.isDeleted = ? AND operation_manual.opera_manual_title LIKE ? AND operation_manual.opera_manua_res_id = ?
)
SELECT 
    opera_manual_id, opera_manual_title, opera_manual_type, opera_manual_status, note,
    (SELECT total FROM total_count) AS total_items,
    COALESCE(CEIL((SELECT total FROM total_count) / NULLIF(CAST(? AS FLOAT), 0)), 0) AS total_pages
FROM operation_manual
WHERE operation_manual.isDeleted = ? AND operation_manual.opera_manual_title LIKE ? AND operation_manual.opera_manua_res_id = ?
LIMIT ? OFFSET ?;
