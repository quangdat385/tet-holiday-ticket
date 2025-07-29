import { ApiResponseProperty } from '@nestjs/swagger';

export class UpdateAvatarResponse {
  @ApiResponseProperty({
    example: 'Update Avatar Successfully'
  })
  public message: string;
  @ApiResponseProperty({
    example: 'avatar.png'
  })
  public metadata: string;
  @ApiResponseProperty({
    example: 200
  })
  public status: number;
}
