CREATE TABLE `dogs`
(
    `id`         bigint(20) unsigned                     NOT NULL AUTO_INCREMENT,
    `created_at` timestamp                               NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp                               NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `name`       VARCHAR(255) COLLATE utf8mb4_unicode_ci NOT NULL,
    `good_boy`   BOOL                                    NOT NULL,

    PRIMARY KEY (`id`),
    UNIQUE KEY `name_idx` (`name`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci
    COMMENT ='Dogs';

INSERT INTO `dogs`
    (`name`, `good_boy`)
VALUES ('Buddy', true),
       ('Luna', false),
       ('Rocky', true);
