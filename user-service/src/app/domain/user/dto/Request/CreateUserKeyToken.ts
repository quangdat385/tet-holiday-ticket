import { Prisma } from '@prisma/client';

export interface CreateUserKeyToken {
  user_id: bigint;
  refresh_token: string;
  private_key: string;
  public_key: string;
  expireAt: Date;
}

export const createPrismaUserKeyToken = (createUserKeyToken: CreateUserKeyToken) => {
  return Prisma.validator<Prisma.UserKeyTokenWhereInput>()({
    ...createUserKeyToken
  });
};
