import { Prisma } from '@prisma/client';

export interface CreateUserBaseDto {
  user_account: string;
  user_password: string;
  user_salt: string;
}

export const createPrismaUserBaseDto = (createUserBase: CreateUserBaseDto) => {
  return Prisma.validator<Prisma.UserBaseWhereInput>()({
    ...createUserBase
  });
};
