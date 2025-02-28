-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `blog_note`(
    `bl_note_id` CHAR(36) NOT NULL PRIMARY KEY,
    `bl_id` CHAR(36) NOT NULL,
    `bl_content` LONGTEXT NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT='blog_note';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `blog_note`;
-- +goose StatementEnd
