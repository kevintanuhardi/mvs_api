-- +goose Up
CREATE TABLE `product`
(
    `id`            int(11)      NOT NULL AUTO_INCREMENT,
		`name` varchar(40) NOT NULL,
		`category_id` int(11),
		`sku_no` varchar(60) NOT NULL, 
    `created_at`    Timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`    Timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
)
    CHARACTER SET = utf8
    COLLATE = utf8_general_ci
    ENGINE = InnoDB;

-- +goose Down
DROP TABLE IF EXISTS `product`;
