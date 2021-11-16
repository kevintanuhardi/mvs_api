-- +goose Up
CREATE TABLE `area_coverage`
(
    `id`            int(11)      NOT NULL AUTO_INCREMENT,
    `employee_id`   int(10)      NOT NULL,
    `company_id`   int(10)      NOT NULL,
    `area_coverage_type`   varchar(20)      NOT NULL,
    `area_coverage`   varchar(50)      NOT NULL,
    `channel_coverage`   varchar(50)      NOT NULL,
    `created_at`    Timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`    Timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
)
    CHARACTER SET = utf8
    COLLATE = utf8_general_ci
    ENGINE = InnoDB;

-- +goose Down
DROP TABLE IF EXISTS `area_coverage`;
