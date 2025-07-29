-- +goose Up
-- +goose StatementBegin
DROP TABLE IF EXISTS pre_go_seat_reservation_99999;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS pre_go_seat_reservation_99999 (
    id BIGINT NOT NULL AUTO_INCREMENT COMMENT 'Primary key',
    train_id BIGINT NOT NULL COMMENT 'Train ID',
    seat_id BIGINT NOT NULL COMMENT 'Seat ID',
    order_number VARCHAR(255) NOT NULL COMMENT 'Order number',
    from_station_id BIGINT NOT NULL COMMENT 'Departure station ID',
    to_station_id BIGINT NOT NULL COMMENT 'Arrival station ID',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id) COMMENT 'Primary key',
    UNIQUE KEY (order_number) COMMENT 'Unique order number',
    FOREIGN KEY (train_id) REFERENCES pre_go_train_99999(id),
    FOREIGN KEY (seat_id) REFERENCES pre_go_seat_99999(id),
    FOREIGN KEY (from_station_id) REFERENCES pre_go_station_99999(id),
    FOREIGN KEY (to_station_id) REFERENCES pre_go_station_99999(id)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = 'Table for seat details';
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS pre_go_seat_reservation_99999;
-- +goose StatementEnd