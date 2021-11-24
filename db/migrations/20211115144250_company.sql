-- +goose Up
CREATE TABLE `company`
(
    `id`            int(11)      NOT NULL AUTO_INCREMENT,
    `code`          varchar(10)  NOT NULL,
    `name`		    varchar(100) NOT NULL,
    `address`       varchar(250) NOT NULL,
    `country`		varchar(100) NOT NULL,
    `province`		varchar(100) NOT NULL,
    `city`		    varchar(100) NOT NULL,
    `district`		varchar(100) NOT NULL,
    `village`		varchar(100) NOT NULL,
    `postal_code`	varchar(100) NOT NULL,
    `phone_number`	varchar(20)  NOT NULL,
    `fax_number`    varchar(20)  NOT NULL,
    `npwp`          varchar(20)  NOT NULL,
    `sppkp`         varchar(20)  NOT NULL,
    `created_at`    Timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`    Timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
)
    CHARACTER SET = utf8
    COLLATE = utf8_general_ci
    ENGINE = InnoDB;

-- +goose Down
DROP TABLE IF EXISTS `company`;
