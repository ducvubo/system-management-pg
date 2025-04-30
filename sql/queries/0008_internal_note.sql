-- name: CreateInternalNote :execresult
INSERT INTO internal_note (
    itn_note_id,itn_note_res_id, itn_note_title, itn_note_content, itn_note_type, createdBy, createdAt, updatedAt
) VALUES (
    ?,?, ?, ?, ?, ?, NOW(), NOW()
);

-- name: GetInternalNote :one
SELECT itn_note_id, itn_note_res_id, itn_note_title, itn_note_content, itn_note_type, isDeleted
FROM internal_note
WHERE itn_note_id = ? AND itn_note_res_id = ?;

-- name: UpdateInternalNote :exec
UPDATE internal_note
SET itn_note_title = ?, itn_note_content = ?, itn_note_type = ?, updatedAt = NOW(), updatedBy = ?
WHERE itn_note_id = ? AND itn_note_res_id = ?;

-- name: DeleteInternalNote :exec
UPDATE internal_note
SET isDeleted = 1, deletedAt = NOW(), deletedBy = ?
WHERE itn_note_id = ? AND itn_note_res_id = ?;

-- name: RestoreInternalNote :exec
UPDATE internal_note
SET isDeleted = 0, deletedAt = NULL, deletedBy = NULL
WHERE itn_note_id = ? AND itn_note_res_id = ?;

-- name: GetListInternalNote :many
WITH total_count AS (
    SELECT COUNT(*) AS total FROM internal_note WHERE internal_note.isDeleted = ? AND internal_note.itn_note_title LIKE ? AND internal_note.itn_note_res_id = ?
)
SELECT 
    itn_note_id, itn_note_title, itn_note_content, itn_note_type,
    (SELECT total FROM total_count) AS total_items,
    COALESCE(CEIL((SELECT total FROM total_count) / NULLIF(CAST(? AS FLOAT), 0)), 0) AS total_pages
FROM internal_note
WHERE internal_note.isDeleted = ? AND internal_note.itn_note_title LIKE ? AND internal_note.itn_note_res_id = ?
LIMIT ? OFFSET ?;