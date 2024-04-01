CREATE TABLE `record` (
                          `id` bigint NOT NULL AUTO_INCREMENT,
                          `user_id` varchar(255) DEFAULT NULL,
                          `record` varchar(255) DEFAULT NULL,
                          PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;