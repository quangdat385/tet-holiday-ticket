-- +goose Up
-- +goose StatementBegin
DROP TABLE IF EXISTS pre_go_communication_notification_user_99999;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE TABLE pre_go_communication_notification_user_99999 (
    user_id BIGINT NOT NULL,
    notification_id BIGINT NOT NULL,
    read_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, notification_id)
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS pre_go_communication_notification_user_99999;
-- +goose StatementEnd