-- +goose Up
-- +goose StatementBegin
ALTER TABLE `devices`
ADD `last_seen` datetime(3) NOT NULL DEFAULT current_timestamp(3);
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