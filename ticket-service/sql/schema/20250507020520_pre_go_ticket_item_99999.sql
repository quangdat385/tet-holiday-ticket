-- +goose Up
-- +goose StatementBegin
DROP TABLE IF EXISTS `pre_go_ticket_item_99999`;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `pre_go_ticket_item_99999` (
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT COMMENT 'Primary key',
    `name` VARCHAR(50) NOT NULL COMMENT 'Ticket title',
    `ticket_id` BIGINT(20) NOT NULL COMMENT 'ID of the associated ticket',
    `train_id` BIGINT(20) NOT NULL COMMENT 'ID of the associated train',
    `description` TEXT COMMENT 'Ticket description',
    `seat_class` ENUM('ECONOMY', 'BUSINESS', 'FIRST') NOT NULL DEFAULT 'ECONOMY' COMMENT 'Seat class (e.g., ECONOMY, BUSINESS, FIRST)',
    `stock_initial` INT(11) NOT NULL DEFAULT 0 COMMENT 'Initial stock quantity (e.g., 1000 tickets)',
    `stock_available` INT(11) NOT NULL DEFAULT 0 COMMENT 'Current available stock (e.g., 900 tickets)',
    `is_stock_prepared` BOOLEAN NOT NULL DEFAULT 0 COMMENT 'Indicates if stock is pre-warmed (0/1)',
    `departure_time` TIMESTAMP NOT NULL COMMENT 'Departure time of the train',
    -- warm up cache
    `price_original` decimal(8, 2) NOT NULL COMMENT 'Original ticket price',
    -- Giá gốc: ví dụ: 100K/ticket
    `price_flash` decimal(8, 2) NOT NULL COMMENT 'Discounted price during flash sale',
    -- Giảm giá khung giờ vàng : ví dụ: 10K/ticket
    `sale_start_time` TIMESTAMP NOT NULL COMMENT 'Flash sale start time',
    `sale_end_time` TIMESTAMP NOT NULL COMMENT 'Flash sale end time',
    `status` INT(11) NOT NULL DEFAULT 0 COMMENT 'Ticket status (e.g., active/inactive)',
    -- Trạng thái của vé (ví dụ: hoạt động/không hoạt động)
    `activity_id` BIGINT(20) NOT NULL COMMENT 'ID of associated activity',
    -- ID của hoạt động liên quan đến vé
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Timestamp of the last update',
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Creation timestamp',
    PRIMARY KEY (`id`),
    KEY `idx_end_time` (`sale_end_time`),
    KEY `idx_start_time` (`sale_start_time`),
    KEY `idx_status` (`status`),
    FOREIGN KEY (`ticket_id`) REFERENCES `pre_go_ticket_99999` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = 'Table for ticket details';
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `pre_go_ticket_item_99999`;
-- +goose StatementEnd