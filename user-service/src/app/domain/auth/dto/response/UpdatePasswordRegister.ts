import { ApiResponseProperty } from '@nestjs/swagger';

export class UpdatePasswordRegisterResponse {
  @ApiResponseProperty({
    example: 'Update Password Register Successfully'
  })
  public message: string;
  @ApiResponseProperty({
    example: 'Update Password Register Successfully'
  })
  public metadata: string;
  @ApiResponseProperty({
    example: 200
  })
  public status: number;
}
