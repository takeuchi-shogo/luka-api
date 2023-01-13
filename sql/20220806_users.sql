
CREATE TABLE `users` (
    `id` int UNSIGNED NOT NULL,
    `screen_name` varchar(50) NOT NULL,
    `display_name` varchar(50) NOT NULL,
    `password` varchar(255) NOT NULL,
    `email` varchar(200) NOT NULL,
    `age` int NULL,
    `gender` int UNSIGNED NOT NULL,
    `prefecture` int UNSIGNED NOT NULL,
    `created_at` int UNSIGNED NOT NULL,
    `updated_at` int UNSIGNED NOT NULL,
    `deleted_at` int DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


ALTER TABLE `users`
    ADD PRIMARY KEY (`id`);


ALTER TABLE `users`
    MODIFY `id` int UNSIGNED NOT NULL AUTO_INCREMENT;
