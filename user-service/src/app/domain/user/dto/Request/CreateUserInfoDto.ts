import { Prisma } from '@prisma/client';

export interface CreateUserInfoDto {
  user_id: bigint;
  user_account: string;
  user_nickname: string | null;
  user_avatar?: string;
  user_state?: number;
  user_gender?: number;
  user_mobile?: string | null;
  user_email?: string | null;
  user_is_authentication: boolean;
}

export const createPrismaUserInfoDto = (createUserInfo: CreateUserInfoDto) => {
  return Prisma.validator<Prisma.UserBaseWhereInput>()({
    ...createUserInfo
  });
};
