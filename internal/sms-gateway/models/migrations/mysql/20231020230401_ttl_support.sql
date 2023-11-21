-- +goose Up
-- +goose StatementBegin
ALTER TABLE `messages`
ADD `valid_until` datetime(3) DEFAULT NULL;
-- +goose StatementEnd
---
-- +goose Down
-- +goose StatementBegin
ALTER TABLE `messages` DROP `valid_until`;
-- +goose StatementEnd