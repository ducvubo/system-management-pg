-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `internal_note`(
    `itn_note_id` CHAR(36) NOT NULL,
    `itn_note_res_id` CHAR(24) NOT NULL,
    `itn_note_title` VARCHAR(255) NULL,
    `itn_note_content` TEXT NULL,
    `itn_note_type` VARCHAR(255) NULL,
    `createdAt` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updatedAt` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `deletedAt` TIMESTAMP NULL,
    `createdBy` VARCHAR(255) NULL,
    `updatedBy` VARCHAR(255) NULL,
    `deletedBy` VARCHAR(255) NULL,
    `isDeleted` INT NULL DEFAULT 0 COMMENT '0: chưa xóa, 1: đã xóa',
    PRIMARY KEY(`itn_note_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT='internal_note';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `internal_note`;
-- +goose StatementEnd
