-- +goose Up
CREATE TABLE `otp`
(
    `id`            int(11)      	NOT NULL AUTO_INCREMENT,
    `owner_id`   	varchar(40)  	NOT NULL,
    `otp`   		varchar(4)   	NOT NULL,
    `type`   		varchar(10)     NOT NULL,
    `exp_time`   	Timestamp     NOT NULL,
    `created_at`    Timestamp    	NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`    Timestamp    	NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
)
    CHARACTER SET = utf8
    COLLATE = utf8_general_ci
    ENGINE = InnoDB;

-- +goose Down
DROP TABLE IF EXISTS `otp`;
