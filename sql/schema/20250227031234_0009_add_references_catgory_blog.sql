-- +goose Up
-- +goose StatementBegin
ALTER TABLE `blog`
ADD CONSTRAINT `fk_blog_category`
FOREIGN KEY (`cat_bl_id`) REFERENCES `category_blog` (`cat_bl_id`) ON DELETE CASCADE ON UPDATE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE `blog` DROP FOREIGN KEY `fk_blog_category`;
-- +goose StatementEnd
