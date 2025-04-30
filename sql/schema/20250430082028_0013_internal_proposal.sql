-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `internal_proposal`(
    `itn_proposal_id` CHAR(36) NOT NULL,
    `itn_proposal_res_id` CHAR(24) NOT NULL,
    `itn_proposal_title` VARCHAR(255) NULL,
    `itn_proposal_content` TEXT NULL,
    `itn_proposal_type` VARCHAR(255) NULL,
    `itn_proposal_status` ENUM('pending', 'approved', 'rejected') DEFAULT 'pending' COMMENT 'Trạng thái đề xuất',
    `createdAt` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updatedAt` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `deletedAt` TIMESTAMP NULL,
    `createdBy` VARCHAR(255) NULL,
    `updatedBy` VARCHAR(255) NULL,
    `deletedBy` VARCHAR(255) NULL,
    `isDeleted` INT NULL DEFAULT 0 COMMENT '0: chưa xóa, 1: đã xóa',
    PRIMARY KEY(`itn_proposal_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT='internal_proposal';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `internal_proposal`;
-- +goose StatementEnd
