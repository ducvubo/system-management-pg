-- +goose Up
-- +goose StatementBegin
ALTER TABLE `user_management_account` ADD CONSTRAINT `user_management_account_usa_id_foreign` FOREIGN KEY(`usa_id`) REFERENCES `user_management_profile`(`us_id`);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE `user_management_account` DROP FOREIGN KEY `user_management_account_usa_id_foreign`;
-- +goose StatementEnd
