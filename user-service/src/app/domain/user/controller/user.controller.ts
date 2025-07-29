import {
  Body,
  Controller,
  Get,
  HttpCode,
  HttpStatus,
  MaxFileSizeValidator,
  Param,
  ParseFilePipe,
  Patch,
  Post,
  Req,
  Res,
  UploadedFiles,
  UseGuards,
  UseInterceptors,
  UsePipes,
  ValidationPipe
} from '@nestjs/common';
import {
  ApiBadRequestResponse,
  ApiBearerAuth,
  ApiConsumes,
  ApiHeader,
  ApiInternalServerErrorResponse,
  ApiOkResponse,
  ApiOperation,
  ApiUnauthorizedResponse
} from '@nestjs/swagger';
import { AuthService } from 'src/app/domain/auth/service/impl/auth.service';
import {
  BadRequetResponse,
  RequestSuccessResponse,
  UnauthorizedResponse
} from 'src/app/domain/user/dto/response/common';
import { GetUserInfoResponse, UserInfoResponse } from 'src/app/domain/user/dto/response/GetUserInfo';
import { UserService } from 'src/app/domain/user/service/impl/user.service';
import { SuccessResponse } from 'src/utils/success.response';
import { Request, Response } from 'express';
import { AccessTokenGuard } from 'src/app/domain/guard/access.token.guard';
import { UpdatePasswordRequest } from 'src/app/domain/user/dto/Request/UpdatePassword';
import { RolesGuard } from 'src/app/domain/guard/role.guard';
import { RoleAllowed } from 'src/app/domain/decorator/roles.decorator';
import { UserRoles } from 'src/app/domain/user/user.type';
import { GrantUserRoleRequest } from 'src/app/domain/user/dto/Request/GrantUserRole';
import { FilesInterceptor } from '@nestjs/platform-express';
import { storage, imageFilter } from 'src/utils/upload.file';
import { UpdateUserInfoRequest } from 'src/app/domain/user/dto/Request/UpdateUserInfo';
import { UserInfo } from '@prisma/client';
import { File } from 'buffer';
import { join } from 'path';
import { UpdateUserInfoResponse } from 'src/app/domain/user/dto/response/UpdateUserInfo';
import { UpdatePasswordResponse } from 'src/app/domain/auth/dto/response/UpdatePassword';
import { UpdateAvatarResponse } from 'src/app/domain/user/dto/response/UploadAvatar';
import { UpdateAvatarRequest } from 'src/app/domain/user/dto/Request/UpdateAvatar';
import { ContactUs } from 'src/app/domain/user/dto/Request/ContactUs';
@Controller('user')
@UsePipes(
  new ValidationPipe({
    whitelist: true,
    transform: true
  })
)
export class UserController {
  constructor(
    private readonly authService: AuthService,
    private readonly userService: UserService
    // eslint-disable-next-line prettier/prettier
  ) { }
  @Get('get-user-info/:user_id')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ description: 'get user info' })
  @ApiOkResponse({
    description: 'get user info successfully',
    type: GetUserInfoResponse
  })
  @ApiInternalServerErrorResponse({
    description: 'internal server error occurred'
  })
  @ApiBadRequestResponse({ description: 'bad request', type: BadRequetResponse })
  @ApiUnauthorizedResponse({
    description: 'unauthorized access token or user_id',
    type: UnauthorizedResponse
  })
  @ApiConsumes('application/json')
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
  @UseGuards(AccessTokenGuard)
  async getUserInfo(@Param('user_id') user_id: bigint, @Res() res: Response, @Req() req: Request): Promise<any> {
    const context = { param: { user_id } };
    req['context'] = context;
    res.header('x-device-id', req.device);
    new SuccessResponse<UserInfoResponse | null>({
      message: 'Get user info successfully',
      metadata: await this.userService.getUserInfoByUserId(user_id),
      statusCode: HttpStatus.OK
    }).send(res, req);
  }
  @Patch('update-password/:user_id')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ description: 'user update password' })
  @ApiOkResponse({
    description: 'user update password successfully',
    type: UpdatePasswordResponse
  })
  @ApiInternalServerErrorResponse({
    description: 'internal server error occurred'
  })
  @ApiBadRequestResponse({ description: 'bad request', type: BadRequetResponse })
  @ApiUnauthorizedResponse({
    description: 'unauthorized access token or user_id',
    type: UnauthorizedResponse
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
  @UseGuards(AccessTokenGuard, RolesGuard)
  @RoleAllowed(UserRoles.ADMIN, UserRoles.USER)
  async updatePassword(
    @Param('user_id') user_id: bigint,
    @Body() body: UpdatePasswordRequest,
    @Res() res: Response,
    @Req() req: Request
  ): Promise<any> {
    const context = { body: body, token: req.access_token };
    req['context'] = context;
    res.header('x-device-id', req.device);
    new SuccessResponse<boolean>({
      message: 'Update password successfully',
      metadata: await this.userService.updateUserPassword(user_id, body),
      statusCode: HttpStatus.OK
    }).send(res, req);
  }
  @Post('grant-user-role')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ description: 'user grant user role' })
  @ApiOkResponse({
    description: 'user grant user role successfully',
    type: RequestSuccessResponse
  })
  @ApiInternalServerErrorResponse({
    description: 'internal server error occurred'
  })
  @ApiBadRequestResponse({ description: 'bad request', type: BadRequetResponse })
  @ApiUnauthorizedResponse({
    description: 'unauthorized access token or user_id',
    type: UnauthorizedResponse
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
  @UseGuards(AccessTokenGuard, RolesGuard)
  @RoleAllowed(UserRoles.ADMIN)
  async grantUserRole(@Body() body: GrantUserRoleRequest, @Res() res: Response, @Req() req: Request): Promise<any> {
    const context = { body };
    req['context'] = context;
    res.header('x-device-id', req.device);
    new SuccessResponse<string>({
      message: 'Grant user role successfully',
      metadata: await this.userService.grantUserRole(body.user_id, body.role),
      statusCode: HttpStatus.OK
    }).send(res, req);
  }
  @Post('upload-avatar')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ description: 'user upload avatar' })
  @ApiOkResponse({
    description: 'user upload avatar successfully',
    type: UpdateAvatarResponse
  })
  @ApiInternalServerErrorResponse({
    description: 'internal server error occurred'
  })
  @ApiBadRequestResponse({ description: 'bad request', type: BadRequetResponse })
  @ApiUnauthorizedResponse({
    description: 'unauthorized access token or user_id',
    type: UnauthorizedResponse
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
  @ApiConsumes('multipart/form-data')
  @UseGuards(AccessTokenGuard, RolesGuard)
  @RoleAllowed(UserRoles.USER)
  @UseInterceptors(FilesInterceptor('avatar', 1, { storage, fileFilter: imageFilter }))
  uploadAvatar(
    @UploadedFiles(
      new ParseFilePipe({
        validators: [new MaxFileSizeValidator({ maxSize: 1000000 })]
      })
    )
    files: [Express.Multer.File],
    @Body() _body: UpdateAvatarRequest, // Replace 'any' with the appropriate DTO for updating user
    @Res() res: Response,
    @Req() req: Request
  ) {
    const context = { param: { files } };
    req['context'] = context;
    res.header('x-device-id', req.device);
    new SuccessResponse<string>({
      message: 'Upload avatar successfully',
      metadata: files[0].path,
      statusCode: HttpStatus.OK
    }).send(res, req);
  }
  @Patch('update-user-info/:user_id')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ description: 'update user' })
  @ApiOkResponse({
    description: 'update user successfully',
    type: UpdateUserInfoResponse
  })
  @ApiInternalServerErrorResponse({
    description: 'internal server error occurred'
  })
  @ApiBadRequestResponse({ description: 'bad request', type: BadRequetResponse })
  @ApiUnauthorizedResponse({
    description: 'unauthorized access token or user_id',
    type: UnauthorizedResponse
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
  @UseGuards(AccessTokenGuard, RolesGuard)
  @RoleAllowed(UserRoles.USER)
  updateUserInfo(
    @Param('user_id') user_id: bigint,
    @Body() body: UpdateUserInfoRequest, // Replace 'any' with the appropriate DTO for updating user
    @Res() res: Response,
    @Req() req: Request
  ) {
    const context = { param: { user_id }, body };
    req['context'] = context;
    res.header('x-device-id', req.device);
    new SuccessResponse<Promise<UserInfo>>({
      message: 'Upload avatar successfully',
      metadata: this.userService.updateUserInfo(user_id, body),
      statusCode: HttpStatus.OK
    }).send(res, req);
  }
  @Get('get-user-avatar/:avatar')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ description: 'get user avatar' })
  @ApiOkResponse({
    description: 'get user avatar successfully',
    type: File
  })
  @ApiInternalServerErrorResponse({
    description: 'internal server error occurred'
  })
  @ApiBadRequestResponse({ description: 'bad request', type: BadRequetResponse })
  @ApiUnauthorizedResponse({
    description: 'unauthorized access token or user_id',
    type: UnauthorizedResponse
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
  @UseGuards(AccessTokenGuard, RolesGuard)
  @RoleAllowed(UserRoles.USER)
  getUserAvatar(@Param('avatar') avatar: string, @Res() res: Response, @Req() req: Request) {
    const avatarPath = join(process.cwd(), 'src', 'public', 'img', 'avatar', avatar);
    const context = { param: { avatar }, avatarPath };
    req['context'] = context;
    res.header('x-device-id', req.device);
    res.sendFile(avatarPath, (err) => {
      if (err) {
        res.status(404).json({ message: 'File not found' });
      }
    });
  }
  @Post('contact-us')
  @HttpCode(HttpStatus.OK)
  @ApiOperation({ description: 'contact us' })
  @ApiOkResponse({
    description: 'contact us successfully',
    type: RequestSuccessResponse
  })
  @ApiInternalServerErrorResponse({
    description: 'internal server error occurred'
  })
  @ApiBadRequestResponse({ description: 'bad request', type: BadRequetResponse })
  @ApiConsumes('application/json')
  async contactUs(@Body() body: ContactUs, @Res() res: Response, @Req() req: Request): Promise<any> {
    const context = { body };
    req['context'] = context;
    new SuccessResponse<string>({
      message: 'Contact us successfully',
      metadata: await this.userService.contactUs(body),
      statusCode: HttpStatus.OK
    }).send(res, req);
  }
}
