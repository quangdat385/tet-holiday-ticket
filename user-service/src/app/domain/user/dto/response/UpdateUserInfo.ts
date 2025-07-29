import { ApiResponseProperty } from '@nestjs/swagger';
import { UserInfoResponse } from 'src/app/domain/user/dto/response/GetUserInfo';

export class UpdateUserInfoResponse {
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
  public metadata: UserInfoResponse;
  @ApiResponseProperty({
    example: 200
  })
  public status: number;
}
