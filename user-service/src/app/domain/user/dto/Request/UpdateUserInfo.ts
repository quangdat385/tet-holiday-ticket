import { ApiProperty } from '@nestjs/swagger';
import { IsDefined, IsNumber, IsOptional, IsString } from 'class-validator';

export class UpdateUserInfoRequest {
  @ApiProperty({
    description: 'user id',
    example: 123456789,
    required: true
  })
  @IsNumber()
  @IsDefined()
  user_id: bigint;
  @ApiProperty({
    description: 'user account',
    example: 'user123',
    required: false
  })
  @IsString()
  @IsOptional()
  user_account?: string;
  @ApiProperty({
    description: 'user nickname',
    example: 'User Name',
    required: true
  })
  @IsString()
  @IsOptional()
  user_nickname?: string;
  @ApiProperty({
    description: 'user avatar',
    example: 'avatar.jpg',
    required: false
  })
  @IsString()
  @IsOptional()
  user_avatar?: string;
  @ApiProperty({
    description: 'user state',
    example: 1,
    required: false
  })
  @IsNumber()
  @IsOptional()
  user_state?: number;
  @ApiProperty({
    description: 'user gender',
    example: 1,
    required: false
  })
  @IsNumber()
  @IsOptional()
  user_gender?: number;
  @ApiProperty({
    description: 'user mobile',
    example: '1234567890',
    required: false
  })
  user_mobile?: string;
  @ApiProperty({
    description: 'user email',
    example: 'example@yahoo.com',
    required: false
  })
  user_email?: string;
}
