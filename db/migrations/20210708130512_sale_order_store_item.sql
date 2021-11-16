-- +goose Up
CREATE TABLE `sale_order_store_item`
(
    `id`                int(11)       NOT NULL AUTO_INCREMENT,
    `order_store_id`    int(11)       NOT NULL,
    `sku`               varchar(140)  NOT NULL,
    `name`              varchar(1000) NOT NULL,
    `uom`               varchar(140)  NOT NULL,
    `quantity`          decimal(12,2) NOT NULL,
    `price_unit`        decimal(12,2) NOT NULL,
    `created_at`        Timestamp     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`        Timestamp     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `order_store_item_ibfk_1` (`order_store_id`),
    CONSTRAINT `sale_order_store_item_ibfk_1` FOREIGN KEY (`order_store_id`) REFERENCES `sale_order_store` (`id`) ON DELETE CASCADE
)
    CHARACTER SET = utf8
    COLLATE = utf8_general_ci
    ENGINE = InnoDB;

-- +goose Down
DROP TABLE IF EXISTS `sale_order_store_item`;
