-- +goose Up
-- +goose StatementBegin
DROP TABLE IF EXISTS pre_go_seat_99999;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE TABLE pre_go_seat_99999 (
    id BIGINT NOT NULL AUTO_INCREMENT COMMENT 'Primary key',
    train_id BIGINT NOT NULL COMMENT 'Train ID',
    seat_number VARCHAR(255) NOT NULL COMMENT 'Seat number',
    seat_class ENUM('ECONOMY', 'BUSINESS', 'FIRST') NOT NULL,
    status INT NOT NULL DEFAULT 0 COMMENT 'Seat status (e.g., available, booked)',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id) COMMENT 'Primary key',
    UNIQUE KEY uniq_train_seat (train_id, seat_number),
    KEY idx_train_id (train_id),
    KEY idx_seat_number (seat_number),
    FOREIGN KEY (train_id) REFERENCES pre_go_train_99999(id) ON DELETE CASCADE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = 'Table for seat details';
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS pre_go_seat_99999;
-- +goose StatementEnd