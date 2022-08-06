
CREATE TABLE `user_tokens` (
  `id` int UNSIGNED NOT NULL,
  `user_id` int UNSIGNED NOT NULL,
  `token` varchar(255) NOT NULL,
  `token_expire_at` int UNSIGNED NOT NULL,
  `refresh_token` varchar(255) NOT NULL,
  `refresh_token_expire_at` int UNSIGNED NOT NULL,
  `created_at` int UNSIGNED NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

ALTER TABLE `user_tokens`
    ADD PRIMARY KEY (`id`);

ALTER TABLE `user_tokens`
  MODIFY `id` int UNSIGNED NOT NULL AUTO_INCREMENT;