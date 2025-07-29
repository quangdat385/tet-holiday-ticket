-- +goose Up
-- +goose StatementBegin
DROP TABLE IF EXISTS pre_go_communication_notification_99999;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS pre_go_communication_notification_99999 (
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `from` BIGINT NOT NULL,
    `to` JSON NULL DEFAULT NULL,
    content JSON NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_from (`from`),
    INDEX idx_to (`to`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = 'notification table';
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS pre_go_communication_notification_99999;
-- +goose StatementEnd