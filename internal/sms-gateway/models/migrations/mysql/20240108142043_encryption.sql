-- +goose Up
-- +goose StatementBegin
ALTER TABLE `messages`
ADD `is_encrypted` tinyint(1) unsigned NOT NULL DEFAULT false;
-- +goose StatementEnd
-- +goose StatementBegin
ALTER TABLE `message_recipients`
MODIFY COLUMN `phone_number` varchar(128) NOT NULL;
-- +goose StatementEnd
---
-- +goose Down
-- +goose StatementBegin
ALTER TABLE `message_recipients`
MODIFY COLUMN `phone_number` varchar(16) NOT NULL;
-- +goose StatementEnd
-- +goose StatementBegin
ALTER TABLE `messages` DROP `is_encrypted`;
-- +goose StatementEnd