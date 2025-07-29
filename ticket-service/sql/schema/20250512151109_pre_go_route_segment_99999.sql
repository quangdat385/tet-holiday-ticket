-- +goose Up
-- +goose StatementBegin
DROP TABLE IF EXISTS pre_go_route_segment_99999;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS pre_go_route_segment_99999 (
    id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT 'Primary key',
    train_id BIGINT NOT NULL COMMENT 'ID of the associated train',
    from_station_id BIGINT NOT NULL COMMENT 'ID of the departure station',
    to_station_id BIGINT NOT NULL COMMENT 'ID of the arrival station',
    segment_order INT NOT NULL COMMENT 'Order of the segment in the route',
    distance_km INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (train_id) REFERENCES pre_go_train_99999(id),
    FOREIGN KEY (from_station_id) REFERENCES pre_go_station_99999(id),
    FOREIGN KEY (to_station_id) REFERENCES pre_go_station_99999(id)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = 'Table for route segments';
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS pre_go_route_segment_99999;
-- +goose StatementEnd