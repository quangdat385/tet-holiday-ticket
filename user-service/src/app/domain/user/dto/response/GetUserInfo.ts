import { ApiResponseProperty } from '@nestjs/swagger';
import { UserInfo } from '@prisma/client';

export class UserInfoResponse {
  user_account: string;
  user_email: string | null;
  user_nickname: string | null;
  user_avatar: string | null;
  user_birthday: Date | null;
  user_id: number;
  user_state: number;
  user_gender: number;
  user_mobile: string | null;
  user_is_authentication: boolean;
}

export class GetUserInfoResponse {
  @ApiResponseProperty({
    example: 'Get User Info Successfully'
  })
  public message: string;
  @ApiResponseProperty({
    example: {
      user_id: 1,
      user_account: 'user_account',
      user_name: 'user_name',
      user_email: 'user_email',
      user_mobile: 'user_phone',
      user_nickname: 'user_nickname',
      user_avatar: 'user_avatar',
      user_gender: 'male',
      user_birthday: 'user_birthday',
      user_address: 'user_address',
      user_is_authentication: true
    }
  })
  public metadata: UserInfo;
  @ApiResponseProperty({
    example: 200
  })
  public status: number;
}
