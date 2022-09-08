CREATE DATABASE IF NOT EXISTS `weight_data`;

USE `weight_data`;


CREATE TABLE IF NOT EXISTS `main` (
  `date` date NOT NULL,
  `weight` decimal(5,1) NOT NULL,
  PRIMARY KEY (`date`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE IF NOT EXISTS `users` (
  `id` VARCHAR(128) NOT NULL,
  `username` VARCHAR(256) NOT NULL,
  `password` VARCHAR(256) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE IF NOT EXISTS `test` (
  `date` date NOT NULL,
  `weight` decimal(5,1) NOT NULL,
  PRIMARY KEY (`date`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


CREATE USER 'archimedes'@'%' IDENTIFIED BY '@r(h1m3d3s';

GRANT ALL PRIVILEGES ON weight_data . * TO 'archimedes'@'%';

FLUSH PRIVILEGES;