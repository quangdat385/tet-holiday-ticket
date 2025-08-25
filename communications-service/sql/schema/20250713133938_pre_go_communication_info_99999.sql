-- +goose Up
-- +goose StatementBegin
DROP TABLE IF EXISTS pre_go_communication_info_99999;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS pre_go_communication_info_99999 (
    id BIGINT NOT NULL AUTO_INCREMENT COMMENT 'Primary key for the communication info table',
    user_id BIGINT NOT NULL UNIQUE,
    status BOOLEAN DEFAULT FALSE,
    value VARCHAR(255) DEFAULT NULL UNIQUE,
    type VARCHAR(255) DEFAULT 'socket_id',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = 'Pre-go communication info table';
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS pre_go_communication_info_99999;
-- +goose StatementEnd