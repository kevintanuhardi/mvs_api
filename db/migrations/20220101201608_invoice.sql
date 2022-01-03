-- +goose Up
CREATE TABLE `invoice`
(
    `id`            int(11)      NOT NULL AUTO_INCREMENT,
		`invoice_no`		int(11)			NOT NULL,
		`invoice_date`	Timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
		`payment_period_in_days` int(3) NOT NULL,
		`shipping_contact_name` varchar(40) NOT NULL,
		`shipping_contact_phone` varchar(40) NOT NULL,
		`shipping_address` varchar(200) NOT NULL,
		`shipping_city` varchar(40) NOT NULL,
		`shipping_postal_code` varchar(10) NOT NULL,
		`billing_contact_name` varchar(40),
		`billing_contact_phone` varchar(40),
		`billing_address` varchar(200),
		`billing_city` varchar(40),
		`billing_postal_code` varchar(10),
		`status`			enum('unpaid', 'paid', 'cancelled')	NOT NULL 	DEFAULT "unpaid",
    `created_at`    Timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`    Timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
)
    CHARACTER SET = utf8
    COLLATE = utf8_general_ci
    ENGINE = InnoDB;

-- +goose Down
DROP TABLE IF EXISTS `invoice`;
