import { Controller, Get, HttpCode, HttpStatus, Req, Res, UseGuards, UsePipes, ValidationPipe } from '@nestjs/common';
import { ApiExcludeEndpoint, ApiTags } from '@nestjs/swagger';
import { Request, Response } from 'express';
import { OAuthLoginResponse } from 'src/app/domain/auth/dto/request/FacebookAuth';
import { LoginResponse } from 'src/app/domain/auth/dto/response/Login';
import { AuthService } from 'src/app/domain/auth/service/impl/auth.service';
import { GoogleOauthGuard } from 'src/app/domain/guard/google_auth.guard';
import { generateOtp } from 'src/utils/otp';
import { SuccessResponse } from 'src/utils/success.response';
@Controller('auth/google')
@UsePipes(
  new ValidationPipe({
    whitelist: true,
    transform: true
  })
)
@ApiTags('auth')
export class GoogleController {
  constructor(
    private readonly authService: AuthService
    // eslint-disable-next-line prettier/prettier
  ) { }
  @Get()
  @UseGuards(GoogleOauthGuard)
  // eslint-disable-next-line @typescript-eslint/no-unused-vars, prettier/prettier
  async googleAuth(@Req() _req: Request) { }
  @ApiExcludeEndpoint()
  @HttpCode(HttpStatus.CREATED)
  @Get('callback')
  @UseGuards(GoogleOauthGuard)
  async googleAuthRedirect(@Req() req: Request, @Res() res: Response) {
    const ip = req.headers['x-forwarded-for'] || req.socket.remoteAddress;
    const device = req.headers['x-device-id'] ? (req.headers['x-device-id'] as string) : generateOtp().toString();
    const response = await this.authService.oAuth2Login(req.user as OAuthLoginResponse, ip as string, device);
    res.header('x-device-id', device);
    console.log('response', response);
    if (response && typeof response === 'string') {
      new SuccessResponse<string>({
        message: 'Login successfully',
        metadata: response,
        statusCode: HttpStatus.OK
      }).send(res, req);
    } else if (response && typeof response === 'object') {
      this.authService.setCookie(res, response.refresh_token);
      new SuccessResponse<LoginResponse>({
        message: 'Login successfully',
        metadata: response,
        statusCode: HttpStatus.OK
      }).send(res, req);
      return res.json(response);
    }
  }
}
