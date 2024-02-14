-- +goose Up
-- +goose StatementBegin
ALTER TABLE `devices`
ADD `last_seen` datetime NOT NULL;
-- +goose StatementEnd
-- +goose StatementBegin
UPDATE `devices`
SET `last_seen` = `updated_at`;
-- +goose StatementEnd
---
-- +goose Down
-- +goose StatementBegin
ALTER TABLE `devices` DROP `last_seen`;
-- +goose StatementEnd