-- AlterTable
ALTER TABLE `pre_nestjs_acc_user_base_9999` MODIFY `user_created_at` TIMESTAMP(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
    MODIFY `user_updated_at` TIMESTAMP(0) NOT NULL;

-- AlterTable
ALTER TABLE `pre_nestjs_acc_user_info_9999` MODIFY `user_birthday` TIMESTAMP(0) NULL,
    MODIFY `created_at` TIMESTAMP(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
    MODIFY `updated_at` TIMESTAMP(0) NOT NULL;

-- AlterTable
ALTER TABLE `pre_nestjs_acc_user_key_token_9999` MODIFY `expire_at` TIMESTAMP(0) NOT NULL;

-- AlterTable
ALTER TABLE `pre_nestjs_acc_user_menu_9999` MODIFY `created_at` TIMESTAMP(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
    MODIFY `updated_at` TIMESTAMP(0) NOT NULL;

-- AlterTable
ALTER TABLE `pre_nestjs_acc_user_role_9999` MODIFY `created_at` TIMESTAMP(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
    MODIFY `updated_at` TIMESTAMP(0) NOT NULL;

-- AlterTable
ALTER TABLE `pre_nestjs_acc_user_role_menu_9999` MODIFY `created_at` TIMESTAMP(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
    MODIFY `updated_at` TIMESTAMP(0) NOT NULL;

-- AlterTable
ALTER TABLE `pre_nestjs_acc_user_two_factor_9999` MODIFY `two_factor_created_at` TIMESTAMP(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
    MODIFY `two_factor_updated_at` TIMESTAMP(0) NOT NULL;

-- AlterTable
ALTER TABLE `pre_nestjs_acc_user_verify_9999` MODIFY `verify_created_at` TIMESTAMP(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
    MODIFY `verify_updated_at` TIMESTAMP(0) NOT NULL;
