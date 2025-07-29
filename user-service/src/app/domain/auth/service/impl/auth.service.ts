import { SendEmailCommandOutput } from '@aws-sdk/client-ses';
import { BadRequestException, forwardRef, Inject, Injectable } from '@nestjs/common';
import { JwtService } from '@nestjs/jwt';
import { JwtPayload } from 'jsonwebtoken';
import { PrismaClient, UserBase, UserInfo, UserKeyToken, UserVerify } from '@prisma/client';
import appConstant, { VerifyType } from 'src/app/app.constant';
import { CreateTokenPayLoad, Role } from 'src/app/domain/auth/dto/request/CreateToken';
import { LoginRequest } from 'src/app/domain/auth/dto/request/Login';
import { RegisterRequest } from 'src/app/domain/auth/dto/request/Register';
import { UpdatePasswordRegisterRequest } from 'src/app/domain/auth/dto/request/UpdatePasswordRegister';
import VerifyOtpRequest from 'src/app/domain/auth/dto/request/VerifyOtp';
import { LoginResponse } from 'src/app/domain/auth/dto/response/Login';
import { AuthServiceInterface } from 'src/app/domain/auth/service/auth.servie.interface';
import { UserService } from 'src/app/domain/user/service/impl/user.service';
import { hashPassword } from 'src/utils/bycript';
import { generateHash, getUserBaseKey, getUserHashKey } from 'src/utils/crypto';
import { generateOtp } from 'src/utils/otp';
import { compare } from 'bcrypt';
import { randomBytes } from 'crypto';
import { DistributeCacheService } from 'src/app/domain/cache/service/impl/cache.distribute';
import { UserBaseType } from 'src/app/domain/auth/auth.type';
import { Response } from 'express';
import { OAuthLoginResponse } from 'src/app/domain/auth/dto/request/FacebookAuth';
const prisma = new PrismaClient();

