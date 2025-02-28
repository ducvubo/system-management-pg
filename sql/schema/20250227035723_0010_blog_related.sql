-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `blog_related`(
    `bl_id` CHAR(36) NOT NULL,
    `bl_rlt_id` CHAR(36) NOT NULL,
    PRIMARY KEY(`bl_id`, `bl_rlt_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT='blog_related';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `blog_related`;
-- +goose StatementEnd
