import { Prisma } from '@prisma/client';

export interface CreateUserVerifyDto {
  verify_otp: string;
  verify_key: string;
  verify_type: number;
  verify_key_hash: string;
}

export const createPrismaUserVerifyDto = (createUserVerify: CreateUserVerifyDto) => {
  return Prisma.validator<Prisma.UserVerifyWhereInput>()({
    ...createUserVerify
  });
};
