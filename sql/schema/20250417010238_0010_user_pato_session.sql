-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `user_pato_session`(
    `uss_pt_id` CHAR(36) NOT NULL PRIMARY KEY,
    `usa_pt_id` INT NOT NULL,
    `uss_pt_rf` TEXT NOT NULL,
    `uss_pt_key_at` TEXT NOT NULL,
    `uss_pt_key_rf` TEXT NOT NULL,
    `uss_pt_client_id` VARCHAR(255) NOT NULL,
    `uss_pt_login_time` TIMESTAMP NOT NULL,
    `uss_pt_logout_time` TIMESTAMP NULL,
    `createdAt` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updatedAt` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `deletedAt` TIMESTAMP NULL,
    `createdBy` VARCHAR(255) NULL,
    `updatedBy` VARCHAR(255) NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT='user_pato_session';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `user_pato_session`;
-- +goose StatementEnd
