import { SendEmailCommandOutput } from '@aws-sdk/client-ses';
import { BadRequestException, forwardRef, Inject, Injectable, NotAcceptableException } from '@nestjs/common';
import { ConfigService } from '@nestjs/config';
import { UserBase, UserInfo, UserKeyToken, UserMenu, UserUserRole, UserVerify } from '@prisma/client';
import { omit } from 'lodash';
import { UserBaseType } from 'src/app/domain/auth/auth.type';
import { AuthService } from 'src/app/domain/auth/service/impl/auth.service';
import { DistributeCacheService } from 'src/app/domain/cache/service/impl/cache.distribute';
import { LocalCacheService } from 'src/app/domain/cache/service/impl/cache.local';
import { DatabaseService } from 'src/app/domain/database/database.service';
import { ContactUs } from 'src/app/domain/user/dto/Request/ContactUs';
import { createPrismaUserBaseDto, CreateUserBaseDto } from 'src/app/domain/user/dto/Request/CreateUserBaseDto';
import { createPrismaUserInfoDto, CreateUserInfoDto } from 'src/app/domain/user/dto/Request/CreateUserInfoDto';
import { createPrismaUserKeyToken, CreateUserKeyToken } from 'src/app/domain/user/dto/Request/CreateUserKeyToken';
import { createPrismaUserUserRole, CreateUserUserRole } from 'src/app/domain/user/dto/Request/CreateUserUserRole';
import { createPrismaUserVerifyDto, CreateUserVerifyDto } from 'src/app/domain/user/dto/Request/CreateUserVerifyDto';
import { UpdatePasswordRequest } from 'src/app/domain/user/dto/Request/UpdatePassword';
import { UpdateUserbaseDto } from 'src/app/domain/user/dto/Request/UpdateUserbaseDto';
import { UpdateUserInfoRequest } from 'src/app/domain/user/dto/Request/UpdateUserInfo';
import { UserInfoResponse } from 'src/app/domain/user/dto/response/GetUserInfo';
import { UserRoleNumber, UserServiceInterface } from 'src/app/domain/user/service/user.service.interface';
import { UserRoleType } from 'src/app/domain/user/user.type';
import { SendMailService } from 'src/app/domain/utils/service/send-email.service';
import { hashPassword } from 'src/utils/bycript';

@Injectable()
export class UserService implements UserServiceInterface {
  constructor(
    private localCache: LocalCacheService,
    private distributeCache: DistributeCacheService,
    private db: DatabaseService,
    @Inject(forwardRef(() => AuthService))
    private authService: AuthService,
    private emailService: SendMailService,
    private configService: ConfigService
    // eslint-disable-next-line prettier/prettier
  ) { }
  createUserMenu(user_id: bigint, menu_id: number): Promise<UserMenu> {
    console.log('createUserMenu', user_id, menu_id);
    throw new Error('Method not implemented.');
  }
  async updateUserLogin(userBase: UpdateUserbaseDto): Promise<UserBase | null> {
    return await this.db.userBase.update({
      where: {
        user_id: userBase.user_id
      },
      data: {
        user_login_ip: userBase.user_login_ip,
        user_login_time: userBase.user_login_time
      }
    });
  }
  async updatePassswordByUserAccount(user_account: string, password: string): Promise<UserBase | null> {
    return await this.db.userBase.update({
      where: {
        user_account: user_account
      },
      data: {
        user_password: password
      }
    });
  }
  async updateUserPassword(user_id: bigint, body: UpdatePasswordRequest): Promise<boolean> {
    if (body.password !== body.confirm_password) {
      throw new NotAcceptableException('password and confirm password not match');
    }
    const passwordHash: string = await hashPassword(body.password);
    const userBase = await this.db.userBase.update({
      where: {
        user_id: user_id
      },
      data: {
        user_password: passwordHash
      }
    });
    if (!userBase) {
      return false;
    }
    return true;
  }
  async getUserBaseByAccount(user_account: string): Promise<UserBase | null> {
    return await this.db.userBase.findUnique({
      where: {
        user_account: user_account
      }
    });
  }

