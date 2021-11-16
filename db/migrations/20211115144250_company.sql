-- +goose Up
CREATE TABLE `company`
(
    `id`            int(11)      NOT NULL AUTO_INCREMENT,
    `code`   varchar(10)      NOT NULL,
    `name`		varchar(100) NOT NULL,
    `created_at`    Timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`    Timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
)
    CHARACTER SET = utf8
    COLLATE = utf8_general_ci
    ENGINE = InnoDB;

-- +goose Down
DROP TABLE IF EXISTS `company`;
