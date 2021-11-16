-- +goose Up
CREATE TABLE `sale_order`
(
    `id`            int(11)      NOT NULL AUTO_INCREMENT,
    `customer_id`   int(11)      NOT NULL,
    `trx_id`        varchar(140) NOT NULL,
    `status`        varchar(40)  NOT NULL,
    `created_at`    Timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`    Timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
)
    CHARACTER SET = utf8
    COLLATE = utf8_general_ci
    ENGINE = InnoDB;

-- +goose Down
DROP TABLE IF EXISTS `sale_order`;
