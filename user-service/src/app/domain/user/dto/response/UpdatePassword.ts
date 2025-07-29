import { ApiResponseProperty } from '@nestjs/swagger';

export class UpdatePasswordResponse {
  @ApiResponseProperty({
    example: 'Update Password Successfully'
  })
  public message: string;
  @ApiResponseProperty({
    example: 'update password successfully'
  })
  public metadata: string;
  @ApiResponseProperty({
    example: 200
  })
  public status: number;
}
