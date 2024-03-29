CREATE TABLE `follows` (
    `id` int UNSIGNED NOT NULL,
    `user_id` int UNSIGNED NOT NULL,
    `to_user_id` int UNSIGNED NOT NULL,
    `created_at` int UNSIGNED NOT NULL,
    `updated_at` int UNSIGNED NOT NULL,
    `deleted_at` int DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


ALTER TABLE `follows`
    ADD PRIMARY KEY (`id`);


ALTER TABLE `follows`
    MODIFY `id` int UNSIGNED NOT NULL AUTO_INCREMENT;
