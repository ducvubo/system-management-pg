-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `operational_costs` (
    `opera_cost_id` CHAR(36) NOT NULL,
    `opera_cost_res_id` CHAR(24) NOT NULL,
    `opera_cost_type` VARCHAR(100) NOT NULL,
    `opera_cost_amount` DECIMAL(10,2) NOT NULL,
    `opera_cost_description` TEXT NULL,
    `opera_cost_date` DATE NOT NULL,
    `opera_cost_status` ENUM('pending', 'paid', 'canceled') DEFAULT 'pending' COMMENT 'Trạng thái chi phí',
    `createdAt` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updatedAt` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `deletedAt` TIMESTAMP NULL,
    `createdBy` VARCHAR(255) NULL,
    `updatedBy` VARCHAR(255) NULL,
    `deletedBy` VARCHAR(255) NULL,
    `isDeleted` INT NULL DEFAULT 0 COMMENT '0: chưa xóa, 1: đã xóa',
    PRIMARY KEY (`opera_cost_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='operational_costs';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `operational_costs`;
-- +goose StatementEnd
