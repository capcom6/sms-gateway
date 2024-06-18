-- +goose Up
-- +goose StatementBegin
CREATE TABLE `webhooks` (
    `id` BIGINT UNSIGNED AUTO_INCREMENT,
    `ext_id` varchar(36) NOT NULL,
    `user_id` varchar(32) NOT NULL,
    `url` varchar(256) NOT NULL,
    `event` varchar(32) NOT NULL,
    `created_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `updated_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    `deleted_at` datetime(3) NULL,
    PRIMARY KEY (`id`),
    UNIQUE INDEX `unq_webhooks_user_extid` (`user_id`, `ext_id`),
    CONSTRAINT `fk_webhooks_user` FOREIGN KEY (`user_id`) REFERENCES `users`(`id`) ON DELETE CASCADE
);
-- +goose StatementEnd
---
-- +goose Down
-- +goose StatementBegin
DROP TABLE `webhooks`;
-- +goose StatementEnd