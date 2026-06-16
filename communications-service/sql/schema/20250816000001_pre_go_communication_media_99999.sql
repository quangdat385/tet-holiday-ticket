-- +goose Up
-- +goose StatementBegin
DROP TABLE IF EXISTS pre_go_communication_media_99999;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS pre_go_communication_media_99999 (
    id BIGINT NOT NULL AUTO_INCREMENT COMMENT 'Primary key for the media table',
    message_id BIGINT NULL COMMENT 'Associated message ID',
    conversation_id BIGINT NULL COMMENT 'Associated conversation ID',
    user_id BIGINT NOT NULL COMMENT 'Uploader user ID',
    file_name VARCHAR(255) NOT NULL COMMENT 'Stored file name (UUID-based)',
    original_name VARCHAR(255) NOT NULL COMMENT 'Original file name',
    mime_type VARCHAR(100) NOT NULL COMMENT 'MIME type of the file',
    size BIGINT NOT NULL COMMENT 'File size in bytes',
    url VARCHAR(500) NOT NULL COMMENT 'Accessible URL of the file',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    INDEX idx_message_id (message_id),
    INDEX idx_conversation_id (conversation_id),
    INDEX idx_user_id (user_id)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = 'ticket pre-go communication media table';
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS pre_go_communication_media_99999;
-- +goose StatementEnd
