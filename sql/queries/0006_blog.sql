-- name: GetBlogByID :one
SELECT 
    bl_id,
    cat_bl_id,
    bl_title,
    bl_description,
    bl_slug,
    bl_image,
    bl_content,
    bl_status,
    bl_type,
    bl_view,
    bl_published_time,
    bl_published_schedule
FROM blog
WHERE bl_id = ?;

-- name: CheckBlogSlug :many
SELECT 
    bl_id,
    bl_slug
FROM blog
WHERE bl_slug = ?;

-- name: CreateBlog :execresult
INSERT INTO blog (
    bl_id,
    cat_bl_id,
    bl_title,
    bl_description,
    bl_slug,
    bl_status,
    bl_image,
    bl_content,
    bl_type,
    bl_view,
    createdAt,
    updatedAt,
    createdBy,
    updatedBy
) VALUES (
    ?,
    ?,
    ?,
    ?,  
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    NOW(),
    NOW(),
    ?,
    ?
);

-- name: UpdateBlog :exec
UPDATE blog
SET cat_bl_id = ?,
    bl_title = ?,
    bl_description = ?,
    bl_slug = ?,
    bl_image = ?,
    bl_content = ?,
    bl_view = ?,
    updatedAt = NOW(),
    updatedBy = ?
WHERE bl_id = ?;

-- name: DeleteBlog :exec
UPDATE blog
SET isDeleted = 1,
    deletedAt = NOW(),
    deletedBy = ?
WHERE bl_id = ?;

-- name: RestoreBlog :exec
UPDATE blog
SET isDeleted = 0,
    deletedAt = NULL,
    deletedBy = NULL
WHERE bl_id = ?;

-- name: GetListBlog :many
WITH total_count AS (
    SELECT COUNT(*) AS total FROM blog WHERE blog.isDeleted = ? AND blog.bl_title LIKE ?
)
SELECT 
    bl_id,
    cat_bl_id,
    bl_title,
    bl_description,
    bl_slug,
    bl_image,
    bl_content,
    bl_status,
    bl_type,
    bl_view,
    bl_published_time,
    bl_published_schedule,
    (SELECT total FROM total_count) AS total_items,
    CEIL((SELECT total FROM total_count) / CAST(? AS FLOAT)) AS total_pages
FROM blog
WHERE blog.isDeleted = ? AND blog.bl_title LIKE ?
LIMIT ? OFFSET ?;

-- name: GetListBlogByCategory :many
WITH total_count AS (
    SELECT COUNT(*) AS total FROM blog WHERE blog.isDeleted = ? AND blog.cat_bl_id = ?
)
SELECT 
    bl_id,
    cat_bl_id,
    bl_title,
    bl_description,
    bl_slug,
    bl_image,
    bl_content,
    bl_status,
    bl_type,
    bl_view,
    bl_published_time,
    bl_published_schedule,
    (SELECT total FROM total_count) AS total_items,
    CEIL((SELECT total FROM total_count) / CAST(? AS FLOAT)) AS total_pages
FROM blog
WHERE blog.isDeleted = ? AND blog.cat_bl_id = ?
LIMIT ? OFFSET ?;

-- name: UpdateStatusBlog :exec
UPDATE blog
SET bl_status = ?,
    updatedAt = NOW(),
    updatedBy = ?
WHERE bl_id = ?;