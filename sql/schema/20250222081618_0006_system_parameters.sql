-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `system_parameters`(
    `sys_para_id` CHAR(36) NOT NULL PRIMARY KEY,
    `sys_para_description` VARCHAR(255)  NULL,
    `sys_para_value` LONGTEXT NOT NULL,
    `createdAt` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updatedAt` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `createdBy` VARCHAR(255) NULL,
    `updatedBy` VARCHAR(255) NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT='system_parameters';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `system_parameters`;
-- +goose StatementEnd
