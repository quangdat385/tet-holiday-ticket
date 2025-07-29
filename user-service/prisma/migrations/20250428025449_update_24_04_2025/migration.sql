/*
  Warnings:

  - Made the column `user_account` on table `pre_nestjs_acc_user_info_9999` required. This step will fail if there are existing NULL values in that column.

*/
-- AlterTable
ALTER TABLE `pre_nestjs_acc_user_info_9999` MODIFY `user_account` VARCHAR(255) NOT NULL;

-- CreateTable
CREATE TABLE `UserKeyToken` (
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `refresh_token` VARCHAR(600) NOT NULL,
    `user_id` BIGINT NOT NULL,
    `public_key` VARCHAR(600) NOT NULL,
    `private_key` VARCHAR(600) NOT NULL,
    `expire_at` DATETIME(3) NOT NULL,

    UNIQUE INDEX `UserKeyToken_refresh_token_key`(`refresh_token`),
    PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- AddForeignKey
ALTER TABLE `UserKeyToken` ADD CONSTRAINT `UserKeyToken_user_id_fkey` FOREIGN KEY (`user_id`) REFERENCES `pre_nestjs_acc_user_base_9999`(`user_id`) ON DELETE CASCADE ON UPDATE CASCADE;
