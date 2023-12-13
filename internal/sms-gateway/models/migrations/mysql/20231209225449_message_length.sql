-- +goose Up
-- +goose StatementBegin
ALTER TABLE `messages`
MODIFY COLUMN `message` text NOT NULL;
-- +goose StatementEnd
---
-- +goose Down
-- +goose StatementBegin
ALTER TABLE `messages`
MODIFY COLUMN `message` tinytext NOT NULL;
-- +goose StatementEnd