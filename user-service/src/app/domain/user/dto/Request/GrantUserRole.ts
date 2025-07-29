import { ApiProperty } from '@nestjs/swagger';
import { IsDefined, IsNumber } from 'class-validator';
import { UserRoleNumber } from 'src/app/domain/user/service/user.service.interface';

export class GrantUserRoleRequest {
  @ApiProperty({
    description: 'user id',
    example: 1234567890,
    required: true
  })
  @IsDefined()
  @IsNumber()
  user_id: bigint;
  @ApiProperty({
    description: 'user role id',
    example: 1234567890,
    required: true
  })
  @IsDefined()
  @IsNumber()
  role: UserRoleNumber;
}
