import { ApiProperty } from '@nestjs/swagger';
import { IsDefined, IsNumber, IsString } from 'class-validator';

export class CreateUserMenu {
  @ApiProperty({
    description: 'user menu_name',
    example: 'User',
    required: true
  })
  @IsDefined()
  @IsString()
  menu_name: string;
  @ApiProperty({
    description: 'user menu_pid',
    example: '1234567890',
    required: true
  })
  @IsDefined()
  @IsString()
  menu_pid: string;
  @ApiProperty({
    description: 'user menu_prefix',
    example: 'user',
    required: true
  })
  @IsDefined()
  @IsString()
  menu_prefix: string;
  @ApiProperty({
    description: 'user menu_url',
    example: '/user',
    required: true
  })
  @IsDefined()
  @IsString()
  menu_url: string;
  @ApiProperty({
    description: 'user user_created',
    example: 123456789,
    required: true
  })
  @IsDefined()
  @IsNumber()
  user_created: bigint;
}
