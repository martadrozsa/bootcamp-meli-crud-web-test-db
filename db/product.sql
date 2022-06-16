CREATE DATABASE IF NOT EXISTS `product`;

USE `product`;

DROP TABLE IF EXISTS `product`;

CREATE TABLE `product` (
    `id` varchar(36) NOT NULL,
    `name` varchar(100) NOT NULL,
    `product_type` varchar(100) NOT NULL,
    `description` varchar(255) NOT NULL,
    `quantity` int(10) unsigned NOT NULL,
    `price` float(10,2) unsigned NOT NULL,
  PRIMARY KEY (`uuid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;