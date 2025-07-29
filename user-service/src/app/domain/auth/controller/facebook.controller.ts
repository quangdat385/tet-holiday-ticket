// Native.

// Package.
import { Controller, Get, HttpCode, HttpStatus, Req, Res, UseGuards, UsePipes, ValidationPipe } from '@nestjs/common';
import {
  ApiBadRequestResponse,
  ApiConsumes,
  ApiExcludeEndpoint,
  ApiOkResponse,
  ApiOperation,
  ApiTags
} from '@nestjs/swagger';
import { Request, Response } from 'express';
import { OAuthLoginResponse } from 'src/app/domain/auth/dto/request/FacebookAuth';
import { BadAuthRequetResponse } from 'src/app/domain/auth/dto/response/common';
import { LoginResponse } from 'src/app/domain/auth/dto/response/Login';
import { RegisterUserDto } from 'src/app/domain/auth/dto/response/Register';
import { AuthService } from 'src/app/domain/auth/service/impl/auth.service';
import { FacebookOauthGuard } from 'src/app/domain/guard/facebook_auth.guard';
import { generateOtp } from 'src/utils/otp';
import { SuccessResponse } from 'src/utils/success.response';

@Controller('auth/facebook')
@UsePipes(
  new ValidationPipe({
    whitelist: true,
    transform: true
  })
)
@ApiTags('auth')
export class FacebookController {
  constructor(
    private readonly authService: AuthService
    // eslint-disable-next-line prettier/prettier
  ) { }
  @Get()
  @HttpCode(HttpStatus.CREATED)
  @ApiOkResponse({ type: RegisterUserDto, description: 'Create User Success' })
  @ApiBadRequestResponse({ type: BadAuthRequetResponse, description: 'Bad Request' })
  @ApiOperation({ description: 'Login User By Facebook() ' })
  @ApiConsumes('application/json')
  @UseGuards(FacebookOauthGuard)
  // eslint-disable-next-line @typescript-eslint/no-unused-vars, prettier/prettier
  async facebookleAuth(@Req() _req: Request) { }
  @ApiExcludeEndpoint()
  @HttpCode(HttpStatus.CREATED)
  @UseGuards(FacebookOauthGuard)
  @Get('callback')
  async facebookAuthRedirect(@Req() req: Request, @Res() res: Response) {
    const ip = req.headers['x-forwarded-for'] || req.socket.remoteAddress;
    const device = req.headers['x-device-id'] ? (req.headers['x-device-id'] as string) : generateOtp().toString();
    const response = await this.authService.oAuth2Login(req.user as OAuthLoginResponse, ip as string, device);
    res.header('x-device-id', device);
    if (response && typeof response === 'string') {
      new SuccessResponse<string>({
        message: 'Send Email Verification To Your Email Successfully',
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
