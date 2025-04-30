-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `operation_manual` (
    `opera_manual_id` CHAR(36) NOT NULL,
    `opera_manua_res_id` CHAR(24) NOT NULL,
    `opera_manual_title` VARCHAR(255) NOT NULL,
    `opera_manual_content` TEXT NOT NULL,
    `opera_manual_type` VARCHAR(100) NOT NULL,
    `opera_manual_status` ENUM('active', 'archived', 'deleted') DEFAULT 'active' COMMENT 'Trạng thái tài liệu',
    `note` TEXT NULL,
    `createdAt` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updatedAt` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `deletedAt` TIMESTAMP NULL,
    `createdBy` VARCHAR(255) NULL,
    `updatedBy` VARCHAR(255) NULL,
    `deletedBy` VARCHAR(255) NULL,
    `isDeleted` INT NULL DEFAULT 0 COMMENT '0: chưa xóa, 1: đã xóa',
    PRIMARY KEY (`opera_manual_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='operation_manual';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `operation_manual`;
-- +goose StatementEnd
