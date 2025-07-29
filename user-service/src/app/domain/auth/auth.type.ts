import { UserRoleType } from 'src/app/domain/user/user.type';

export interface UserBaseType {
  user_id: bigint;
  user_account: string;
  user_salt: string;
  user_roles: UserRoleType[];
}
