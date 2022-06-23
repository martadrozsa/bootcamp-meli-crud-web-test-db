CREATE DATABASE IF NOT EXISTS `product`;

USE `product`;

DROP TABLE IF EXISTS `product`;

use product;

CREATE TABLE `product` (
    `id` MEDIUMINT NOT NULL AUTO_INCREMENT,
    `name` varchar(100) NOT NULL,
    `product_type` varchar(100) NOT NULL,
    `description` varchar(255) NOT NULL,
    `quantity` int(10) unsigned NOT NULL,
    `price` float(10,2) unsigned NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

drop table product;