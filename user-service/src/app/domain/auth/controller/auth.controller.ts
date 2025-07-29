import {
  Controller,
  Post,
  Body,
  UsePipes,
  ValidationPipe,
  Res,
  Req,
  HttpCode,
  HttpStatus,
  UseGuards
} from '@nestjs/common';
import {
  ApiBadRequestResponse,
  ApiBearerAuth,
  ApiConsumes,
  ApiCookieAuth,
  ApiHeader,
  ApiInternalServerErrorResponse,
  ApiNotFoundResponse,
  ApiOkResponse,
  ApiOperation,
  ApiTags,
  ApiUnauthorizedResponse
} from '@nestjs/swagger';
import { RegisterRequest, ReSendVerifyOtpRequest } from 'src/app/domain/auth/dto/request/Register';
import { AuthService } from 'src/app/domain/auth/service/impl/auth.service';
import { Response, Request } from 'express';
import { SuccessResponse } from 'src/utils/success.response';
import { RegisterUserDto } from 'src/app/domain/auth/dto/response/Register';
import {
  BadAuthRequetResponse,
  NotFoundResponse,
  UnauthorizedAuthResponse
} from 'src/app/domain/auth/dto/response/common';
import VerifyOtpResponse from 'src/app/domain/auth/dto/response/VerifyOtp';
import VerifyOtpRequest from 'src/app/domain/auth/dto/request/VerifyOtp';
import { UpdatePasswordRegisterRequest } from 'src/app/domain/auth/dto/request/UpdatePasswordRegister';
import { UpdatePasswordRegisterResponse } from 'src/app/domain/auth/dto/response/UpdatePasswordRegister';
import { ResendVerifyOtpResponse } from 'src/app/domain/auth/dto/response/ResendVerifyOtp';
import { ForgotPasswordResponse } from 'src/app/domain/auth/dto/response/ForgotPassword';
import { VerifyForgotPasswordResponse } from 'src/app/domain/auth/dto/response/VerifyForgotPassword';
import { LoginResponse, LoginResponseDto } from 'src/app/domain/auth/dto/response/Login';
import { LogOutResponse } from 'src/app/domain/auth/dto/response/Logout';
import { AccessTokenGuard } from 'src/app/domain/guard/access.token.guard';
import { UserBaseType } from 'src/app/domain/auth/auth.type';
import { LoginRequest } from 'src/app/domain/auth/dto/request/Login';
import { RefreshTokenGuard } from 'src/app/domain/guard/refresh.token.guard';
import { CreateTokenPayLoad } from 'src/app/domain/auth/dto/request/CreateToken';
import { UpdatePasswordResponse } from 'src/app/domain/auth/dto/response/UpdatePassword';
import { generateOtp } from 'src/utils/otp';

