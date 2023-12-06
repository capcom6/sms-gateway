-- +goose Up
-- +goose StatementBegin
CREATE UNIQUE INDEX IF NOT EXISTS `unq_messages_id_device` ON `messages`(`ext_id`, `device_id`);
-- +goose StatementEnd
-- +goose StatementBegin
DROP INDEX IF EXISTS `unq_messages_device_id` ON `messages`;
-- +goose StatementEnd
--
-- +goose Down
-- +goose StatementBegin
CREATE UNIQUE INDEX IF NOT EXISTS `unq_messages_device_id` ON `messages`(`device_id`, `ext_id`);
-- +goose StatementEnd
-- +goose StatementBegin
DROP INDEX IF EXISTS `unq_messages_id_device` ON `messages`;
-- +goose StatementEnd