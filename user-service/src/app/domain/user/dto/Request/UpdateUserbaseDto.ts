import { Prisma } from '@prisma/client';

export interface UpdateUserbaseDto {
  user_id: bigint;
  user_login_time?: Date;
  user_logout_time?: Date;
  user_login_ip: string;
}

export function UpdatePrismaUserBase(userBase: UpdateUserbaseDto): UpdateUserbaseDto {
  return Prisma.validator<Prisma.UserBaseWhereInput>()({
    ...userBase
  });
}
