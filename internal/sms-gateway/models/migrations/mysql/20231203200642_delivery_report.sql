-- +goose Up
-- +goose StatementBegin
ALTER TABLE `messages`
ADD `with_delivery_report` tinyint(1) unsigned DEFAULT false;
-- +goose StatementEnd
---
-- +goose Down
-- +goose StatementBegin
ALTER TABLE `messages` DROP `with_delivery_report`;
-- +goose StatementEnd