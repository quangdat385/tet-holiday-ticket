-- +goose Up
-- +goose StatementBegin
DROP TABLE IF EXISTS pre_go_communication_conversation_99999;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS pre_go_communication_conversation_99999 (
    id BIGINT NOT NULL AUTO_INCREMENT COMMENT 'Primary key',
    title VARCHAR(255) DEFAULT NULL,
    description VARCHAR(255) DEFAULT '',
    type VARCHAR(50) DEFAULT 'personal',
    background VARCHAR(50) DEFAULT 'white',
    emoji VARCHAR(50) DEFAULT 'haha',
    is_deleted BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    INDEX idx_is_deleted (is_deleted)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = 'Pre-go communication conversation table';
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS pre_go_communication_conversation_99999;
-- +goose StatementEnd