-- name: GetBlogNoteByBlogId :many
SELECT 
    bl_note_id,
    bl_id,
    bl_content
FROM blog_note
WHERE bl_id = ?;

-- name: CreateBlogNote :execresult
INSERT INTO blog_note (
    bl_note_id,
    bl_id,
    bl_content
) VALUES (
    ?,
    ?,
    ?
);

-- name: DeleteBlogNote :exec
DELETE FROM blog_note
WHERE bl_id = ?;
