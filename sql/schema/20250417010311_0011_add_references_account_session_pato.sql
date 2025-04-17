-- +goose Up
-- +goose StatementBegin
ALTER TABLE
    `user_pato_session` ADD CONSTRAINT `user_pato_session_usa_id_foreign` FOREIGN KEY(`usa_pt_id`) REFERENCES `user_pato_account`(`usa_pt_id`);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE `user_pato_session` DROP FOREIGN KEY `user_pato_session_usa_id_foreign`;
-- +goose StatementEnd
