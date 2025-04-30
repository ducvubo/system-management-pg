-- name: CreateEquipmentMaintenance :execresult
INSERT INTO equipment_maintenance (
    eqp_mtn_id,eqp_mtn_res_id, eqp_mtn_name,  eqp_mtn_location, eqp_mtn_issue_description,
    eqp_mtn_reported_by, eqp_mtn_performed_by, eqp_mtn_date_reported, eqp_mtn_date_fixed, eqp_mtn_cost, eqp_mtn_note,
    createdBy, createdAt, updatedAt
) VALUES (
    ?,?,  ?, ?, ?, ?, ?, ?,  ?, ?, ?, ?, NOW(), NOW()
);

-- name: GetEquipmentMaintenance :one
SELECT eqp_mtn_id, eqp_mtn_name,  eqp_mtn_location, eqp_mtn_issue_description,
       eqp_mtn_reported_by, eqp_mtn_performed_by, eqp_mtn_date_reported, eqp_mtn_date_fixed, eqp_mtn_cost, eqp_mtn_note, isDeleted
FROM equipment_maintenance
WHERE eqp_mtn_id = ? AND eqp_mtn_res_id = ?;

-- name: UpdateEquipmentMaintenance :exec
UPDATE equipment_maintenance
SET eqp_mtn_name = ?, eqp_mtn_location = ?, eqp_mtn_issue_description = ?,
    eqp_mtn_reported_by = ?, eqp_mtn_performed_by = ?, eqp_mtn_date_reported = ?, eqp_mtn_date_fixed = ?, eqp_mtn_cost = ?, 
    eqp_mtn_note = ?, updatedAt = NOW(), updatedBy = ?
WHERE eqp_mtn_id = ? AND eqp_mtn_res_id = ?;

-- name: DeleteEquipmentMaintenance :exec
UPDATE equipment_maintenance
SET isDeleted = 1, deletedAt = NOW(), deletedBy = ?
WHERE eqp_mtn_id = ? AND eqp_mtn_res_id = ?;

-- name: RestoreEquipmentMaintenance :exec
UPDATE equipment_maintenance
SET isDeleted = 0, deletedAt = NULL, deletedBy = NULL
WHERE eqp_mtn_id = ? AND eqp_mtn_res_id = ?;

-- name: UpdateEquipmentMaintenanceStatus :exec
UPDATE equipment_maintenance
SET eqp_mtn_status = ?, updatedAt = NOW(), updatedBy = ?
WHERE eqp_mtn_id = ? AND eqp_mtn_res_id = ?;

-- name: GetListEquipmentMaintenance :many
WITH total_count AS (
    SELECT COUNT(*) AS total FROM equipment_maintenance WHERE equipment_maintenance.isDeleted = ? AND equipment_maintenance.eqp_mtn_name LIKE ? AND equipment_maintenance.eqp_mtn_res_id = ?
)
SELECT 
    eqp_mtn_id, eqp_mtn_name, eqp_mtn_location, eqp_mtn_issue_description,
    eqp_mtn_reported_by, eqp_mtn_performed_by, eqp_mtn_date_reported, eqp_mtn_date_fixed, eqp_mtn_cost, eqp_mtn_note, eqp_mtn_status,
    (SELECT total FROM total_count) AS total_items,
    COALESCE(CEIL((SELECT total FROM total_count) / NULLIF(CAST(? AS FLOAT), 0)), 0) AS total_pages
FROM equipment_maintenance
WHERE equipment_maintenance.isDeleted = ? AND equipment_maintenance.eqp_mtn_name LIKE ? AND equipment_maintenance.eqp_mtn_res_id = ?
LIMIT ? OFFSET ?;