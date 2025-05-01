-- name: CreateOperationalCosts :execresult
INSERT INTO operational_costs (
    opera_cost_id,opera_cost_res_id, opera_cost_type, opera_cost_amount, opera_cost_description,
    opera_cost_date, createdBy, createdAt, updatedAt
) VALUES (
    ?, ?, ?, ?, ?, ?, ?, NOW(), NOW()
);

-- name: GetOperationalCosts :one
SELECT opera_cost_id, opera_cost_type, opera_cost_amount, opera_cost_description,
       opera_cost_date, opera_cost_status, isDeleted
FROM operational_costs
WHERE opera_cost_id = ? AND opera_cost_res_id = ?;

-- name: UpdateOperationalCosts :exec
UPDATE operational_costs
SET opera_cost_type = ?, opera_cost_amount = ?, opera_cost_description = ?, opera_cost_date = ?,
    updatedAt = NOW(), updatedBy = ?
WHERE opera_cost_id = ? AND opera_cost_res_id = ?;

-- name: DeleteOperationalCosts :exec
UPDATE operational_costs
SET isDeleted = 1, deletedAt = NOW(), deletedBy = ?
WHERE opera_cost_id = ? AND opera_cost_res_id = ?;

-- name: RestoreOperationalCosts :exec
UPDATE operational_costs
SET isDeleted = 0, deletedAt = NULL, deletedBy = NULL
WHERE opera_cost_id = ? AND opera_cost_res_id = ?;

-- name: UpdateOperationalCostsStatus :exec
UPDATE operational_costs
SET opera_cost_status = ?, updatedAt = NOW(), updatedBy = ?
WHERE opera_cost_id = ? AND opera_cost_res_id = ?;

-- name: GetListOperationalCostss :many
WITH total_count AS (
    SELECT COUNT(*) AS total FROM operational_costs
    WHERE operational_costs.isDeleted = ? AND operational_costs.opera_cost_type LIKE ? AND operational_costs.opera_cost_res_id = ?
)
SELECT 
    opera_cost_id, opera_cost_type, opera_cost_amount, opera_cost_date, opera_cost_status,opera_cost_description,
    (SELECT total FROM total_count) AS total_items,
    COALESCE(CEIL((SELECT total FROM total_count) / NULLIF(CAST(? AS FLOAT), 0)), 0) AS total_pages
FROM operational_costs
WHERE operational_costs.isDeleted = ? AND operational_costs.opera_cost_type LIKE ? AND operational_costs.opera_cost_res_id = ?
LIMIT ? OFFSET ?;
