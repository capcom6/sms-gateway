-- +goose Up
-- +goose StatementBegin
CREATE TABLE `message_states` (
    `id` BIGINT UNSIGNED AUTO_INCREMENT,
    `message_id` BIGINT UNSIGNED NOT NULL,
    `state` enum(
        'Pending',
        'Sent',
        'Processed',
        'Delivered',
        'Failed'
    ) NOT NULL,
    `updated_at` datetime(3) NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE INDEX `unq_message_states_message_id_state` (`message_id`, `state`),
    CONSTRAINT `fk_messages_states` FOREIGN KEY (`message_id`) REFERENCES `messages`(`id`) ON DELETE CASCADE
);
-- +goose StatementEnd
---
-- +goose Down
-- +goose StatementBegin
DROP TABLE `message_states`;
-- +goose StatementEnd