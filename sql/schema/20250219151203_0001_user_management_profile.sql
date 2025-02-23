-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `user_management_profile`(
    `us_id` CHAR(36) NOT NULL,
    `us_name` VARCHAR(255) NULL,
    `us_avatar` TEXT NULL,
    `us_phone` VARCHAR(255) NULL,
    `us_gender` VARCHAR(255) NULL,
    `us_address` VARCHAR(255) NULL,
    `us_birthday` TIMESTAMP NULL,
    `createdAt` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updatedAt` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `deletedAt` TIMESTAMP NULL,
    `createdBy` VARCHAR(255) NULL,
    `updatedBy` VARCHAR(255) NULL,
    `deletedBy` VARCHAR(255) NULL,
    `isDeleted` INT NULL DEFAULT 0 COMMENT '0: chưa xóa, 1: đã xóa',
    PRIMARY KEY(`us_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT='user_management_account';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `user_management_profile`;
-- +goose StatementEnd
