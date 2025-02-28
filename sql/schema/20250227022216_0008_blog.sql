-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `blog`(
    `bl_id` CHAR(36) NOT NULL,
    `cat_bl_id` CHAR(36) NOT NULL,
    `bl_title` VARCHAR(255) NOT NULL,
    `bl_description` VARCHAR(255) NULL,
    `bl_slug` VARCHAR(255) NOT NULL,
    `bl_image` LONGTEXT NULL,
    `bl_content` LONGTEXT NOT NULL,
    `bl_status` INT NULL DEFAULT 0 COMMENT '0: Bản nháp, 1: chờ duyệt, 2: Từ chối duyệt, 3 đợi xuất bản, 4: lên lịch xuất bản, 5: Đã xuất bản, 6: Không xuất bản',
    `bl_type` INT NULL DEFAULT 0 COMMENT '0: Bài viết, 1: Video, 3: Ảnh',
    `bl_view` INT NULL DEFAULT 0,
    `bl_published_time` TIMESTAMP NULL,
    `bl_published_schedule` TIMESTAMP NULL,
    `createdAt` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updatedAt` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `deletedAt` TIMESTAMP NULL,
    `createdBy` VARCHAR(255) NULL,
    `updatedBy` VARCHAR(255) NULL,
    `deletedBy` VARCHAR(255) NULL,
    `isDeleted` INT NULL DEFAULT 0 COMMENT '0: chưa xóa, 1: đã xóa',
    PRIMARY KEY(`bl_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT='blog';

-- ALTER TABLE `blog`
-- ADD CONSTRAINT `fk_blog_category`
-- FOREIGN KEY (`cat_bl_id`) REFERENCES `category_blog` (`cat_bl_id`) ON DELETE CASCADE ON UPDATE CASCADE;

-- CREATE TABLE IF NOT EXISTS `blog_related` (
--     `bl_id` CHAR(36) NOT NULL,
--     `bl_related_id` VARCHAR(255) NOT NULL,
--     PRIMARY KEY (`bl_id`, `bl_related_id`), -- Sửa: Khóa chính là tổ hợp 2 cột
--     FOREIGN KEY (`bl_id`) REFERENCES `blog` (`bl_id`) ON DELETE CASCADE ON UPDATE CASCADE
-- ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Related articles for blog';

-- CREATE TABLE IF NOT EXISTS `blog_note` (
--     `bl_id` CHAR(36) NOT NULL,
--     `bl_note` VARCHAR(4000) NOT NULL,
--     PRIMARY KEY (`bl_id`, `bl_note`(255)), -- Sửa: Giới hạn độ dài cho khóa chính trong MySQL
--     FOREIGN KEY (`bl_id`) REFERENCES `blog` (`bl_id`) ON DELETE CASCADE ON UPDATE CASCADE
-- ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Notes for blog articles';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `blog`;
-- ALTER TABLE `blog` DROP FOREIGN KEY `blog_ibfk_1`;
-- DROP TABLE IF EXISTS `blog_related`;
-- DROP TABLE IF EXISTS `blog_note`;
-- +goose StatementEnd
