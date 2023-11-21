-- +goose Up
-- +goose StatementBegin
CREATE TABLE `users` (
    `id` varchar(32),
    `password_hash` varchar(72) NOT NULL,
    `created_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `updated_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    `deleted_at` datetime(3) NULL,
    PRIMARY KEY (`id`)
);
-- +goose StatementEnd
-- +goose StatementBegin
CREATE TABLE `devices` (
    `id` char(21),
    `name` varchar(128),
    `auth_token` char(21) NOT NULL,
    `push_token` varchar(256),
    `user_id` varchar(32) NOT NULL,
    `created_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `updated_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    `deleted_at` datetime(3) NULL,
    PRIMARY KEY (`id`),
    UNIQUE INDEX `idx_devices_auth_token` (`auth_token`),
    CONSTRAINT `fk_users_devices` FOREIGN KEY (`user_id`) REFERENCES `users`(`id`) ON DELETE CASCADE
);
-- +goose StatementEnd
-- +goose StatementBegin
CREATE TABLE `messages` (
    `id` BIGINT UNSIGNED AUTO_INCREMENT,
    `device_id` char(21) NOT NULL,
    `ext_id` varchar(36) NOT NULL,
    `message` tinytext NOT NULL,
    `state` enum('Pending', 'Sent', 'Delivered', 'Failed') NOT NULL DEFAULT 'Pending',
    `created_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `updated_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    `deleted_at` datetime(3) NULL,
    PRIMARY KEY (`id`),
    UNIQUE INDEX `unq_messages_device_id` (`device_id`, `ext_id`),
    INDEX `idx_messages_device_state` (`device_id`, `state`),
    CONSTRAINT `fk_messages_device` FOREIGN KEY (`device_id`) REFERENCES `devices`(`id`) ON DELETE CASCADE
);
-- +goose StatementEnd
-- +goose StatementBegin
CREATE TABLE `message_recipients` (
    `message_id` BIGINT UNSIGNED,
    `phone_number` char(11),
    `state` enum('Pending', 'Sent', 'Delivered', 'Failed') NOT NULL DEFAULT 'Pending',
    PRIMARY KEY (`message_id`, `phone_number`),
    CONSTRAINT `fk_messages_recipients` FOREIGN KEY (`message_id`) REFERENCES `messages`(`id`) ON DELETE CASCADE
);
-- +goose StatementEnd
-------------------------------------------------------------------------------
-- +goose Down
-- +goose StatementBegin
DROP TABLE `message_recipients`;
-- +goose StatementEnd
-- +goose StatementBegin
DROP TABLE `messages`;
-- +goose StatementEnd
-- +goose StatementBegin
DROP TABLE `devices`;
-- +goose StatementEnd
-- +goose StatementBegin
DROP TABLE `users`;
-- +goose StatementEnd