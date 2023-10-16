-- +goose Up
-- +goose StatementBegin
ALTER TABLE `messages`
MODIFY COLUMN `state` enum(
        'Pending',
        'Processed',
        'Sent',
        'Delivered',
        'Failed'
    ) NOT NULL DEFAULT 'Pending';
-- +goose StatementEnd
-- +goose StatementBegin
ALTER TABLE `message_recipients`
MODIFY COLUMN `state` enum(
        'Pending',
        'Processed',
        'Sent',
        'Delivered',
        'Failed'
    ) NOT NULL DEFAULT 'Pending';
-- +goose StatementEnd
--
-- +goose Down
-- +goose StatementBegin
ALTER TABLE `message_recipients`
MODIFY COLUMN `state` ENUM('Pending', 'Sent', 'Delivered', 'Failed') NOT NULL DEFAULT 'Pending';
-- +goose StatementEnd
-- +goose StatementBegin
ALTER TABLE `messages`
MODIFY COLUMN `state` ENUM('Pending', 'Sent', 'Delivered', 'Failed') NOT NULL DEFAULT 'Pending';
-- +goose StatementEnd