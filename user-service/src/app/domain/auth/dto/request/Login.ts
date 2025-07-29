import { ApiProperty } from '@nestjs/swagger';
import { IsDefined, IsString, IsStrongPassword } from 'class-validator';

export class LoginRequest {
  @ApiProperty({
    description: 'user account',
    example: 'example@yahoo.com',
    required: true
  })
  @IsDefined()
  @IsString()
  user_account: string;
  @ApiProperty({
    description: 'user password',
    example: '12345678@default',
    required: true
  })
  @IsDefined()
  @IsStrongPassword({
    minLength: 8,
    minLowercase: 1,
    minUppercase: 1,
    minNumbers: 1,
    minSymbols: 1
  })
  @IsString()
  password: string;
}
