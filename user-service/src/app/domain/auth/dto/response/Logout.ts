import { ApiResponseProperty } from '@nestjs/swagger';

export class LogOutResponse {
  @ApiResponseProperty({
    example: 'Log Out Successfully'
  })
  public message: string;
  @ApiResponseProperty({
    example: {
      message: 'Logout Successful'
    }
  })
  public metadata: object;
  @ApiResponseProperty({
    example: 200
  })
  public status: number;
}
