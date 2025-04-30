-- name: CreateOperationalCost :execresult
INSERT INTO operational_costs (
    opera_cost_id,opera_cost_res_id, opera_cost_type, opera_cost_amount, opera_cost_description,
    opera_cost_date, opera_cost_status, createdBy, createdAt, updatedAt
) VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW()
);

-- name: GetOperationalCost :one
SELECT opera_cost_id, opera_cost_type, opera_cost_amount, opera_cost_description,
       opera_cost_date, opera_cost_status, isDeleted
FROM operational_costs
WHERE opera_cost_id = ? AND opera_cost_res_id = ?;

-- name: UpdateOperationalCost :exec
UPDATE operational_costs
SET opera_cost_type = ?, opera_cost_amount = ?, opera_cost_description = ?, opera_cost_date = ?,
    updatedAt = NOW(), updatedBy = ?
WHERE opera_cost_id = ? AND opera_cost_res_id = ?;

-- name: DeleteOperationalCost :exec
UPDATE operational_costs
SET isDeleted = 1, deletedAt = NOW(), deletedBy = ?
WHERE opera_cost_id = ? AND opera_cost_res_id = ?;

-- name: RestoreOperationalCost :exec
UPDATE operational_costs
SET isDeleted = 0, deletedAt = NULL, deletedBy = NULL
WHERE opera_cost_id = ? AND opera_cost_res_id = ?;

-- name: UpdateOperationalCostStatus :exec
UPDATE operational_costs
SET opera_cost_status = ?, updatedAt = NOW(), updatedBy = ?
WHERE opera_cost_id = ? AND opera_cost_res_id = ?;

-- name: GetListOperationalCosts :many
WITH total_count AS (
    SELECT COUNT(*) AS total FROM operational_costs
    WHERE operational_costs.isDeleted = ? AND operational_costs.opera_cost_type LIKE ? AND operational_costs.opera_cost_res_id = ?
)
SELECT 
    opera_cost_id, opera_cost_type, opera_cost_amount, opera_cost_date, opera_cost_status,
    (SELECT total FROM total_count) AS total_items,
    COALESCE(CEIL((SELECT total FROM total_count) / NULLIF(CAST(? AS FLOAT), 0)), 0) AS total_pages
FROM operational_costs
WHERE operational_costs.isDeleted = ? AND operational_costs.opera_cost_type LIKE ? AND operational_costs.opera_cost_res_id = ?
LIMIT ? OFFSET ?;
