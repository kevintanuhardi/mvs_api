-- +goose Up
CREATE TABLE `user`
(
    `id`            int(11)      NOT NULL AUTO_INCREMENT,
    `employee_id`   varchar(40)      NOT NULL,
    `company_id`		int(11) NOT NULL,
    `active`        boolean			  NOT NULL,
    `phone_number`	varchar(25)			  NOT NULL,
    `email`					varchar(40)			  NOT NULL,
    `password`			varchar(255)			  NOT NULL,
    `role_id`			int(11)			  NOT NULL,
    `created_at`    Timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`    Timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
)
    CHARACTER SET = utf8
    COLLATE = utf8_general_ci
    ENGINE = InnoDB;

-- +goose Down
DROP TABLE IF EXISTS `user`;
