CREATE DATABASE IF NOT EXISTS php;
CREATE DATABASE IF NOT EXISTS golang;

# CREATE USER 'my'@'%' IDENTIFIED BY 'my';
GRANT CREATE, ALTER, INDEX, LOCK TABLES, REFERENCES, UPDATE, DELETE, DROP, SELECT, INSERT ON `php`.* TO 'my'@'%';
GRANT CREATE, ALTER, INDEX, LOCK TABLES, REFERENCES, UPDATE, DELETE, DROP, SELECT, INSERT ON `golang`.* TO 'my'@'%';

FLUSH PRIVILEGES;

#php

CREATE
    TABLE `php`.`10e1-row`
(
    `id`                       INT                                   NOT NULL AUTO_INCREMENT,
    `uid`                      VARCHAR(255)                          NULL,
    `manufacturer_part_number` VARCHAR(255)                          NULL,
    `manufacturer`             VARCHAR(255)                          NULL,
    `quantity`                 VARCHAR(255)                          NULL,
    `created_at`               TIMESTAMP on update CURRENT_TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;

CREATE
    TABLE `php`.`10e2-row`
(
    `id`                       INT                                   NOT NULL AUTO_INCREMENT,
    `uid`                      VARCHAR(255)                          NULL,
    `manufacturer_part_number` VARCHAR(255)                          NULL,
    `manufacturer`             VARCHAR(255)                          NULL,
    `quantity`                 VARCHAR(255)                          NULL,
    `created_at`               TIMESTAMP on update CURRENT_TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;

CREATE
    TABLE `php`.`10e3-row`
(
    `id`                       INT                                   NOT NULL AUTO_INCREMENT,
    `uid`                      VARCHAR(255)                          NULL,
    `manufacturer_part_number` VARCHAR(255)                          NULL,
    `manufacturer`             VARCHAR(255)                          NULL,
    `quantity`                 VARCHAR(255)                          NULL,
    `created_at`               TIMESTAMP on update CURRENT_TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;

CREATE
    TABLE `php`.`10e4-row`
(
    `id`                       INT                                   NOT NULL AUTO_INCREMENT,
    `uid`                      VARCHAR(255)                          NULL,
    `manufacturer_part_number` VARCHAR(255)                          NULL,
    `manufacturer`             VARCHAR(255)                          NULL,
    `quantity`                 VARCHAR(255)                          NULL,
    `created_at`               TIMESTAMP on update CURRENT_TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;

CREATE
    TABLE `php`.`10e5-row`
(
    `id`                       INT                                   NOT NULL AUTO_INCREMENT,
    `uid`                      VARCHAR(255)                          NULL,
    `manufacturer_part_number` VARCHAR(255)                          NULL,
    `manufacturer`             VARCHAR(255)                          NULL,
    `quantity`                 VARCHAR(255)                          NULL,
    `created_at`               TIMESTAMP on update CURRENT_TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;

CREATE
    TABLE `php`.`10e6-row`
(
    `id`                       INT                                   NOT NULL AUTO_INCREMENT,
    `uid`                      VARCHAR(255)                          NULL,
    `manufacturer_part_number` VARCHAR(255)                          NULL,
    `manufacturer`             VARCHAR(255)                          NULL,
    `quantity`                 VARCHAR(255)                          NULL,
    `created_at`               TIMESTAMP on update CURRENT_TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;

#golang

CREATE
    TABLE `golang`.`10e1-row`
(
    `id`                       INT                                   NOT NULL AUTO_INCREMENT,
    `uid`                      VARCHAR(255)                          NULL,
    `manufacturer_part_number` VARCHAR(255)                          NULL,
    `manufacturer`             VARCHAR(255)                          NULL,
    `quantity`                 VARCHAR(255)                          NULL,
    `created_at`               TIMESTAMP on update CURRENT_TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;

CREATE
    TABLE `golang`.`10e2-row`
(
    `id`                       INT                                   NOT NULL AUTO_INCREMENT,
    `uid`                      VARCHAR(255)                          NULL,
    `manufacturer_part_number` VARCHAR(255)                          NULL,
    `manufacturer`             VARCHAR(255)                          NULL,
    `quantity`                 VARCHAR(255)                          NULL,
    `created_at`               TIMESTAMP on update CURRENT_TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;

CREATE
    TABLE `golang`.`10e3-row`
(
    `id`                       INT                                   NOT NULL AUTO_INCREMENT,
    `uid`                      VARCHAR(255)                          NULL,
    `manufacturer_part_number` VARCHAR(255)                          NULL,
    `manufacturer`             VARCHAR(255)                          NULL,
    `quantity`                 VARCHAR(255)                          NULL,
    `created_at`               TIMESTAMP on update CURRENT_TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;

CREATE
    TABLE `golang`.`10e4-row`
(
    `id`                       INT                                   NOT NULL AUTO_INCREMENT,
    `uid`                      VARCHAR(255)                          NULL,
    `manufacturer_part_number` VARCHAR(255)                          NULL,
    `manufacturer`             VARCHAR(255)                          NULL,
    `quantity`                 VARCHAR(255)                          NULL,
    `created_at`               TIMESTAMP on update CURRENT_TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;

CREATE
    TABLE `golang`.`10e5-row`
(
    `id`                       INT                                   NOT NULL AUTO_INCREMENT,
    `uid`                      VARCHAR(255)                          NULL,
    `manufacturer_part_number` VARCHAR(255)                          NULL,
    `manufacturer`             VARCHAR(255)                          NULL,
    `quantity`                 VARCHAR(255)                          NULL,
    `created_at`               TIMESTAMP on update CURRENT_TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;

CREATE
    TABLE `golang`.`10e6-row`
(
    `id`                       INT                                   NOT NULL AUTO_INCREMENT,
    `uid`                      VARCHAR(255)                          NULL,
    `manufacturer_part_number` VARCHAR(255)                          NULL,
    `manufacturer`             VARCHAR(255)                          NULL,
    `quantity`                 VARCHAR(255)                          NULL,
    `created_at`               TIMESTAMP on update CURRENT_TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;