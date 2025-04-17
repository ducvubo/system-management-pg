-- +goose Up
-- +goose StatementBegin
ALTER TABLE `user_pato_account` ADD CONSTRAINT `user_pato_account_usa_pt_id_foreign` FOREIGN KEY(`usa_pt_id`) REFERENCES `user_pato_profile`(`us_pt_id`);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE `user_pato_account` DROP FOREIGN KEY `user_pato_account_usa_pt_id_foreign`;
-- +goose StatementEnd
