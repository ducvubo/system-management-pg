-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `user_pato_profile`(
    `us_pt_id` CHAR(36) NOT NULL,
    `us_pt_name` VARCHAR(255) NULL,
    `us_pt_avatar` TEXT NULL,
    `us_pt_phone` VARCHAR(255) NULL,
    `us_pt_gender` VARCHAR(255) NULL,
    `us_pt_address` VARCHAR(255) NULL,
    `us_pt_birthday` TIMESTAMP NULL,
    `createdAt` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updatedAt` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `deletedAt` TIMESTAMP NULL,
    `createdBy` VARCHAR(255) NULL,
    `updatedBy` VARCHAR(255) NULL,
    `deletedBy` VARCHAR(255) NULL,
    `isDeleted` INT NULL DEFAULT 0 COMMENT '0: chưa xóa, 1: đã xóa',
    PRIMARY KEY(`us_pt_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT='user_pato_profile';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `user_pato_profile`;
-- +goose StatementEnd
