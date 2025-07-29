import { ApiProperty } from '@nestjs/swagger';
import { IsDefined, IsEnum, IsString } from 'class-validator';
import { VerifyType } from 'src/app/app.constant';

export class RegisterRequest {
  @ApiProperty({
    description: 'user verify key',
    example: 'example@gmail.com',
    required: true
  })
  @IsDefined()
  @IsString()
  verify_key: string;
  @ApiProperty({
    description: 'user verify type',
    example: 1,
    required: true
  })
  @IsDefined()
  @IsEnum(VerifyType)
  verify_type: VerifyType;
  @ApiProperty({
    description: 'user verify purpose',
    example: 'TEST_USER',
    required: true
  })
  @IsDefined()
  @IsString()
  verify_purpose: string;
}
export class ReSendVerifyOtpRequest {
  @ApiProperty({
    description: 'user verify key',
    example: 'example@gmail.com',
    required: true
  })
  @IsDefined()
  @IsString()
  verify_key: string;
  @ApiProperty({
    description: 'user verify type',
    example: 1,
    required: true
  })
  @IsDefined()
  @IsEnum(VerifyType)
  verify_type: VerifyType;
}
