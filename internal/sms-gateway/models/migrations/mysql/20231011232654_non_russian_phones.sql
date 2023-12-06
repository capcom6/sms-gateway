-- +goose Up
-- +goose StatementBegin
ALTER TABLE `message_recipients`
MODIFY COLUMN `phone_number` varchar(16) NOT NULL;
-- +goose StatementEnd
--
-- +goose Down
-- +goose StatementBegin
ALTER TABLE `message_recipients`
MODIFY COLUMN `phone_number` char(11) NOT NULL;
-- +goose StatementEnd