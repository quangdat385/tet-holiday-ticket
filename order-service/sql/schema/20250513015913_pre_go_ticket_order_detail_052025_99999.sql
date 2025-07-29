-- +goose Up
-- +goose StatementBegin
DROP TABLE IF EXISTS pre_go_ticket_order_detail_052025_99999;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS pre_go_ticket_order_detail_052025_99999 (
    `id` INT(11) NOT NULL AUTO_INCREMENT COMMENT 'Primary key',
    `ticket_item_id` BIGINT(20) NOT NULL COMMENT 'Ticket item ID',
    `order_number` VARCHAR(50) NOT NULL COMMENT 'Reference order number',
    `passenger_name` VARCHAR(50) NOT NULL COMMENT 'Passenger name',
    `departure_station` VARCHAR(50) NOT NULL COMMENT 'Departure station',
    `arrival_station` VARCHAR(50) NOT NULL COMMENT 'Arrival station',
    `departure_time` DATETIME NOT NULL COMMENT 'Departure time',
    `passenger_id` BIGINT(20) NOT NULL COMMENT 'Passenger ID',
    `seat_class` ENUM('ECONOMY', 'BUSINESS', 'FIRST') NOT NULL COMMENT 'Seat class',
    `ticket_price` DECIMAL(10, 3) NOT NULL COMMENT 'Ticket price',
    `seat_number` VARCHAR(10) NOT NULL COMMENT 'Seat number',
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Last update time',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Creation time',
    PRIMARY KEY (`id`) USING BTREE,
    KEY `ticket_item_id` (`ticket_item_id`)
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS pre_go_ticket_order_detail_052025_99999;
-- +goose StatementEnd