import { ApiResponseProperty } from '@nestjs/swagger';

export class VerifyForgotPasswordResponse {
  @ApiResponseProperty({
    example: 'Verify forgot password successfully'
  })
  public message: string;
  @ApiResponseProperty({
    example: {
      message: 'Verify forgot password successfully'
    }
  })
  public metadata: string;
  @ApiResponseProperty({
    example: 200
  })
  public status: number;
}
