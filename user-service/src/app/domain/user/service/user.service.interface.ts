import { SendEmailCommandOutput } from '@aws-sdk/client-ses';
import { UserBase, UserInfo, UserKeyToken, UserMenu, UserUserRole, UserVerify } from '@prisma/client';
import { UserBaseType } from 'src/app/domain/auth/auth.type';
import { CreateUserBaseDto } from 'src/app/domain/user/dto/Request/CreateUserBaseDto';
import { CreateUserInfoDto } from 'src/app/domain/user/dto/Request/CreateUserInfoDto';
import { CreateUserKeyToken } from 'src/app/domain/user/dto/Request/CreateUserKeyToken';
import { CreateUserUserRole } from 'src/app/domain/user/dto/Request/CreateUserUserRole';
import { CreateUserVerifyDto } from 'src/app/domain/user/dto/Request/CreateUserVerifyDto';
import { UpdateUserbaseDto } from 'src/app/domain/user/dto/Request/UpdateUserbaseDto';
import { UserInfoResponse } from 'src/app/domain/user/dto/response/GetUserInfo';

export type UserRoleNumber = 1 | 2 | 3 | 4;
export interface UserServiceInterface {
  //user base
  checkUserBaseExist(user_account: string): Promise<boolean>;
  getUserBaseByAccount(user_account: string): Promise<UserBase | null>;
  createUserBase(userBase: CreateUserBaseDto): Promise<UserBase>;
  updateUserLogin(userBase: UpdateUserbaseDto): Promise<UserBase | null>;
  //user user_role
  createUserUserRole(userUserRole: CreateUserUserRole): Promise<UserUserRole>;
  grantUserRole(user_id: bigint, role: UserRoleNumber): Promise<string>;
  // user verify
  getUserVerifyByKey(key: string): Promise<UserVerify | null>;
  getUserVerifyByKeyHash(key_hash: string): Promise<UserVerify | null>;
  createUserVerify(userVerify: CreateUserVerifyDto): Promise<UserVerify>;
  updateUserVerify(verify_id: number, userVerify: Partial<UserVerify>): Promise<UserVerify | null>;
  // user info
  createUserInfo(userInfo: CreateUserInfoDto): Promise<UserInfo>;
  getUserInfoByUserId(user_id: bigint): Promise<UserInfoResponse | null>;
  // user role
  findUserRolesByUserId(user_id: bigint): Promise<UserBaseType | null>;
  // key token
  createUserKeyToken(userKeyToken: CreateUserKeyToken): Promise<UserKeyToken>;
  updateUserKeyToken(token: string, expireAt: Date): Promise<UserKeyToken | null>;
  getUserKeyTokenById(id: bigint): Promise<UserKeyToken | null>;
  // user menu
  createUserMenu(user_id: bigint, menu_id: number): Promise<UserMenu>;
  // user role menu
  // otp
  getOtp(key: string): Promise<string | null>;
  setOtp(key: string, otp: number, ttl?: number): Promise<string>;
  delOtp(key: string): Promise<void>;
  // send email
  sendTextEmailOtp(
    subject: string,
    body: string,
    toAddress: string,
    fromAddress?: string
  ): Promise<SendEmailCommandOutput>;
}