  async checkUserBaseExist(user_account: string): Promise<boolean> {
    const userBase = await this.db.userBase.findUnique({
      where: {
        user_account: user_account
      }
    });
    if (userBase) {
      return true;
    }
    return false;
  }
  async getUserVerifyByKey(verify_key: string): Promise<UserVerify | null> {
    const userVerify = await this.db.userVerify.findUnique({
      where: {
        verify_key
      }
    });
    if (userVerify) {
      return userVerify;
    }
    return null;
  }
  async getUserVerifyByKeyHash(key_hash: string): Promise<UserVerify | null> {
    const userVerify = await this.db.userVerify.findFirst({
      where: {
        verify_key_hash: key_hash
      }
    });
    if (userVerify) {
      return userVerify;
    }
    return null;
  }
  async getUserInfoByUserId(user_id: bigint): Promise<UserInfoResponse | null> {
    const userInfo = await this.db.userInfo.findUnique({
      where: {
        user_id: user_id
      },
      select: {
        user_id: true,
        user_account: true,
        user_email: true,
        user_mobile: true,
        user_nickname: true,
        user_avatar: true,
        user_birthday: true,
        user_gender: true,
        user_is_authentication: true,
        user_state: true
      }
    });
    if (!userInfo) {
      return null;
    }
    return {
      ...userInfo,
      user_id: Number(userInfo.user_id)
    };
  }
  async updateUserVerify(verify_id: number, userVerify: Partial<UserVerify>): Promise<UserVerify | null> {
    const userVerifyFound = await this.db.userVerify.findUnique({
      where: {
        verify_id
      }
    });
    if (!userVerifyFound) {
      return null;
    }
    return this.db.userVerify.update({
      where: {
        verify_id
      },
      data: userVerify
    });
  }
  async getUserKeyTokenById(id: bigint): Promise<UserKeyToken | null> {
    return await this.db.userKeyToken.findUnique({
      where: {
        id
      }
    });
  }
  async updateUserKeyToken(token: string, expireAt: Date): Promise<UserKeyToken | null> {
    const userKeyToken = await this.db.userKeyToken.update({
      where: {
        refresh_token: token
      },
      data: {
        expireAt: expireAt
      }
    });
    if (!userKeyToken) {
      return null;
    }
    return userKeyToken;
  }
  async grantUserRole(user_id: bigint, role: UserRoleNumber): Promise<string> {
    if (role === 1) {
      throw new NotAcceptableException('user role is not allowed to be set to 1');
    }
    await this.db.userUserRole.upsert({
      where: {
        user_id_role_id: {
          user_id: user_id,
          role_id: role
        }
      },
      create: {
        user_id: user_id,
        role_id: role
      },
      update: {
        user_id: user_id,
        role_id: role
      }
    });
    return 'user role granted successfully';
  }
  async getOtp(key: string): Promise<string | null> {
    const otp = await this.distributeCache.getString(key);
    if (otp) {
      return otp;
    }
    return null;
  }
  async setOtp(key: string, otp: number, ttl: number): Promise<string> {
    return this.distributeCache.setStringWithTtl(key, otp.toString(), ttl);
  }
  async delOtp(key: string): Promise<void> {
    return await this.distributeCache.del(key);
  }
  async sendTextEmailOtp(
    subject: string,
    body: string,
    toAddress: string,
    fromAddress?: string
  ): Promise<SendEmailCommandOutput> {
    const from = fromAddress ? fromAddress : this.configService.get<string>('from_address');
    console.log('sendTextEmailOtp', subject, body, toAddress, from);
    return this.emailService.sendVerifyEmail(from as string, toAddress, subject, body);
  }
  async contactUs(body: ContactUs): Promise<string> {
    const { firstName, lastName, email, phone, message } = body;
    const subject = `Contact Us from ${firstName} ${lastName}`;
    const bodyContent = `
      <p>First Name: ${firstName}</p>
      <p>Last Name: ${lastName}</p>
      <p>Email: ${email}</p>
      <p>Phone: ${phone}</p>
      <p>Message: ${message}</p>
    `;
    const fromAddress = this.configService.get<string>('from_address');
    await this.emailService.sendVerifyEmailHtml(fromAddress as string, 'cddd4f@yahoo.com', subject, bodyContent);
    return 'Contact us message sent successfully';
  }
  // Create user
  async createUserVerify(userVerify: CreateUserVerifyDto): Promise<UserVerify> {
    return await this.db.userVerify.create({
      data: createPrismaUserVerifyDto({ ...userVerify })
    });
  }
  async createUserBase(userBase: CreateUserBaseDto): Promise<UserBase> {
    return await this.db.userBase.create({
      data: createPrismaUserBaseDto({
        ...userBase
      })
    });
  }
  async createUserInfo(userInfo: CreateUserInfoDto): Promise<UserInfo> {
    return await this.db.userInfo.create({
      data: createPrismaUserInfoDto({
        ...userInfo
      })
    });
  }
  async createUserKeyToken(userKeyToken: CreateUserKeyToken): Promise<UserKeyToken> {
    return this.db.userKeyToken.create({
      data: createPrismaUserKeyToken({ ...userKeyToken })
    });
  }
  async createUserUserRole(userUserRole: CreateUserUserRole): Promise<UserUserRole> {
    return await this.db.userUserRole.create({
      data: createPrismaUserUserRole({
        ...userUserRole
      })
    });
  }
  async findUserRolesByUserId(user_id: bigint): Promise<UserBaseType | null> {
    const user = await this.db.userBase.findUnique({
      where: {
        user_id: user_id
      },
      include: {
        user_roles: {
          include: {
            userRole: {
              include: {
                role_menus: {
                  include: {
                    user_menu: true
                  }
                }
              }
            }
          }
        }
      }
    });
    if (!user) {
      return null;
    }
    const userRoles = user.user_roles.map((userRole) => {
      return {
        role_id: userRole.role_id,
        role_name: userRole.userRole.role_name,
        role_description: userRole.userRole.role_description,
        role_menus: userRole.userRole.role_menus.map((roleMenu) => {
          return {
            menu_id: roleMenu.user_menu.menu_id,
            menu_name: roleMenu.user_menu.menu_name,
            menu_url: roleMenu.user_menu.menu_url,
            menu_prefix: roleMenu.user_menu.menu_prefix,
            menu_pid: roleMenu.user_menu.menu_pid
          };
        })
      };
    });
    return {
      user_id: user.user_id,
      user_account: user.user_account,
      user_salt: user.user_salt,
      user_roles: userRoles as UserRoleType[]
    };
  }
  async deleteUserKeyTokenById(id: bigint): Promise<boolean> {
    const userKeyToken = await this.db.userKeyToken.findUnique({
      where: {
        id
      }
    });
    if (!userKeyToken) {
      return false;
    }
    const result = await this.db.userKeyToken.delete({
      where: {
        id
      }
    });
    if (!result) {
      return false;
    }
    return true;
  }
  async updateUserInfo(user_id: bigint, body: UpdateUserInfoRequest) {
    if (user_id !== body.user_id) {
      throw new BadRequestException('user_id not match');
    }
    const userInfo = await this.db.userInfo.findUnique({
      where: {
        user_id: body.user_id
      }
    });
    if (!userInfo) {
      throw new BadRequestException('user info not found');
    }
    const updateData = omit(body, ['user_id']);
    return await this.db.userInfo.update({
      where: {
        user_id: body.user_id
      },
      data: updateData
    });
  }
}
