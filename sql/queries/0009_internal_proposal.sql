-- name: CreateInternalProposal :execresult
INSERT INTO internal_proposal (
    itn_proposal_id,itn_proposal_res_id, itn_proposal_title, itn_proposal_content, itn_proposal_type, createdBy, createdAt, updatedAt
) VALUES (
    ?,?, ?, ?, ?, ?, NOW(), NOW()
);

-- name: GetInternalProposal :one
SELECT itn_proposal_id, itn_proposal_res_id, itn_proposal_title, itn_proposal_content, itn_proposal_type, isDeleted
FROM internal_proposal
WHERE itn_proposal_id = ? AND itn_proposal_res_id = ?;

-- name: UpdateInternalProposal :exec
UPDATE internal_proposal
SET itn_proposal_title = ?, itn_proposal_content = ?, itn_proposal_type = ?, updatedAt = NOW(), updatedBy = ?
WHERE itn_proposal_id = ? AND itn_proposal_res_id = ?;

-- name: DeleteInternalProposal :exec
UPDATE internal_proposal
SET isDeleted = 1, deletedAt = NOW(), deletedBy = ?
WHERE itn_proposal_id = ? AND itn_proposal_res_id = ?;

-- name: RestoreInternalProposal :exec
UPDATE internal_proposal
SET isDeleted = 0, deletedAt = NULL, deletedBy = NULL
WHERE itn_proposal_id = ? AND itn_proposal_res_id = ?;

-- name: UpdateInternalProposalStatus :exec
UPDATE internal_proposal
SET itn_proposal_status = ?, updatedAt = NOW(), updatedBy = ?
WHERE itn_proposal_id = ? AND itn_proposal_res_id = ?;

-- name: GetListInternalProposal :many
WITH total_count AS (
    SELECT COUNT(*) AS total FROM internal_proposal WHERE internal_proposal.isDeleted = ? AND internal_proposal.itn_proposal_title LIKE ? AND internal_proposal.itn_proposal_res_id = ?
)
SELECT 
    itn_proposal_id, itn_proposal_title, itn_proposal_content, itn_proposal_type, itn_proposal_status,
    (SELECT total FROM total_count) AS total_items,
    COALESCE(CEIL((SELECT total FROM total_count) / NULLIF(CAST(? AS FLOAT), 0)), 0) AS total_pages
FROM internal_proposal
WHERE internal_proposal.isDeleted = ? AND internal_proposal.itn_proposal_title LIKE ? AND internal_proposal.itn_proposal_res_id = ?
LIMIT ? OFFSET ?;