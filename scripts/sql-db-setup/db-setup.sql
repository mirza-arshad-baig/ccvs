CREATE DATABASE credit_cards;

CREATE TABLE `credit_cards` (
	`id` INT(11) NOT NULL AUTO_INCREMENT,
	`card_number` VARCHAR(16) NOT NULL COLLATE 'utf8mb4_general_ci',
	`created_at` DATETIME NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`) USING BTREE,
	UNIQUE INDEX `card_number` (`card_number`) USING BTREE
)
COLLATE='utf8mb4_general_ci'
ENGINE=InnoDB
AUTO_INCREMENT=3
;
