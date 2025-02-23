-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `user_management_account`(
    `usa_id` CHAR(36) NOT NULL,
    `usa_email` VARCHAR(255) NOT NULL,
    `usa_salt` VARCHAR(255) NOT NULL,
    `usa_password` VARCHAR(255) NOT NULL,
    `usa_active_time` TIMESTAMP NULL,
    `usa_locked_time` TIMESTAMP NULL,
    `usa_recover_pass_time` TIMESTAMP NULL,
    `usa_verify_time` TIMESTAMP NULL,
    `usa_verify_code` VARCHAR(255) NULL,
    `usa_recover_pass_code` VARCHAR(255)  NULL,
    `usa_active` INT NULL COMMENT '0: kích hoạt',
    `usa_locked` INT NULL,
    `createdAt` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updatedAt` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `deletedAt` TIMESTAMP NULL,
    `createdBy` VARCHAR(255) NULL,
    `updatedBy` VARCHAR(255) NULL,
    `deletedBy` VARCHAR(255) NULL,
    `isDeleted` INT NULL DEFAULT 0 COMMENT '0: chưa xóa, 1: đã xóa',
    PRIMARY KEY(`usa_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT='user_management_account';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `user_management_account`;
-- +goose StatementEnd
