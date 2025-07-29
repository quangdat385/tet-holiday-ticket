import { ApiResponseProperty } from '@nestjs/swagger';

export class ForgotPasswordResponse {
  @ApiResponseProperty({
    example: 'Forgot password successfully'
  })
  public message: string;
  @ApiResponseProperty({
    example: {
      message: 'Forgot password successfully',
      state: 4
    }
  })
  public metadata: string;
  @ApiResponseProperty({
    example: 200
  })
  public status: number;
}
