-- +goose Up
-- +goose StatementBegin
ALTER TABLE `messages`
ADD `is_hashed` tinyint(1) unsigned NOT NULL DEFAULT false;
-- +goose StatementEnd
---
-- +goose Down
-- +goose StatementBegin
ALTER TABLE `messages` DROP `is_hashed`;
-- +goose StatementEnd