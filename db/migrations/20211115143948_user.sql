-- +goose Up
CREATE TABLE `user`
(
    `id`            int(11)      NOT NULL AUTO_INCREMENT,
    `name`   varchar(63)      NOT NULL,
    `is_active`        boolean			  NOT NULL,
    `phone_number`	varchar(25)			  NOT NULL,
    `email`					varchar(40)			  NOT NULL,
    `password`			varchar(255)			  NOT NULL,
    `role`			enum('admin', 'super_admin')			  NOT NULL,
    `created_at`    Timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`    Timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
)
    CHARACTER SET = utf8
    COLLATE = utf8_general_ci
    ENGINE = InnoDB;

-- +goose Down
DROP TABLE IF EXISTS `user`;
