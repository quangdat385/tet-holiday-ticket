CREATE DATABASE IF NOT EXISTS vetautet DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;
-- 1. ticket table
CREATE TABLE IF NOT EXISTS `vetautet`.`ticket` (
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT COMMENT 'Primary key',
    `name` VARCHAR(50) NOT NULL COMMENT 'ticket name',
    `description` TEXT COMMENT 'ticket description',
    `start_time` DATETIME NOT NULL COMMENT 'ticket sale start time',
    `end_time` DATETIME NOT NULL COMMENT 'ticket sale end time',
    `status` INT(11) NOT NULL DEFAULT 0 COMMENT 'ticket sale activity status',
    -- 0: deactive, 1: activity
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Last update time',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Creation time',
    PRIMARY KEY (`id`),
    KEY `idx_end_time` (`end_time`),
    -- Very high query runtime
    KEY `idx_start_time` (`start_time`),
    -- Very high query runtime
    KEY `idx_status` (`status`) -- Very high query runtime
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = 'ticket table';
-- 2. ticket detail (item) table
CREATE TABLE IF NOT EXISTS `vetautet`.`ticket_item` (
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT COMMENT 'Primary key',
    `name` VARCHAR(50) NOT NULL COMMENT 'Ticket title',
    `description` TEXT COMMENT 'Ticket description',
    `stock_initial` INT(11) NOT NULL DEFAULT 0 COMMENT 'Initial stock quantity (e.g., 1000 tickets)',
    `stock_available` INT(11) NOT NULL DEFAULT 0 COMMENT 'Current available stock (e.g., 900 tickets)',
    `is_stock_prepared` BOOLEAN NOT NULL DEFAULT 0 COMMENT 'Indicates if stock is pre-warmed (0/1)',
    -- warm up cache
    `price_original` BIGINT(20) NOT NULL COMMENT 'Original ticket price',
    -- Giá gốc: ví dụ: 100K/ticket
    `price_flash` BIGINT(20) NOT NULL COMMENT 'Discounted price during flash sale',
    -- Giảm giá khung giờ vàng : ví dụ: 10K/ticket
    `sale_start_time` DATETIME NOT NULL COMMENT 'Flash sale start time',
    `sale_end_time` DATETIME NOT NULL COMMENT 'Flash sale end time',
    `status` INT(11) NOT NULL DEFAULT 0 COMMENT 'Ticket status (e.g., active/inactive)',
    -- Trạng thái của vé (ví dụ: hoạt động/không hoạt động)
    `activity_id` BIGINT(20) NOT NULL COMMENT 'ID of associated activity',
    -- ID của hoạt động liên quan đến vé
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Timestamp of the last update',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Creation timestamp',
    PRIMARY KEY (`id`),
    KEY `idx_end_time` (`sale_end_time`),
    KEY `idx_start_time` (`sale_start_time`),
    KEY `idx_status` (`status`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = 'Table for ticket details';
-- 3. ticket order table
CREATE TABLE IF NOT EXISTS `vetautet`.`ticket_order_202502` (
    `id` INT(11) NOT NULL AUTO_INCREMENT COMMENT 'Primary key',
    `order_number` VARCHAR(50) NOT NULL COMMENT 'Order number',
    `order_amount` BIGINT(20) NOT NULL COMMENT 'Order amount',
    `terminal_id` INT(11) NOT NULL COMMENT 'Terminal ID',
    `order_date` Timestamp NOT NULL COMMENT 'Order date',
    `order_notes` VARCHAR(255) COMMENT 'Order notes',
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Last update time',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Creation time',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `order_number` (`order_number`),
    KEY `order_date` (`order_date`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = 'Table for ticket orders';
-- 4. ticket order detail table
CREATE TABLE IF NOT EXISTS `vetautet`.`ticket_order_detail_202502` (
    `id` INT(11) NOT NULL AUTO_INCREMENT COMMENT 'Primary key',
    `ticket_item_id` BIGINT(20) NOT NULL COMMENT 'Ticket item ID',
    `order_number` VARCHAR(50) NOT NULL COMMENT 'Reference order number',
    `passenger_name` VARCHAR(50) NOT NULL COMMENT 'Passenger name',
    `departure_station` VARCHAR(50) NOT NULL COMMENT 'Departure station',
    `arrival_station` VARCHAR(50) NOT NULL COMMENT 'Arrival station',
    `departure_time` DATETIME NOT NULL COMMENT 'Departure time',
    `passenger_id` VARCHAR(20) NOT NULL COMMENT 'Passenger ID',
    `seat_class` ENUM('ECONOMY', 'BUSINESS', 'FIRST') NOT NULL COMMENT 'Seat class',
    `ticket_price` DECIMAL(10, 3) NOT NULL COMMENT 'Ticket price',
    `seat_number` VARCHAR(10) NOT NULL COMMENT 'Seat number',
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Last update time',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Creation time',
    PRIMARY KEY (`id`) USING BTREE,
    KEY `order_id` (`order_id`),
    KEY `ticket_item_id` (`ticket_item_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = 'Table for ticket order details';
-- INSERT MOCK DATA
-- Insert data into `ticket` table
INSERT INTO `vetautet`.`ticket` (
        `name`,
        `description`,
        `start_time`,
        `end_time`,
        `status`,
        `updated_at`,
        `created_at`
    )
VALUES (
        'Đợt Mở Bán Vé Ngày 12/12',
        'Sự kiện mở bán vé đặc biệt cho ngày 12/12',
        '2024-12-12 00:00:00',
        '2024-12-12 23:59:59',
        1,
        CURRENT_TIMESTAMP,
        CURRENT_TIMESTAMP
    ),
    (
        'Đợt Mở Bán Vé Ngày 01/01',
        'Sự kiện mở bán vé cho ngày đầu năm mới 01/01',
        '2025-01-01 00:00:00',
        '2025-01-01 23:59:59',
        1,
        CURRENT_TIMESTAMP,
        CURRENT_TIMESTAMP
    );
-- Insert data into `ticket_item` table corresponding to each event in `ticket` table
INSERT INTO `vetautet`.`ticket_item` (
        `name`,
        `description`,
        `stock_initial`,
        `stock_available`,
        `is_stock_prepared`,
        `price_original`,
        `price_flash`,
        `sale_start_time`,
        `sale_end_time`,
        `status`,
        `activity_id`,
        `updated_at`,
        `created_at`
    )
VALUES -- Ticket items for the 12/12 event
    (
        'Vé Sự Kiện 12/12 - Hạng Phổ Thông',
        'Vé phổ thông cho sự kiện ngày 12/12',
        1000,
        1000,
        0,
        100000,
        10000,
        '2024-12-12 00:00:00',
        '2024-12-12 23:59:59',
        1,
        1,
        CURRENT_TIMESTAMP,
        CURRENT_TIMESTAMP
    ),
    (
        'Vé Sự Kiện 12/12 - Hạng VIP',
        'Vé VIP cho sự kiện ngày 12/12',
        500,
        500,
        0,
        200000,
        15000,
        '2024-12-12 00:00:00',
        '2024-12-12 23:59:59',
        1,
        1,
        CURRENT_TIMESTAMP,
        CURRENT_TIMESTAMP
    ),
    -- Ticket items for the 01/01 event
    (
        'Vé Sự Kiện 01/01 - Hạng Phổ Thông',
        'Vé phổ thông cho sự kiện ngày 01/01',
        2000,
        2000,
        0,
        100000,
        10000,
        '2025-01-01 00:00:00',
        '2025-01-01 23:59:59',
        1,
        2,
        CURRENT_TIMESTAMP,
        CURRENT_TIMESTAMP
    ),
    (
        'Vé Sự Kiện 01/01 - Hạng VIP',
        'Vé VIP cho sự kiện ngày 01/01',
        1000,
        1000,
        0,
        200000,
        15000,
        '2025-01-01 00:00:00',
        '2025-01-01 23:59:59',
        1,
        2,
        CURRENT_TIMESTAMP,
        CURRENT_TIMESTAMP
    );
-- 1. Danh sách ga
CREATE TABLE station (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL
);
-- 2. Chuyến tàu
CREATE TABLE train (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    code VARCHAR(20) NOT NULL,
    name VARCHAR(100) NOT NULL,
    departure_station_id BIGINT NOT NULL,
    arrival_station_id BIGINT NOT NULL,
    departure_time DATETIME NOT NULL,
    arrival_time DATETIME NOT NULL,
    direction VARCHAR(20) NOT NULL,
    FOREIGN KEY (departure_station_id) REFERENCES station(id),
    FOREIGN KEY (arrival_station_id) REFERENCES station(id)
);
-- 3. Đợt mở bán vé
CREATE TABLE ticket (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(50) NOT NULL,
    description TEXT,
    start_time DATETIME NOT NULL,
    end_time DATETIME NOT NULL,
    status TINYINT NOT NULL DEFAULT 0,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- 4. Loại vé cụ thể
CREATE TABLE ticket_item (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    ticket_id BIGINT NOT NULL,
    train_id BIGINT NOT NULL,
    seat_class ENUM('ECONOMY', 'BUSINESS', 'FIRST') NOT NULL,
    price_original BIGINT NOT NULL,
    price_flash BIGINT NOT NULL,
    sale_start_time DATETIME NOT NULL,
    sale_end_time DATETIME NOT NULL,
    stock_initial INT NOT NULL DEFAULT 0,
    stock_available INT NOT NULL DEFAULT 0,
    is_stock_prepared BOOLEAN NOT NULL DEFAULT 0,
    status TINYINT NOT NULL DEFAULT 0,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (ticket_id) REFERENCES ticket(id),
    FOREIGN KEY (train_id) REFERENCES train(id)
);
-- 5. Danh sách ghế
CREATE TABLE seat (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    train_id BIGINT NOT NULL,
    seat_number VARCHAR(10) NOT NULL,
    seat_class ENUM('ECONOMY', 'BUSINESS', 'FIRST') NOT NULL,
    carriage_number INT NOT NULL,
    FOREIGN KEY (train_id) REFERENCES train(id)
);
-- 6. Đơn hàng
CREATE TABLE ticket_order (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    user_id BIGINT NOT NULL,
    order_number VARCHAR(50) NOT NULL UNIQUE,
    order_amount BIGINT NOT NULL,
    order_date DATETIME NOT NULL,
    status TINYINT NOT NULL DEFAULT 0,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- 7. Chi tiết đơn hàng
CREATE TABLE ticket_order_detail (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    ticket_order_id BIGINT NOT NULL,
    ticket_item_id BIGINT NOT NULL,
    seat_id BIGINT NOT NULL,
    passenger_name VARCHAR(50) NOT NULL,
    passenger_id VARCHAR(20) NOT NULL,
    from_station_id BIGINT NOT NULL,
    to_station_id BIGINT NOT NULL,
    departure_time DATETIME NOT NULL,
    ticket_price BIGINT NOT NULL,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (ticket_order_id) REFERENCES ticket_order(id),
    FOREIGN KEY (ticket_item_id) REFERENCES ticket_item(id),
    FOREIGN KEY (seat_id) REFERENCES seat(id),
    FOREIGN KEY (from_station_id) REFERENCES station(id),
    FOREIGN KEY (to_station_id) REFERENCES station(id)
);
-- 8. Đặt chỗ ghế theo đoạn
CREATE TABLE seat_reservation (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    seat_id BIGINT NOT NULL,
    train_id BIGINT NOT NULL,
    ticket_order_detail_id BIGINT NOT NULL,
    from_station_id BIGINT NOT NULL,
    to_station_id BIGINT NOT NULL,
    FOREIGN KEY (seat_id) REFERENCES seat(id),
    FOREIGN KEY (train_id) REFERENCES train(id),
    FOREIGN KEY (ticket_order_detail_id) REFERENCES ticket_order_detail(id),
    FOREIGN KEY (from_station_id) REFERENCES station(id),
    FOREIGN KEY (to_station_id) REFERENCES station(id)
);
-- 9. Đoạn lộ trình
CREATE TABLE route_segment (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    train_id BIGINT NOT NULL,
    from_station_id BIGINT NOT NULL,
    to_station_id BIGINT NOT NULL,
    segment_order INT NOT NULL,
    distance_km INT NOT NULL,
    FOREIGN KEY (train_id) REFERENCES train(id),
    FOREIGN KEY (from_station_id) REFERENCES station(id),
    FOREIGN KEY (to_station_id) REFERENCES station(id)
);
INSERT INTO pre_go_route_segment_99999 (
        train_id,
        from_station_id,
        to_station_id,
        segment_order,
        distance_km
    )
VALUES (1, 1, 2, 1, 56),
    -- Hà Nội -> Hà Nam
    (1, 2, 3, 2, 61),
    -- Hà Nam -> Nam Định
    (1, 3, 4, 3, 32),
    -- Nam Định -> Ninh Bình
    (1, 4, 5, 4, 93),
    -- Ninh Bình -> Thanh Hóa
    (1, 5, 6, 5, 139),
    -- Thanh Hóa -> Nghệ An
    (1, 6, 7, 6, 92),
    -- Nghệ An -> Hà Tĩnh
    (1, 7, 8, 7, 130),
    -- Hà Tĩnh -> Quảng Bình
    (1, 8, 9, 8, 110),
    -- Quảng Bình -> Quảng Trị
    (1, 9, 10, 9, 60),
    -- Quảng Trị -> Thừa Thiên Huế
    (1, 10, 11, 10, 103),
    -- Thừa Thiên Huế -> Đà Nẵng
    (1, 11, 12, 11, 87),
    -- Đà Nẵng -> Quảng Nam
    (1, 12, 13, 12, 91),
    -- Quảng Nam -> Quảng Ngãi
    (1, 13, 14, 13, 119),
    -- Quảng Ngãi -> Bình Định
    (1, 14, 15, 14, 92),
    -- Bình Định -> Phú Yên
    (1, 15, 16, 15, 120),
    -- Phú Yên -> Khánh Hòa
    (1, 16, 17, 16, 60),
    -- Khánh Hòa -> Ninh Thuận
    (1, 17, 18, 17, 105),
    -- Ninh Thuận -> Bình Thuận
    (1, 18, 19, 18, 150),
    -- Bình Thuận -> Đồng Nai
    (1, 19, 20, 19, 90),
    -- Đồng Nai -> Bình Dương
    (1, 20, 21, 20, 40);
-- Bình Dương -> TP Hồ Chí Minh
INSERT INTO pre_go_station_99999 (id, name, code, status)
VALUES (1, 'Hà Nội', 'OKX-HNI', 1),
    (2, 'Hà Nam', 'OKX-HNA', 1),
    (3, 'Nam Định', 'OKX-NDI', 1),
    (4, 'Ninh Bình', 'OKX-NBI', 1),
    (5, 'Thanh Hóa', 'OKX-THO', 1),
    (6, 'Nghệ An', 'OKX-NAN', 1),
    (7, 'Hà Tĩnh', 'OKX-HTH', 1),
    (8, 'Quảng Bình', 'OKX-QBI', 1),
    (9, 'Quảng Trị', 'OKX-QTR', 1),
    (10, 'Thừa Thiên Huế', 'OKX-TTH', 1),
    (11, 'Đà Nẵng', 'OKX-DNA', 1),
    (12, 'Quảng Nam', 'OKX-QNA', 1),
    (13, 'Quảng Ngãi', 'OKX-QNG', 1),
    (14, 'Bình Định', 'OKX-BDI', 1),
    (15, 'Phú Yên', 'OKX-PYE', 1),
    (16, 'Khánh Hòa', 'OKX-KHO', 1),
    (17, 'Ninh Thuận', 'OKX-NTH', 1),
    (18, 'Bình Thuận', 'OKX-BTH', 1),
    (19, 'Đồng Nai', 'OKX-DNAI', 1),
    (20, 'Bình Dương', 'OKX-BDU', 1),
    (21, 'TP Hồ Chí Minh', 'OKX-HCM', 1);