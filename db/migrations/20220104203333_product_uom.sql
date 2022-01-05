-- +goose Up
CREATE TABLE `product_uom`
(
    `id`            int(11)      NOT NULL AUTO_INCREMENT,
		`name` varchar(40) NOT NULL,
		`product_id` int(11) NOT NULL,
		`conversion` int(11) NOT NULL,
		`is_sale` boolean NOT NULL DEFAULT true,
    `created_at`    Timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`    Timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
)
    CHARACTER SET = utf8
    COLLATE = utf8_general_ci
    ENGINE = InnoDB;

-- +goose Down
DROP TABLE IF EXISTS `product_uom`;
