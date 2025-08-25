-- +goose Up
-- +goose StatementBegin
DROP TABLE IF EXISTS pre_go_communication_message_read_99999;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE TABLE pre_go_communication_message_read_99999 (
    message_id BIGINT,
    user_id BIGINT,
    read_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (message_id, user_id)
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS pre_go_communication_message_read_99999;
-- +goose StatementEnd