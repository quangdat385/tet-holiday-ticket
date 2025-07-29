-- CreateTable
CREATE TABLE `pre_nestjs_acc_user_verify_9999` (
    `verify_id` INTEGER NOT NULL AUTO_INCREMENT,
    `verify_otp` VARCHAR(6) NOT NULL,
    `verify_key` VARCHAR(255) NOT NULL,
    `verify_key_hash` VARCHAR(255) NOT NULL,
    `verify_type` INTEGER NOT NULL DEFAULT 1,
    `is_verified` INTEGER NOT NULL DEFAULT 0,
    `is_deleted` INTEGER NOT NULL DEFAULT 0,
    `verify_created_at` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `verify_updated_at` DATETIME(3) NOT NULL,

    UNIQUE INDEX `pre_nestjs_acc_user_verify_9999_verify_key_key`(`verify_key`),
    INDEX `idx_verify_otp`(`verify_otp`),
    PRIMARY KEY (`verify_id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- CreateTable
CREATE TABLE `pre_nestjs_acc_user_base_9999` (
    `user_id` BIGINT NOT NULL AUTO_INCREMENT,
    `user_account` VARCHAR(255) NOT NULL,
    `user_password` VARCHAR(255) NOT NULL,
    `user_salt` VARCHAR(255) NOT NULL,
    `user_login_time` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `user_logout_time` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `user_login_ip` VARCHAR(255) NULL,
    `is_two_factor_enabled` INTEGER NOT NULL DEFAULT 0,
    `user_created_at` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `user_updated_at` DATETIME(3) NOT NULL,

    UNIQUE INDEX `pre_nestjs_acc_user_base_9999_user_account_key`(`user_account`),
    PRIMARY KEY (`user_id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- CreateTable
CREATE TABLE `pre_nestjs_acc_user_info_9999` (
    `user_id` BIGINT NOT NULL,
    `user_account` VARCHAR(255) NULL,
    `user_nickname` VARCHAR(255) NULL,
    `user_avatar` VARCHAR(255) NULL,
    `user_state` INTEGER NOT NULL DEFAULT 0,
    `user_gender` INTEGER NOT NULL DEFAULT 0,
    `user_mobile` VARCHAR(20) NULL,
    `user_birthday` DATETIME(3) NULL,
    `user_email` VARCHAR(255) NULL,
    `user_is_authentication` BOOLEAN NOT NULL DEFAULT false,
    `created_at` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `updated_at` DATETIME(3) NOT NULL,

    UNIQUE INDEX `pre_nestjs_acc_user_info_9999_user_account_key`(`user_account`),
    INDEX `idx_user_mobile`(`user_mobile`),
    INDEX `idx_user_email`(`user_email`),
    INDEX `idx_user_state`(`user_state`),
    INDEX `idx_user_is_authentication`(`user_is_authentication`),
    PRIMARY KEY (`user_id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- CreateTable
CREATE TABLE `pre_nestjs_acc_user_two_factor_9999` (
    `two_factor_id` BIGINT NOT NULL AUTO_INCREMENT,
    `user_id` BIGINT NOT NULL,
    `two_factor_auth_type` INTEGER NOT NULL DEFAULT 0,
    `two_factor_auth_secret` VARCHAR(255) NULL,
    `two_factor_phone` VARCHAR(20) NULL,
    `two_factor_email` VARCHAR(255) NULL,
    `two_factor_is_active` INTEGER NOT NULL DEFAULT 0,
    `two_factor_created_at` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `two_factor_updated_at` DATETIME(3) NOT NULL,

    INDEX `idx_user_id`(`user_id`),
    INDEX `idx_two_factor_auth_type`(`two_factor_auth_type`),
    PRIMARY KEY (`two_factor_id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- CreateTable
CREATE TABLE `pre_nestjs_acc_user_role_9999` (
    `role_id` BIGINT NOT NULL AUTO_INCREMENT,
    `role_name` VARCHAR(25) NOT NULL,
    `role_description` VARCHAR(255) NULL,
    `user_created` BIGINT NOT NULL,
    `created_at` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `updated_at` DATETIME(3) NOT NULL,

    UNIQUE INDEX `pre_nestjs_acc_user_role_9999_role_name_key`(`role_name`),
    PRIMARY KEY (`role_id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- CreateTable
CREATE TABLE `pre_nestjs_acc_user_user_role_9999` (
    `user_id` BIGINT NOT NULL,
    `role_id` BIGINT NOT NULL,

    PRIMARY KEY (`user_id`, `role_id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- CreateTable
CREATE TABLE `pre_nestjs_acc_user_menu_9999` (
    `menu_id` BIGINT NOT NULL AUTO_INCREMENT,
    `menu_name` VARCHAR(255) NOT NULL,
    `menu_pid` VARCHAR(255) NOT NULL,
    `menu_prefix` VARCHAR(255) NOT NULL,
    `menu_url` VARCHAR(255) NOT NULL,
    `user_crearted` BIGINT NOT NULL,
    `created_at` DATETIME(3) NOT NULL,
    `updated_at` DATETIME(3) NOT NULL,

    UNIQUE INDEX `pre_nestjs_acc_user_menu_9999_menu_name_key`(`menu_name`),
    UNIQUE INDEX `pre_nestjs_acc_user_menu_9999_menu_pid_key`(`menu_pid`),
    UNIQUE INDEX `pre_nestjs_acc_user_menu_9999_menu_prefix_key`(`menu_prefix`),
    PRIMARY KEY (`menu_id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- CreateTable
CREATE TABLE `pre_nestjs_acc_user_role_menu_9999` (
    `role_id` BIGINT NOT NULL,
    `menu_id` BIGINT NOT NULL,
    `created_at` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `updated_at` DATETIME(3) NOT NULL,

    PRIMARY KEY (`role_id`, `menu_id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- AddForeignKey
ALTER TABLE `pre_nestjs_acc_user_user_role_9999` ADD CONSTRAINT `pre_nestjs_acc_user_user_role_9999_user_id_fkey` FOREIGN KEY (`user_id`) REFERENCES `pre_nestjs_acc_user_base_9999`(`user_id`) ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE `pre_nestjs_acc_user_user_role_9999` ADD CONSTRAINT `pre_nestjs_acc_user_user_role_9999_role_id_fkey` FOREIGN KEY (`role_id`) REFERENCES `pre_nestjs_acc_user_role_9999`(`role_id`) ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE `pre_nestjs_acc_user_menu_9999` ADD CONSTRAINT `pre_nestjs_acc_user_menu_9999_user_crearted_fkey` FOREIGN KEY (`user_crearted`) REFERENCES `pre_nestjs_acc_user_base_9999`(`user_id`) ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE `pre_nestjs_acc_user_role_menu_9999` ADD CONSTRAINT `pre_nestjs_acc_user_role_menu_9999_role_id_fkey` FOREIGN KEY (`role_id`) REFERENCES `pre_nestjs_acc_user_role_9999`(`role_id`) ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE `pre_nestjs_acc_user_role_menu_9999` ADD CONSTRAINT `pre_nestjs_acc_user_role_menu_9999_menu_id_fkey` FOREIGN KEY (`menu_id`) REFERENCES `pre_nestjs_acc_user_menu_9999`(`menu_id`) ON DELETE CASCADE ON UPDATE CASCADE;
