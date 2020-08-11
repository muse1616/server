-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE
IF EXISTS `user`;

CREATE TABLE `user` (
	`id` INT (10) UNSIGNED NOT NULL AUTO_INCREMENT,
	`created_at` datetime DEFAULT NULL,
	`updated_at` datetime DEFAULT NULL,
	`deleted_at` datetime DEFAULT NULL,
	`email` VARCHAR (255) NOT NULL,
	`password` VARCHAR (255) NOT NULL,
	`nickname` VARCHAR (255) DEFAULT NULL,
	`random_id` SMALLINT DEFAULT 0,
	PRIMARY KEY (`id`),
	UNIQUE (`email`),
	KEY `idx_student_deleted_at` (`deleted_at`)
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4;
ALTER TABLE `user`ADD INDEX email_index(`email`)