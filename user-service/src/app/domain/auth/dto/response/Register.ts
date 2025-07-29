import { ApiResponseProperty } from '@nestjs/swagger';

export class RegisterUserDto {
  @ApiResponseProperty({
    example: 'Send Verify Email Successfully'
  })
  public message: string;
  @ApiResponseProperty({
    example: 'Send Verify Email Successfully'
  })
  public metadata: string;
  @ApiResponseProperty({
    example: 200
  })
  public status: number;
}
