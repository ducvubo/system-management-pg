-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `category_blog`(
    `cat_bl_id` CHAR(36) NOT NULL,
    `cat_bl_name` VARCHAR(255) NOT NULL,
    `cat_bl_description` VARCHAR(255) NULL,
    `cat_bl_slug` VARCHAR(255) NOT NULL,
    `cat_bl_order` INT NULL DEFAULT 0,
    `cat_bl_status` INT NULL DEFAULT 0 COMMENT '0: inactive, 1: active',
    `createdAt` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updatedAt` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `deletedAt` TIMESTAMP NULL,
    `createdBy` VARCHAR(255) NULL,
    `updatedBy` VARCHAR(255) NULL,
    `deletedBy` VARCHAR(255) NULL,
    `isDeleted` INT NULL DEFAULT 0 COMMENT '0: chưa xóa, 1: đã xóa',
    PRIMARY KEY(`cat_bl_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT='category_blog';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `category_blog`;
-- +goose StatementEnd
