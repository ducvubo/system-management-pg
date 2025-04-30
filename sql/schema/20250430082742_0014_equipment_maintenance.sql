-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `equipment_maintenance` (
    `eqp_mtn_id` CHAR(36) NOT NULL,
    `eqp_mtn_res_id` CHAR(24) NOT NULL,
    `eqp_mtn_name` VARCHAR(255) NULL,
    `eqp_mtn_location` VARCHAR(255) NULL,
    `eqp_mtn_issue_description` TEXT NULL,
    `eqp_mtn_reported_by` VARCHAR(255) NULL,
    `eqp_mtn_performed_by` VARCHAR(255) NULL,
    `eqp_mtn_date_reported` DATE NOT NULL,
    `eqp_mtn_date_fixed` DATE NULL,
    `eqp_mtn_cost` DECIMAL(38,0) NULL,
    `eqp_mtn_note` TEXT NULL,
    `eqp_mtn_status` ENUM('pending', 'in_progress', 'done', 'rejected') DEFAULT 'pending' COMMENT 'pending, in_progress, done, rejected',
    `createdAt` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updatedAt` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deletedAt` TIMESTAMP NULL,
    `createdBy` VARCHAR(255) NULL,
    `updatedBy` VARCHAR(255) NULL,
    `deletedBy` VARCHAR(255) NULL,
    `isDeleted` INT DEFAULT 0 COMMENT '0: chưa xóa, 1: đã xóa',
    PRIMARY KEY (`eqp_mtn_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='equipment_maintenance';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `equipment_maintenance`;
-- +goose StatementEnd
