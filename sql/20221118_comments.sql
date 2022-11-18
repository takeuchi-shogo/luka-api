

CREATE TABLE `comments` (
    `id` int UNSIGNED NOT NULL,
    `user_id` varchar(50) NOT NULL,
    `thread_id` varchar(50) NOT NULL,
    `content` text NOT NULL,
    `created_at` int UNSIGNED NOT NULL,
    `updated_at` int UNSIGNED NOT NULL,
    `deleted_at` int DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


ALTER TABLE `comments`
    ADD PRIMARY KEY (`id`);


ALTER TABLE `commtents`
    MODIFY `id` int UNSIGNED NOT NULL AUTO_INCREMENT;
