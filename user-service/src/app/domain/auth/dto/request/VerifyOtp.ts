import { ApiProperty } from '@nestjs/swagger';
import { IsDefined, IsString } from 'class-validator';

export default class VerifyOtpRequest {
  @ApiProperty({
    description: 'user verify key',
    example: 'example@gmail.com',
    required: true
  })
  @IsDefined()
  @IsString()
  verify_key: string;
  @ApiProperty({
    description: 'user verify otp',
    example: '123456',
    required: true
  })
  @IsDefined()
  @IsString()
  verify_otp: string;
}
