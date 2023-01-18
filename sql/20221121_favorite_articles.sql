CREATE TABLE `favorite_articles` (
    `id` int UNSIGNED NOT NULL,
    `user_id` int UNSIGNED NOT NULL,
    `article_id` int UNSIGNED NOT NULL,
    `created_at` int UNSIGNED NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


ALTER TABLE `favorite_articles`
    ADD PRIMARY KEY (`id`);


ALTER TABLE `favorite_articles`
    MODIFY `id` int UNSIGNED NOT NULL AUTO_INCREMENT;