@Controller('auth')
@UsePipes(
  new ValidationPipe({
    whitelist: true,
    transform: true
  })
)
@ApiTags('auth')
export class AuthController {
  constructor(
    private readonly authService: AuthService
    // eslint-disable-next-line prettier/prettier
  ) { }
  @Post('register')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ description: 'user register' })
  @ApiOkResponse({
    description: 'user registered successfully',
    type: RegisterUserDto
  })
  @ApiInternalServerErrorResponse({
    description: 'internal server error occurred'
  })
  @ApiBadRequestResponse({ description: 'bad request', type: BadAuthRequetResponse })
  @ApiConsumes('application/json')
  async register(@Body() registerRequest: RegisterRequest, @Res() res: Response, @Req() req: Request): Promise<any> {
    const context = { body: registerRequest };
    req['context'] = context;
    new SuccessResponse<string>({
      message: 'Register successfully',
      metadata: await this.authService.register(registerRequest),
      statusCode: HttpStatus.OK
    }).send(res, req);
  }
  @Post('verify-otp')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ description: 'verify otp' })
  @ApiOkResponse({
    description: 'verify otp successfully',
    type: VerifyOtpResponse
  })
  @ApiInternalServerErrorResponse({
    description: 'internal server error occurred'
  })
  @ApiBadRequestResponse({ description: 'bad request', type: BadAuthRequetResponse })
  @ApiConsumes('application/json')
  async verifyOtp(@Body() verifyOtpRequest: VerifyOtpRequest, @Res() res: Response, @Req() req: Request): Promise<any> {
    const context = { body: verifyOtpRequest };
    console.log('verifyOtpRequest', verifyOtpRequest);
    req['context'] = context;
    new SuccessResponse<{ token: string }>({
      message: 'Verify OTP successfully',
      metadata: await this.authService.verifyOtp(verifyOtpRequest),
      statusCode: HttpStatus.OK
    }).send(res, req);
  }
  @Post('update-password-register')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ description: 'update password register' })
  @ApiOkResponse({
    description: 'update password register successfully',
    type: UpdatePasswordRegisterResponse
  })
  @ApiInternalServerErrorResponse({
    description: 'internal server error occurred'
  })
  @ApiBadRequestResponse({ description: 'bad request', type: BadAuthRequetResponse })
  @ApiConsumes('application/json')
  async updatePasswordRegister(
    @Body() updatePasswordRegister: UpdatePasswordRegisterRequest,
    @Res() res: Response,
    @Req() req: Request
  ): Promise<any> {
    const context = { ...updatePasswordRegister };
    req['context'] = context;
    new SuccessResponse<string>({
      message: 'Update password register successfully',
      metadata: await this.authService.updatePasswordRegister(updatePasswordRegister),
      statusCode: HttpStatus.OK
    }).send(res, req);
  }
  @Post('update-password')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ description: 'update password' })
  @ApiOkResponse({
    description: 'update password register successfully',
    type: UpdatePasswordResponse
  })
  @ApiInternalServerErrorResponse({
    description: 'internal server error occurred'
  })
  @ApiBadRequestResponse({ description: 'bad request', type: BadAuthRequetResponse })
  @ApiConsumes('application/json')
  async updatePassword(
    @Body() updatePassword: UpdatePasswordRegisterRequest,
    @Res() res: Response,
    @Req() req: Request
  ): Promise<any> {
    const context = { ...updatePassword };
    req['context'] = context;
    new SuccessResponse<string>({
      message: 'Update password register successfully',
      metadata: await this.authService.updatePassword(updatePassword),
      statusCode: HttpStatus.OK
    }).send(res, req);
  }
  @Post('resend-verify-otp')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ description: 'resend verify otp' })
  @ApiOkResponse({
    description: 'resend verify otp successfully',
    type: ResendVerifyOtpResponse
  })
  @ApiInternalServerErrorResponse({
    description: 'internal server error occurred'
  })
  @ApiBadRequestResponse({ description: 'bad request', type: BadAuthRequetResponse })
  @ApiConsumes('application/json')
  async resendVerifyOtp(
    @Body() resendVerifyOtp: ReSendVerifyOtpRequest,
    @Res() res: Response,
    @Req() req: Request
  ): Promise<any> {
    const context = { body: resendVerifyOtp };
    req['context'] = context;
    new SuccessResponse<string>({
      message: 'Resend verify otp successfully',
      metadata: await this.authService.reSendVerifyOtp(resendVerifyOtp),
      statusCode: HttpStatus.OK
    }).send(res, req);
  }
  @Post('forgot-password')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ description: 'forget password' })
  @ApiOkResponse({
    description: 'forget password successfully',
    type: ForgotPasswordResponse
  })
  @ApiInternalServerErrorResponse({
    description: 'internal server error occurred'
  })
  @ApiBadRequestResponse({ description: 'bad request', type: BadAuthRequetResponse })
  @ApiConsumes('application/json')
  async forgetPassword(
    @Body() forgetPassword: ReSendVerifyOtpRequest,
    @Res() res: Response,
    @Req() req: Request
  ): Promise<any> {
    const context = { body: forgetPassword };
    req['context'] = context;
    new SuccessResponse<{
      message: string;
      state: number;
    }>({
      message: 'Forget password successfully',
      metadata: await this.authService.forgotPassword(forgetPassword),
      statusCode: HttpStatus.OK
    }).send(res, req);
  }
  @Post('verify-forgot-password')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ description: 'verify forgot password' })
  @ApiOkResponse({
    description: 'verify forgot password successfully',
    type: VerifyForgotPasswordResponse
  })
  @ApiInternalServerErrorResponse({
    description: 'internal server error occurred'
  })
  @ApiBadRequestResponse({ description: 'bad request', type: BadAuthRequetResponse })
  @ApiConsumes('application/json')
  async verifyForgotPassword(
    @Body() verifyForgotPassword: VerifyOtpRequest,
    @Res() res: Response,
    @Req() req: Request
  ): Promise<any> {
    const context = { body: verifyForgotPassword };
    req['context'] = context;
    new SuccessResponse<{ token: string }>({
      message: 'Verify OTP forgot password successfully',
      metadata: await this.authService.verifyForgotPassword(verifyForgotPassword),
      statusCode: HttpStatus.OK
    }).send(res, req);
  }
  @Post('login')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ description: 'user login' })
  @ApiOkResponse({
    description: 'user has been login successfully',
    type: LoginResponseDto
  })
  @ApiInternalServerErrorResponse({
    description: 'internal server error occurred'
  })
  @ApiBadRequestResponse({ description: 'bad request', type: BadAuthRequetResponse })
  @ApiConsumes('application/json')
  @ApiHeader({
    name: 'x-device-id',
    description: 'device id',
    required: true,
    example: '1234567890'
  })
  async login(@Body() loginRequest: LoginRequest, @Res() res: Response, @Req() req: Request): Promise<any> {
    const ipAddress = req.headers['x-forwarded-for'] || req.socket.remoteAddress;
    const device = req.headers['x-device-id'] ? (req.headers['x-device-id'] as string) : generateOtp().toString();
    const context = { body: loginRequest, ipAddress, device };
    req['context'] = context;
    const data = await this.authService.login(loginRequest, ipAddress as string, device);
    this.authService.setCookie(res, data.refresh_token);
    res.header('x-device-id', device);
    new SuccessResponse<LoginResponse>({
      message: 'Login successfully',
      metadata: data,
      statusCode: HttpStatus.OK
    }).send(res, req);
  }
  @Post('logout')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ description: 'user logout' })
  @ApiOkResponse({
    description: 'user has been logout successfully',
    type: LogOutResponse
  })
  @ApiInternalServerErrorResponse({
    description: 'internal server error occurred'
  })
  @ApiBadRequestResponse({ description: 'bad request', type: BadAuthRequetResponse })
  @ApiNotFoundResponse({
    description: 'not found',
    type: NotFoundResponse
  })
  @ApiUnauthorizedResponse({
    description: 'unauthorized access token or user_id',
    type: UnauthorizedAuthResponse
  })
  @ApiBearerAuth('authorization')
  @ApiHeader({
    name: 'x-client-id',
    description: 'client id',
    required: true,
    example: '1234567890'
  })
  @ApiHeader({
    name: 'x-device-id',
    description: 'device id',
    required: true,
    example: '1234567890'
  })
  @ApiConsumes('application/json')
  @UseGuards(AccessTokenGuard)
  async logout(@Res() res: Response, @Req() req: Request): Promise<any> {
    const device = req.device;
    req['context'] = { token: req.access_token };
    const user: UserBaseType = req['user'] as UserBaseType;
    const id = req['token_id'];
    const data = await this.authService.logOut(user.user_id, device, id);
    this.authService.delCookie(res);
    res.header('x-device-id', device);
    new SuccessResponse<string>({
      message: 'Logout successfully',
      metadata: data,
      statusCode: HttpStatus.OK
    }).send(res, req);
  }
  @Post('refresh-token')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ description: 'refresh token' })
  @ApiOkResponse({
    description: 'refresh token successfully',
    type: LoginResponseDto
  })
  @ApiInternalServerErrorResponse({
    description: 'internal server error occurred'
  })
  @ApiBadRequestResponse({ description: 'bad request', type: BadAuthRequetResponse })
  @ApiNotFoundResponse({
    description: 'not found',
    type: NotFoundResponse
  })
  @ApiUnauthorizedResponse({
    description: 'unauthorized access token or user_id',
    type: UnauthorizedAuthResponse
  })
  @ApiHeader({
    name: 'x-client-id',
    description: 'client id',
    required: true,
    example: '1234567890'
  })
  @ApiHeader({
    name: 'x-device-id',
    description: 'device id',
    required: true,
    example: '1234567890'
  })
  @ApiCookieAuth('refresh_token')
  @ApiConsumes('application/json')
  @UseGuards(RefreshTokenGuard)
  async refreshToken(@Res() res: Response, @Req() req: Request): Promise<any> {
    const token: string = req.refresh_token as string;
    let payload: CreateTokenPayLoad = req['token_payload'] as CreateTokenPayLoad;
    const expireAt: Date = req['exprire_at'] as Date;
    const device = req.device;
    const context = {
      refresh_token: token,
      expireAt,
      payload: {
        ...payload,
        user_id: (payload.user_id as string | bigint).toString()
      }
    };
    req['context'] = context;
    payload = {
      ...payload,
      user_id: BigInt(payload.user_id as string)
    };
    const data = await this.authService.refreshToken(token, payload, expireAt, device);
    this.authService.setCookie(res, data.refresh_token);
    res.header('x-device-id', device);
    new SuccessResponse<{
      access_token: string;
    }>({
      message: 'Refresh token successfully',
      metadata: {
        access_token: data.access_token
      },
      statusCode: HttpStatus.OK
    }).send(res, req);
  }
}
