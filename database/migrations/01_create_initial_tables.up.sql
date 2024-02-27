CREATE TABLE `cat` (
                       `id` integer AUTO_INCREMENT PRIMARY KEY,
                       `no` integer,
                       `name` VARCHAR(250),
                       `image_path` VARCHAR(255),
                       `module_id` VARCHAR(250),
                       `home_display` bool,
                       `highlight` bool,
                       `menu_display` bool,
                       `display` bool,
                       `cat_path` VARCHAR(250),
                       `different_path` VARCHAR(250),
                       `description` VARCHAR(250),
                       `content` VARCHAR(250),
                       `banner_path` VARCHAR(250),
                       `gift` VARCHAR(250),
                       `is_follow` bool,
                       `is_index` bool,
                       `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                       `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE `module` (
                          `id` integer AUTO_INCREMENT PRIMARY KEY,
                          `name` VARCHAR(250),
                          `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                          `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE `content` (
                           `id` integer AUTO_INCREMENT PRIMARY KEY,
                           `no` int,
                           `title` VARCHAR(250),
                           `image` VARCHAR(250),
                           `display` bool,
                           `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                           `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE `company_info` (
                                `id` integer AUTO_INCREMENT PRIMARY KEY,
                                `stite_address` VARCHAR(250),
                                `copyright` VARCHAR(250),
                                `name` VARCHAR(250),
                                `address` VARCHAR(250),
                                `hotline` VARCHAR(250),
                                `phone` VARCHAR(250),
                                `email` VARCHAR(250),
                                `map` VARCHAR(250),
                                `zalo` VARCHAR(250),
                                `message` VARCHAR(250),
                                `skype` VARCHAR(250),
                                `facebook` VARCHAR(250),
                                `twiter` VARCHAR(250),
                                `linkedin` VARCHAR(250),
                                `youtube` VARCHAR(250),
                                `priterest` VARCHAR(250),
                                `instagram` VARCHAR(250),
                                `telegram` VARCHAR(250),
                                `whatsapp` VARCHAR(250),
                                `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE `users` (
                         `id` integer AUTO_INCREMENT PRIMARY KEY,
                         `sku` integer,
                         `username` VARCHAR(250),
                         `password` VARCHAR(250),
                         `name` VARCHAR(250),
                         `phone` VARCHAR(250),
                         `email` VARCHAR(250),
                         `role` VARCHAR(250),
                         `balance` int,
                         `lock` bool,
                         `status` bool,
                         `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                         `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE `deposit_withdraw` (
                                    `id` integer AUTO_INCREMENT PRIMARY KEY,
                                    `code` integer,
                                    `user_id` int,
                                    `is_deposit` bool,
                                    `is_withdraw` bool,
                                    `date` datetime,
                                    `date_approved` datetime,
                                    `balance` int,
                                    `status` bool,
                                    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE `card_info` (
                             `id` integer AUTO_INCREMENT PRIMARY KEY,
                             `user_id` VARCHAR(250),
                             `country` VARCHAR(250),
                             `province` VARCHAR(250),
                             `city` VARCHAR(250),
                             `card_number` VARCHAR(250),
                             `card_name` int,
                             `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                             `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE `oders` (
                         `id` integer AUTO_INCREMENT PRIMARY KEY,
                         `user_id` integer,
                         `coin_code` VARCHAR(250),
--     '60, 120 180, 240, 400'
                         `time` int,
                         `balance` int,
                         `date` datetime,
                         `is_inprogress` bool,
                        --     thang thua
                         `status` bool ,
                         `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                         `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE `finance` (
                           `id` integer AUTO_INCREMENT PRIMARY KEY,
                           `user_id` integer,
                           `expired_date` datetime,
                           `commit_date` datetime,
                           `balance` int,
                           `package_finance_id` int,
                           `status` bool,
                           `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                           `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE `package_finance` (
                                   `id` integer AUTO_INCREMENT PRIMARY KEY,
                                   `user_id` integer,
                                   `expired_date` datetime,
                                   `commit_date` datetime,
                                   `balance` int,
                                   `package_id` int,
                                   `status` bool,
                                   `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                   `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE `package` (
                           `id` integer AUTO_INCREMENT PRIMARY KEY,
                           `name` VARCHAR(250),
                           `balance` int,
                           `percent` DECIMAL,
                           `date_lock` int,
                           `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                           `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

