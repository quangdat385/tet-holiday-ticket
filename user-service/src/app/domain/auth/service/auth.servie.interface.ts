import { UserKeyToken } from '@prisma/client';
import { CreateTokenPayLoad } from 'src/app/domain/auth/dto/request/CreateToken';
import { LoginRequest } from 'src/app/domain/auth/dto/request/Login';
import { RegisterRequest } from 'src/app/domain/auth/dto/request/Register';
import { UpdatePasswordRegisterRequest } from 'src/app/domain/auth/dto/request/UpdatePasswordRegister';
import VerifyOtpRequest from 'src/app/domain/auth/dto/request/VerifyOtp';
import { LoginResponse } from 'src/app/domain/auth/dto/response/Login';
import { Response } from 'express';

export interface AuthServiceInterface {
  register(registerRequest: RegisterRequest): Promise<string>;
  verifyOtp(verifyOtpRequet: VerifyOtpRequest): Promise<{ token: string }>;
  reSendVerifyOtp(reSendVerifyOtp: Omit<RegisterRequest, 'verify_purpose'>): Promise<string>;
  updatePasswordRegister(updatePasswordRegister: UpdatePasswordRegisterRequest): Promise<string>;
  updatePassword(updatePassword: UpdatePasswordRegisterRequest): Promise<string>;
  forgotPassword(forgotPassword: Omit<RegisterRequest, 'verify_purpose'>): Promise<{ message: string; state: number }>;
  verifyForgotPassword(verifyForgotPassword: VerifyOtpRequest): Promise<{ token: string }>;
  login(loginRequest: LoginRequest, ip: string, device: string): Promise<LoginResponse>;
  logOut(userId: bigint, device: string, id: bigint): Promise<string>;
  refreshToken(
    token: string,
    user: CreateTokenPayLoad,
    expireAt: Date,
    device: string
  ): Promise<Omit<LoginResponse, 'user'>>;
  createToKen(payload: CreateTokenPayLoad): Promise<LoginResponse>;
  comparePassword(enteredPassword: string, dbPassword: string): Promise<boolean>;
  setUserToken(userId: bigint, userToken: UserKeyToken, device: string): Promise<boolean>;
  setCookie(res: Response, refresh_token: string): void;
  delCookie(res: Response): void;
}
