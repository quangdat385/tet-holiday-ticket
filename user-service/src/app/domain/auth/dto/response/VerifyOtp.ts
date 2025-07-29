import { ApiResponseProperty } from '@nestjs/swagger';

export default class VerifyOtpResponse {
  @ApiResponseProperty({
    example: 'Send Verify Email Successfully'
  })
  public message: string;
  @ApiResponseProperty({
    example: {
      token: '1234566668971'
    }
  })
  public metadata: {
    token: string;
  };
  @ApiResponseProperty({
    example: 200
  })
  public status: number;
}
