-- +goose Up
-- +goose StatementBegin
DROP TABLE IF EXISTS pre_go_station_99999;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE TABLE pre_go_station_99999 (
    id BIGINT NOT NULL AUTO_INCREMENT COMMENT 'Primary key',
    name VARCHAR(255) NOT NULL COMMENT 'Station name',
    code VARCHAR(255) NOT NULL COMMENT 'Station code',
    status INT NOT NULL DEFAULT 1 COMMENT 'Station status (e.g., active/inactive)',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE KEY (code)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = 'Table for station details';
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS pre_go_station_99999;
-- +goose StatementEnd