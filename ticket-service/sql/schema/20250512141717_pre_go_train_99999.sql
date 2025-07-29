-- +goose Up
-- +goose StatementBegin
DROP TABLE IF EXISTS pre_go_train_99999;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE TABLE pre_go_train_99999 (
    id BIGINT NOT NULL AUTO_INCREMENT COMMENT 'Primary key',
    code VARCHAR(255) NOT NULL COMMENT 'Train number',
    name VARCHAR(255) NOT NULL COMMENT 'Train name',
    departure_station_id BIGINT NOT NULL COMMENT 'Departure station ID',
    arrival_station_id BIGINT NOT NULL COMMENT 'Arrival station ID',
    departure_time TIMESTAMP NOT NULL COMMENT 'Departure time',
    arrival_time TIMESTAMP NOT NULL COMMENT 'Arrival time',
    status INT NOT NULL DEFAULT 0 COMMENT 'Train status (e.g., on time, delayed)',
    direction ENUM('SOUTH', 'NORTH') NOT NULL COMMENT 'Train direction (e.g., SOUTH, NORTH)',
    train_type ENUM('EXPRESS', 'LOCAL') NOT NULL COMMENT 'Train type (e.g., EXPRESS, LOCAL)',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE KEY departure_time (departure_time),
    UNIQUE KEY arrival_time (arrival_time)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = 'Table for train details';
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS pre_go_train_99999;
-- +goose StatementEnd