-- +goose Up
-- +goose StatementBegin
DROP TABLE IF EXISTS pre_go_order_052025_99999;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `pre_go_order_052025_99999` (
    `id` INT(11) NOT NULL AUTO_INCREMENT COMMENT 'Primary key',
    `order_number` VARCHAR(50) NOT NULL COMMENT 'Order number',
    `user_id` BIGINT(11) NOT NULL COMMENT 'User ID',
    `station_code` VARCHAR(50) NOT NULL COMMENT 'Station code',
    `order_amount` decimal(8, 2) NOT NULL COMMENT 'Order amount',
    `terminal_id` BIGINT(11) NOT NULL COMMENT 'Terminal ID',
    `order_date` Timestamp NOT NULL COMMENT 'Order date',
    `order_notes` VARCHAR(255) NOT NULL COMMENT 'Order notes',
    `order_item` JSON NOT NULL COMMENT 'Order items',
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Last update time',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Creation time',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `order_number` (`order_number`),
    KEY `order_date` (`order_date`),
    KEY `user_id` (`user_id`),
    KEY `order_notes` (`order_notes`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = 'Table for ticket orders';
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS pre_go_order_052025_99999;
-- +goose StatementEnd