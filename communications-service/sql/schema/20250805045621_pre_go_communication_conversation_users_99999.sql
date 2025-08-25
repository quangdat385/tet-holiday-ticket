-- +goose Up
-- +goose StatementBegin
DROP TABLE IF EXISTS pre_go_communication_conversation_users_99999;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS pre_go_communication_conversation_users_99999 (
    id BIGINT NOT NULL AUTO_INCREMENT COMMENT 'Primary key',
    conversation_id BIGINT NOT NULL COMMENT 'Conversation ID',
    user_id BIGINT NOT NULL COMMENT 'User ID',
    nick_name VARCHAR(255) DEFAULT '' COMMENT 'User Nickname',
    is_deleted BOOLEAN DEFAULT FALSE COMMENT 'Is the user deleted from the conversation',
    last_message TEXT DEFAULT null COMMENT 'Last message sent by the user in the conversation',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (conversation_id) REFERENCES pre_go_communication_conversation_99999(id) ON DELETE CASCADE,
    INDEX idx_is_deleted (is_deleted)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = 'Pre-go communication conversation users table';
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS pre_go_communication_conversation_users_99999;
-- +goose StatementEnd