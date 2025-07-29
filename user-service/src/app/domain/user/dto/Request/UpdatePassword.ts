import { ApiProperty } from '@nestjs/swagger';
import { IsDefined, IsString, IsStrongPassword } from 'class-validator';

export class UpdatePasswordRequest {
  @ApiProperty({
    description: 'password',
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
  @ApiProperty({
    description: 'user confirm password',
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
  confirm_password: string;
}
