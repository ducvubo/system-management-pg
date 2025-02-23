-- +goose Up
-- +goose StatementBegin
ALTER TABLE
    `user_management_session` ADD CONSTRAINT `user_management_session_usa_id_foreign` FOREIGN KEY(`usa_id`) REFERENCES `user_management_account`(`usa_id`);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE `user_management_session` DROP FOREIGN KEY `user_management_session_usa_id_foreign`;
-- +goose StatementEnd
