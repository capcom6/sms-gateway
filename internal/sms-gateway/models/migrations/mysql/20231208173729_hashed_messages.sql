-- +goose Up
-- +goose StatementBegin
ALTER TABLE `messages`
ADD `is_hashed` tinyint(1) unsigned NOT NULL DEFAULT false;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE INDEX `idx_messages_is_hashed` USING HASH ON `messages` (`is_hashed`);
-- +goose StatementEnd
---
-- +goose Down
-- +goose StatementBegin
ALTER TABLE `messages` DROP `is_hashed`;
-- +goose StatementEnd