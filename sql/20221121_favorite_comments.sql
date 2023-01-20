CREATE TABLE `favorite_comments` (
    `id` int UNSIGNED NOT NULL,
    `user_id` int UNSIGNED NOT NULL,
    `thread_id` int UNSIGNED NOT NULL,
    `created_at` int UNSIGNED NOT NULL,
    `updated_at` int UNSIGNED NOT NULL,
    `deleted_at` int DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


ALTER TABLE `favorite_comments`
    ADD PRIMARY KEY (`id`);


ALTER TABLE `favorite_comments`
    MODIFY `id` int UNSIGNED NOT NULL AUTO_INCREMENT;
