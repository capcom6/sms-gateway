-- +goose Up
-- +goose StatementBegin
ALTER TABLE `message_recipients` DROP PRIMARY KEY,
    ADD UNIQUE `unq_message_recipients_message_id_phone_number` (`message_id`, `phone_number`);
-- +goose StatementEnd
-- +goose StatementBegin
ALTER TABLE `message_recipients`
ADD `id` SERIAL NOT NULL PRIMARY KEY FIRST;
-- +goose StatementEnd
---
-- +goose Down
-- +goose StatementBegin
ALTER TABLE `message_recipients` DROP `id`;
-- +goose StatementEnd
-- +goose StatementBegin
ALTER TABLE `message_recipients` DROP INDEX `unq_message_recipients_message_id_phone_number`,
    ADD PRIMARY KEY (`message_id`, `phone_number`);
-- +goose StatementEnd