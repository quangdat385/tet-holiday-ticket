-- +goose Up
-- +goose StatementBegin
DROP TABLE IF EXISTS pre_go_ticket_segment_price_99999;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS pre_go_ticket_segment_price_99999 (
    id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT 'Primary key',
    ticket_item_id BIGINT NOT NULL COMMENT 'ID of the associated ticket item',
    route_segment_id BIGINT NOT NULL COMMENT 'ID of the associated route segment',
    price DECIMAL(8, 2) NOT NULL COMMENT 'Price for this segment',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (ticket_item_id) REFERENCES pre_go_ticket_item_99999(id),
    FOREIGN KEY (route_segment_id) REFERENCES pre_go_route_segment_99999(id)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = 'Table for ticket segment prices';
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS pre_go_ticket_segment_price_99999;
-- +goose StatementEnd