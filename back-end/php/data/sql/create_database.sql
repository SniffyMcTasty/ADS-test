-- You can not modify the structure of the database here

CREATE TABLE IF NOT EXISTS vehicle_make
(
    vehicle_make_id int(11) unsigned                     NOT NULL AUTO_INCREMENT,
    `name`          varchar(255) COLLATE utf8_unicode_ci NOT NULL,
    `url`           varchar(255) COLLATE utf8_unicode_ci NOT NULL,
    PRIMARY KEY (vehicle_make_id),
    UNIQUE KEY `unq_vehicle_make_name` (`name`)
);

CREATE TABLE IF NOT EXISTS vehicle_model
(
    vehicle_model_id int(11) unsigned                     NOT NULL AUTO_INCREMENT,
    `name`           varchar(255) COLLATE utf8_unicode_ci NOT NULL,
    PRIMARY KEY (vehicle_model_id),
    UNIQUE KEY `unq_vehicle_model_name` (`name`)
);


CREATE TABLE IF NOT EXISTS vehicle
(
    vehicle_id       int(11) unsigned NOT NULL AUTO_INCREMENT,
    vehicle_make_id  int(11) unsigned NOT NULL,
    vehicle_model_id int(11) unsigned NOT NULL,
    vehicle_year     smallint(4) unsigned,
    PRIMARY KEY (vehicle_id),
    CONSTRAINT `fk_vehicle_make_id` FOREIGN KEY (vehicle_make_id) REFERENCES vehicle_make (vehicle_make_id),
    CONSTRAINT `fk_vehicle_model_id` FOREIGN KEY (vehicle_model_id) REFERENCES vehicle_model (vehicle_model_id)
);
