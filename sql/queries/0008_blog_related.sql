-- name: GetRelatedBlogByBlogId :many
SELECT 
    bl_id,
    bl_rlt_id
FROM blog_related
WHERE bl_id = ?;

-- name: CreateRelatedBlog :exec
INSERT INTO blog_related (
    bl_id,
    bl_rlt_id
) VALUES (
    ?,
    ?
);

-- name: DeleteRelatedBlog :exec
DELETE FROM blog_related
WHERE bl_id = ?;