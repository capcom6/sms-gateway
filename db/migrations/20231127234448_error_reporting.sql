-- +goose Up
-- +goose StatementBegin
ALTER TABLE `message_recipients`
ADD `error` varchar(256);
-- +goose StatementEnd
---
-- +goose Down
-- +goose StatementBegin
ALTER TABLE `message_recipients` DROP `error`;
-- +goose StatementEnd