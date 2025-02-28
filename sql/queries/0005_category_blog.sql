-- name: GetCategoryBlog :one
SELECT 
    cat_bl_id,
    cat_bl_name,
    cat_bl_description,
    cat_bl_slug,
    cat_bl_order,
    cat_bl_status
FROM category_blog
WHERE cat_bl_id = ?;

-- name: CreateCategoryBlog :execresult
INSERT INTO category_blog (cat_bl_id, cat_bl_name, cat_bl_description, cat_bl_slug,
    cat_bl_order,
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
    NOW(),
    NOW(),
    ?,
    ?
);

-- name: UpdateCategoryBlog :exec
UPDATE category_blog
SET cat_bl_name = ?,
    cat_bl_description = ?,
    cat_bl_slug = ?,
    cat_bl_order = ?,
    updatedAt = NOW(),
    updatedBy = ?
WHERE cat_bl_id = ?;

-- name: DeleteCategoryBlog :exec
UPDATE category_blog
SET isDeleted = 1,
    deletedAt = NOW(),
    deletedBy = ?
WHERE cat_bl_id = ?;

-- name: RestoreCategoryBlog :exec
UPDATE category_blog
SET isDeleted = 0,
    deletedAt = NULL,
    deletedBy = NULL
WHERE cat_bl_id = ?;

-- name: GetListCategoryBlog :many
WITH total_count AS (
    SELECT COUNT(*) AS total FROM category_blog WHERE category_blog.isDeleted = ? AND category_blog.cat_bl_name LIKE ?
)
SELECT 
    cat_bl_id,
    cat_bl_name,
    cat_bl_description,
    cat_bl_slug,
    cat_bl_order,
    cat_bl_status,
    (SELECT total FROM total_count) AS total_items,
    CEIL((SELECT total FROM total_count) / CAST(? AS FLOAT)) AS total_pages
FROM category_blog
WHERE category_blog.isDeleted = ? AND category_blog.cat_bl_name LIKE ?
LIMIT ? OFFSET ?;

-- name: UpdateStatusCategoryBlog :exec
UPDATE category_blog
SET cat_bl_status = ?,
    updatedAt = NOW(),
    updatedBy = ?
WHERE cat_bl_id = ?;
