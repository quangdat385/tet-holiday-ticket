import { ApiResponseProperty } from '@nestjs/swagger';

export class ResendVerifyOtpResponse {
  @ApiResponseProperty({
    example: 'Resend verify otp successfully'
  })
  public message: string;
  @ApiResponseProperty({
    example: 'Resend verify otp successfully'
  })
  public metadata: string;
  @ApiResponseProperty({
    example: 200
  })
  public status: number;
}
