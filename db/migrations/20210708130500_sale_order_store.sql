-- +goose Up
CREATE TABLE `sale_order_store`
(
    `id`            int(11)      NOT NULL AUTO_INCREMENT,
    `order_id`      int(11)      NOT NULL,
    `store_id`      int(11)      NOT NULL,
    `store_trx_id`  varchar(140) NOT NULL,
    `created_at`    Timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`    Timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `order_store_ibfk_1` (`order_id`),
    CONSTRAINT `sale_order_store_ibfk_1` FOREIGN KEY (`order_id`) REFERENCES `sale_order` (`id`) ON DELETE CASCADE
)
    CHARACTER SET = utf8
    COLLATE = utf8_general_ci
    ENGINE = InnoDB;

-- +goose Down
DROP TABLE IF EXISTS `sale_order_store`;