@Injectable()
export class AuthService implements AuthServiceInterface {
  private readonly salt = '123456';
  constructor(
    @Inject(forwardRef(() => UserService)) private userService: UserService,
    private readonly jwtService: JwtService,
    private distributeCache: DistributeCacheService
    // eslint-disable-next-line prettier/prettier
  ) { }
  getSalt(): string {
    return this.salt;
  }
  async register(registerRequest: RegisterRequest): Promise<string> {
    console.log('registerRequest', registerRequest);
    const hashKey = generateHash(registerRequest.verify_key, this.salt, 'sha256');
    const userFound = await this.userService.checkUserBaseExist(registerRequest.verify_key);
    console.log('userFound', userFound);
    if (userFound) {
      throw new BadRequestException(appConstant.USER_EXISTING);
    }
    const userKey = getUserBaseKey(hashKey);
    const otpFound = await this.userService.getOtp(userKey);
    console.log('otpFound', otpFound);
    if (otpFound) {
      await this.userService.delOtp(userKey); // remove otp from cache
      throw new BadRequestException(appConstant.USER_ALREADY_REGISTER);
    }
    let otp: number;
    if (registerRequest.verify_purpose === 'TEST_USER') {
      otp = 123456;
    }
    otp = generateOtp();
    const setOtp = await this.userService.setOtp(userKey, otp, 15 * 60); // 15 minutes
    if (!setOtp) {
      throw new BadRequestException(appConstant.USER_REGISTER_FAIL);
    }
    const toEmail: string = registerRequest.verify_key.trim();
    switch (registerRequest.verify_type) {
      case VerifyType.Email: {
        const sendEmail: SendEmailCommandOutput = await this.userService.sendTextEmailOtp(
          'Verify OTP Email',
          `Your OTP is ${otp} and valid for 15 minutes`,
          toEmail
        );
        if (!sendEmail) {
          await this.userService.delOtp(userKey); // remove otp from cache
          throw new BadRequestException(appConstant.SEND_VERIFY_EMAIL_FAILED);
        }
        const userVerify: UserVerify | null = await this.userService.getUserVerifyByKey(registerRequest.verify_key);
        if (userVerify) {
          await this.userService.updateUserVerify(userVerify.verify_id, {
            verify_otp: otp.toString()
          });
        } else {
          await this.userService.createUserVerify({
            verify_otp: otp.toString(),
            verify_key: registerRequest.verify_key,
            verify_type: registerRequest.verify_type,
            verify_key_hash: hashKey
          });
        }
        break;
      }
      case VerifyType.Phone: {
        await this.userService.delOtp(userKey); // remove otp from cache
        throw new BadRequestException(appConstant.PHONE_REGISTER_NOT_SUPPORT);
      }
      default: {
        throw new BadRequestException(appConstant.USER_REGISTER_FAIL);
      }
    }
    return 'User registered successfully, Please verify your email otp';
  }
  async verifyOtp(verifyOtpRequet: VerifyOtpRequest): Promise<{ token: string }> {
    const hashKey: string = generateHash(verifyOtpRequet.verify_key, this.salt, 'sha256');
    const userKey: string = getUserBaseKey(hashKey);
    const otpFound = await this.userService.getOtp(userKey);
    if (!otpFound) {
      throw new BadRequestException(appConstant.USER_OTP_NOT_FOUND);
    }
    if (otpFound !== verifyOtpRequet.verify_otp) {
      throw new BadRequestException(appConstant.USER_OTP_NOT_MATCH);
    }
    const userVerify: UserVerify | null = await this.userService.getUserVerifyByKey(verifyOtpRequet.verify_key);
    if (!userVerify) {
      throw new BadRequestException(appConstant.USER_NOT_FOUND);
    }
    if (userVerify.is_verified) {
      throw new BadRequestException(appConstant.USER_ALREADY_VERIFY);
    }
    await this.userService.delOtp(userKey); // remove otp from cache
    await this.userService.updateUserVerify(userVerify.verify_id, { is_verified: 1 });
    return {
      token: userVerify.verify_key_hash
    };
  }
  async updatePasswordRegister(updatePasswordRegister: UpdatePasswordRegisterRequest): Promise<string> {
    const { token, password, confirm_password } = updatePasswordRegister;
    if (password !== confirm_password) {
      throw new BadRequestException(appConstant.PASSWORD_INVALID);
    }
    const userVerify: UserVerify | null = await this.userService.getUserVerifyByKeyHash(token);
    if (!userVerify) {
      throw new BadRequestException(appConstant.USER_NOT_FOUND);
    }
    if (!userVerify.is_verified) {
      throw new BadRequestException(appConstant.USER_NOT_VERIFIED);
    }
    const passwordHash: string = await hashPassword(password);
    await prisma.$transaction(async () => {
      const userBase: UserBase = await this.userService.createUserBase({
        user_account: userVerify.verify_key,
        user_password: passwordHash,
        user_salt: this.salt
      });
      if (!userBase) {
        throw new BadRequestException(appConstant.UPDATE_PASSWORD_REGISTER_FAIL);
      }
      const userUserRole = await this.userService.createUserUserRole({
        user_id: userBase.user_id,
        role_id: 4n
      });
      if (!userUserRole) {
        throw new BadRequestException(appConstant.UPDATE_PASSWORD_REGISTER_FAIL);
      }
      const userInfo: UserInfo = await this.userService.createUserInfo({
        user_id: userBase.user_id,
        user_account: userBase.user_account,
        user_nickname: userBase.user_account,
        user_email: (userVerify.verify_type as VerifyType) === VerifyType.Email ? userBase.user_account : undefined,
        user_mobile: (userVerify.verify_type as VerifyType) === VerifyType.Phone ? userBase.user_account : undefined,
        user_state: 1,
        user_is_authentication: true
      });
      if (!userInfo) {
        throw new BadRequestException(appConstant.UPDATE_PASSWORD_REGISTER_FAIL);
      }
    });
    return 'Update password regiter successfully';
  }
  async reSendVerifyOtp(reSendVerifyOtp: Omit<RegisterRequest, 'verify_purpose'>): Promise<string> {
    const hashKey: string = generateHash(reSendVerifyOtp.verify_key, this.salt, 'sha256');
    const userFound = await this.userService.checkUserBaseExist(reSendVerifyOtp.verify_key);
    if (userFound) {
      throw new BadRequestException(appConstant.USER_EXISTING);
    }
    const userKey = getUserBaseKey(hashKey);
    const otp: number = generateOtp();
    const setOtp = await this.userService.setOtp(userKey, otp, 15 * 60); // 15 minutes
    if (!setOtp) {
      throw new BadRequestException(appConstant.USER_REGISTER_FAIL);
    }
    switch (reSendVerifyOtp.verify_type) {
      case VerifyType.Email: {
        const sendEmail: SendEmailCommandOutput = await this.userService.sendTextEmailOtp(
          'Verify OTP Email',
          `Your OTP is ${otp} and valid for 15 minutes`,
          reSendVerifyOtp.verify_key
        );
        if (!sendEmail) {
          await this.userService.delOtp(userKey); // remove otp from cache
          throw new BadRequestException(appConstant.SEND_VERIFY_EMAIL_FAILED);
        }
        const userVerify: UserVerify | null = await this.userService.getUserVerifyByKey(reSendVerifyOtp.verify_key);
        if (!userVerify) {
          await this.userService.createUserVerify({
            verify_otp: otp.toString(),
            verify_key: reSendVerifyOtp.verify_key,
            verify_type: reSendVerifyOtp.verify_type,
            verify_key_hash: hashKey
          });
        } else {
          await this.userService.updateUserVerify(userVerify.verify_id, {
            verify_otp: otp.toString()
          });
        }
        break;
      }
      case VerifyType.Phone: {
        await this.userService.delOtp(userKey); // remove otp from cache
        throw new BadRequestException(appConstant.PHONE_REGISTER_NOT_SUPPORT);
      }
      default: {
        throw new BadRequestException(appConstant.USER_REGISTER_FAIL);
      }
    }
    return 'Resend verify otp successfully';
  }
  async forgotPassword(
    forgotPassword: Omit<RegisterRequest, 'verify_purpose'>
  ): Promise<{ message: string; state: number }> {
    const hashKey: string = generateHash(forgotPassword.verify_key, this.salt, 'sha256');
    const userKey: string = getUserBaseKey(hashKey);
    const userVerify: UserVerify | null = await this.userService.getUserVerifyByKey(forgotPassword.verify_key);
    if (!userVerify) {
      await this.userService.delOtp(userKey); // remove otp from cache
      return {
        message: appConstant.USER_VERIFY_NOT_FOUND,
        state: 0
      };
    }
    if (!userVerify.is_verified) {
      return {
        message: appConstant.USER_NOT_VERIFIED,
        state: 1
      };
    }
    const checkUserBase: boolean = await this.userService.checkUserBaseExist(forgotPassword.verify_key);
    if (!checkUserBase) {
      await this.userService.delOtp(userKey); // remove otp from cache
      return {
        message: appConstant.USER_NOT_FOUND,
        state: 3
      };
    }
    if (forgotPassword.verify_type === VerifyType.Email) {
      const otp: number = generateOtp();
      const setOtp = await this.userService.setOtp(userKey, otp, 15 * 60); // 15 minutes
      if (!setOtp) {
        throw new BadRequestException(appConstant.SEND_FORGET_PASSWORD_EMAIL_FAILED);
      }
      const sendEmail: SendEmailCommandOutput = await this.userService.sendTextEmailOtp(
        'Verify OTP Email',
        `Your OTP is ${otp} and valid for 15 minutes`,
        forgotPassword.verify_key
      );
      if (!sendEmail) {
        await this.userService.delOtp(userKey); // remove otp from cache
        throw new BadRequestException(appConstant.SEND_FORGET_PASSWORD_EMAIL_FAILED);
      }
    }
    return {
      message: appConstant.SEND_FORGET_PASSWORD_EMAIL_SUCCESS,
      state: 4
    };
  }
  async verifyForgotPassword(verifyForgotPassword: VerifyOtpRequest): Promise<{ token: string }> {
    const hashKey: string = generateHash(verifyForgotPassword.verify_key, this.salt, 'sha256');
    const userKey: string = getUserBaseKey(hashKey);
    const otpFound = await this.userService.getOtp(userKey);
    if (!otpFound) {
      throw new BadRequestException(appConstant.USER_OTP_NOT_FOUND);
    }
    if (otpFound !== verifyForgotPassword.verify_otp) {
      throw new BadRequestException(appConstant.USER_OTP_NOT_MATCH);
    }
    await this.userService.delOtp(userKey); // remove otp from cache
    return {
      token: hashKey
    };
  }
  async updatePassword(updatePassword: UpdatePasswordRegisterRequest): Promise<string> {
    const { token, password, confirm_password } = updatePassword;
    if (password !== confirm_password) {
      throw new BadRequestException(appConstant.PASSWORD_INVALID);
    }
    const userVerify: UserVerify | null = await this.userService.getUserVerifyByKeyHash(token);
    if (!userVerify) {
      throw new BadRequestException(appConstant.USER_NOT_FOUND);
    }
    if (!userVerify.is_verified) {
      throw new BadRequestException(appConstant.USER_NOT_VERIFIED);
    }

    const passwordHash: string = await hashPassword(password);
    const userBaseUpdate: UserBase | null = await this.userService.updatePassswordByUserAccount(
      userVerify.verify_key,
      passwordHash
    );
    if (!userBaseUpdate) {
      throw new BadRequestException(appConstant.UPDATE_PASSWORD_FORGET_FAIL);
    }
    return appConstant.UPDATE_PASSWORD_SUCCESS;
  }
  async loginRespose(userBase: UserBaseType, ip: string, device: string): Promise<LoginResponse> {
    const privateKey = randomBytes(32).toString('hex');
    const publicKey = randomBytes(32).toString('hex');
    const now = new Date();

    const expireAt = new Date(now.getTime() + 24 * 30 * 60 * 60 * 1000);
    const roles = userBase?.user_roles.map((role) => {
      return {
        role_name: role.role_name,
        role_description: role.role_description
      };
    });
    const { user, access_token, refresh_token }: LoginResponse = await this.createToKen({
      user_account: userBase.user_account,
      user_roles: roles as Role[],
      public_key: publicKey,
      private_key: privateKey,
      public_expires: '15m',
      private_expires: '30d'
    });
    const keyToken = await this.userService.createUserKeyToken({
      user_id: userBase.user_id,
      refresh_token: refresh_token,
      public_key: publicKey,
      private_key: privateKey,
      expireAt: expireAt
    });
    if (!keyToken) {
      throw new BadRequestException(appConstant.USER_LOGIN_FAIL);
    }
    const hashKey = generateHash(userBase.user_id.toString(), this.salt, 'sha256');
    const userKey = getUserHashKey(hashKey, device);
    const userBaseKey = getUserBaseKey(hashKey);
    console.log('userKey', userKey);
    await this.distributeCache.setObjectWithTtl('user_key_token', userKey, keyToken, 60 * 60 * 24 * 30); // 30 days
    await this.distributeCache.setObjectWithTtl('user_base', userBaseKey, userBase, 30 * 60); // 30 minutes
    console.log('userBaseRoles', userBase);
    await this.userService.updateUserLogin({
      user_id: userBase.user_id,
      user_login_ip: ip,
      user_login_time: new Date()
    });
    return {
      user: {
        user_account: user.user_account as string,
        user_roles: roles as Role[],
        user_id: Number(userBase.user_id),
        user_salt: this.salt
      },
      access_token,
      refresh_token
    };
  }
  async login(loginRequest: LoginRequest, ip: string, device: string): Promise<LoginResponse> {
    const { user_account, password } = loginRequest;
    const userBase: UserBase | null = await this.userService.getUserBaseByAccount(user_account);
    if (!userBase) {
      throw new BadRequestException(appConstant.USER_NOT_FOUND);
    }
    const userBaseRoles = await this.userService.findUserRolesByUserId(userBase.user_id);
    if (!userBaseRoles) {
      throw new BadRequestException(appConstant.USER_NOT_FOUND);
    }
    const isMatch = await this.comparePassword(password, userBase.user_password);
    if (!isMatch) {
      throw new BadRequestException(appConstant.PASSWORD_INVALID);
    }
    return this.loginRespose(userBaseRoles, ip, device);
  }
  async createToKen(payload: CreateTokenPayLoad): Promise<LoginResponse> {
    const data: JwtPayload = {
      user_account: payload.user_account,
      user_roles: payload.user_roles
    };
    const [accessToken, refreshToken] = await Promise.all([
      this.jwtService.signAsync(data, {
        secret: payload.public_key,
        expiresIn: payload.public_expires || '15m'
      }),
      this.jwtService.signAsync(data, {
        secret: payload.private_key,
        expiresIn: payload.private_expires || '30d'
      })
    ]);
    return {
      user: {
        user_account: payload.user_account,
        user_roles: payload.user_roles
      },
      access_token: accessToken,
      refresh_token: refreshToken
    };
  }
  async createAccessToken(user: CreateTokenPayLoad) {
    try {
      const data: JwtPayload = {
        user_account: user.user_account,
        user_roles: user.user_roles
      };
      const access_token = await this.jwtService.signAsync(data, {
        secret: user.public_key,
        expiresIn: user.public_expires || '15m'
      });
      return access_token;
    } catch (error) {
      throw new BadRequestException(error);
    }
  }
  async refreshToken(
    token: string,
    user: CreateTokenPayLoad,
    expireAt: Date,
    device: string
  ): Promise<Omit<LoginResponse, 'user'>> {
    const access_token = await this.createAccessToken(user);
    if (!access_token) {
      throw new BadRequestException(appConstant.CREATE_ACCESS_TOKEN_FAILED);
    }
    const now = new Date();

    const isUpdate = (new Date(expireAt).getTime() - now.getTime()) / (1000 * 60 * 60) < 1;
    const newExpireAt = new Date(now.getTime() + 24 * 30 * 60 * 60 * 1000);
    if (isUpdate) {
      await this.userService.updateUserKeyToken(token, newExpireAt);
    }
    console.log('user', user);
    console.log('is update ', isUpdate);
    const userToken = await this.userService.getUserKeyTokenById(user.id as bigint);
    console.log('userToken', userToken);
    if (!userToken) {
      throw new BadRequestException(appConstant.REFRESH_TONKEN_FAIL);
    }
    await this.setUserToken(user.user_id as bigint, userToken, device);
    return {
      access_token,
      refresh_token: userToken.refresh_token
    };
  }
  async oAuth2Login(user: OAuthLoginResponse, ip: string, device: string): Promise<LoginResponse | string> {
    const { email } = user;
    const checkUser: UserBase | null = await this.userService.getUserBaseByAccount(email);
    if (checkUser) {
      const userBaseRoles = await this.userService.findUserRolesByUserId(checkUser.user_id);
      if (!userBaseRoles) {
        throw new BadRequestException(appConstant.USER_NOT_FOUND);
      }
      return await this.loginRespose(userBaseRoles, ip, device);
    }

    return await this.register({
      verify_key: email,
      verify_type: VerifyType.Email,
      verify_purpose: 'OAUTH_USER'
    });
  }
  async validateJwtPayload(user_id: bigint): Promise<UserBaseType | null> {
    const hashKey = generateHash(user_id.toString(), this.salt, 'sha256');
    const userBaseKey = getUserBaseKey(hashKey);
    let userBaseRoles = await this.distributeCache.getObject<UserBaseType>('user_base', userBaseKey);
    if (!userBaseRoles) {
      userBaseRoles = await this.userService.findUserRolesByUserId(user_id);
      if (!userBaseRoles) {
        return null;
      }
      await this.distributeCache.setObjectWithTtl('user_base', userBaseKey, userBaseRoles, 30 * 60); // 30 minutes
      return userBaseRoles;
    }
    return userBaseRoles;
  }
  async logOut(userId: bigint, device: string, id: bigint): Promise<string> {
    const hashKey = generateHash(userId.toString(), this.salt, 'sha256');
    const userKey = getUserHashKey(hashKey, device);
    await this.distributeCache.delObject('user_key_token', userKey);
    await this.userService.deleteUserKeyTokenById(id);
    return appConstant.LOGOUT_SUCCESS;
  }
  setCookie(res: Response, refresh_token: string): void {
    res.cookie('refresh_token', refresh_token, {
      httpOnly: true,
      sameSite: 'lax'
    });
  }
  delCookie(res: Response): void {
    res.cookie('refresh_token', '', {
      httpOnly: true,
      sameSite: 'lax',
      maxAge: 0
    });
  }

  async setUserToken(userId: bigint, userToken: UserKeyToken, device: string): Promise<boolean> {
    const hashKey = generateHash(userId.toString(), this.salt, 'sha256');
    const userKey = getUserHashKey(hashKey, device);
    await this.distributeCache.setObjectWithTtl('user_key_token', userKey, userToken, 60 * 24 * 30); // 30 days
    return true;
  }
  async comparePassword(enteredPassword: string, dbPassword: string): Promise<boolean> {
    return await compare(enteredPassword, dbPassword);
  }
}
