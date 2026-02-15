-- You will add in this file your changes for the database
ALTER TABLE vehicle
    ADD COLUMN `state` TINYINT UNSIGNED NOT NULL DEFAULT 1 CHECK (`state` IN (0,1)) AFTER vehicle_year,
    ADD COLUMN `updated` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP AFTER `state`;
    ADD UNIQUE INDEX idx_vehicle_make_model_year (vehicle_make_id, vehicle_model_id, vehicle_year);

CREATE TABLE IF NOT EXISTS `user` (
    `user_id` BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `username` VARCHAR(255) NOT NULL UNIQUE,
    `password_hash` VARCHAR(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Insert defaut user admin / secret123
INSERT INTO `user` (`username`, `password_hash`) VALUES ('admin', '$2a$10$3hPtwZ7TLT8nViTCuneW5O1XPShr5R37HceAWXJmvxhtx9oAozYRa');