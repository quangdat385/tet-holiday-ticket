-- +goose Up
-- +goose StatementBegin
DROP TABLE IF EXISTS pre_go_communication_message_99999;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS pre_go_communication_message_99999 (
    id BIGINT NOT NULL AUTO_INCREMENT COMMENT 'Primary key for the message table',
    conversation_id BIGINT NOT NULL,
    user_id BIGINT NOT NULL,
    status BOOLEAN DEFAULT FALSE,
    message TEXT NOT NULL,
    type VARCHAR(50) NOT NULL DEFAULT 'message',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    INDEX idx_conversation_id (conversation_id),
    INDEX idx_user_id (user_id),
    FULLTEXT INDEX idx_message_text (message)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = 'ticket pre-go communication message table';
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS pre_go_communication_message_99999;
-- +goose StatementEnd