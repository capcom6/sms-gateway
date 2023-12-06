-- +goose Up
-- +goose StatementBegin
ALTER TABLE `messages`
ADD `sim_number` tinyint(1) unsigned;
-- +goose StatementEnd
---
-- +goose Down
-- +goose StatementBegin
ALTER TABLE `messages` DROP `sim_number`;
-- +goose StatementEnd