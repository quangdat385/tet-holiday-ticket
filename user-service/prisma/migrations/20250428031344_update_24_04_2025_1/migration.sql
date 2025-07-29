/*
  Warnings:

  - You are about to drop the `UserKeyToken` table. If the table is not empty, all the data it contains will be lost.

*/
-- DropForeignKey
ALTER TABLE `UserKeyToken` DROP FOREIGN KEY `UserKeyToken_user_id_fkey`;

-- DropTable
DROP TABLE `UserKeyToken`;

-- CreateTable
CREATE TABLE `pre_nestjs_acc_user_key_token_9999` (
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `refresh_token` VARCHAR(600) NOT NULL,
    `user_id` BIGINT NOT NULL,
    `public_key` VARCHAR(600) NOT NULL,
    `private_key` VARCHAR(600) NOT NULL,
    `expire_at` DATETIME(3) NOT NULL,

    UNIQUE INDEX `pre_nestjs_acc_user_key_token_9999_refresh_token_key`(`refresh_token`),
    INDEX `idx_user_id`(`user_id`),
    PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- AddForeignKey
ALTER TABLE `pre_nestjs_acc_user_key_token_9999` ADD CONSTRAINT `pre_nestjs_acc_user_key_token_9999_user_id_fkey` FOREIGN KEY (`user_id`) REFERENCES `pre_nestjs_acc_user_base_9999`(`user_id`) ON DELETE CASCADE ON UPDATE CASCADE;
