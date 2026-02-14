-- You will add in this file your changes for the database
ALTER TABLE vehicle
    ADD COLUMN `state` TINYINT UNSIGNED NOT NULL DEFAULT 1 CHECK (`state` IN (0,1)) AFTER vehicle_year,
    ADD COLUMN `updated` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP AFTER `state`;

