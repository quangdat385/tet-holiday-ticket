import { ApiProperty } from '@nestjs/swagger';
import { IsDefined, IsEmail, IsString } from 'class-validator';

export class ContactUs {
  @ApiProperty({
    description: 'First name of the user',
    example: 'John',
    type: String
  })
  @IsString()
  @IsDefined()
  firstName: string;
  @ApiProperty({
    description: 'Last name of the user',
    example: 'Doe',
    type: String
  })
  @IsString()
  @IsDefined()
  lastName: string;
  @ApiProperty({
    description: 'Email address of the user',
    example: 'example@gmail.com',
    type: String
  })
  @IsEmail()
  @IsDefined()
  email: string;
  @ApiProperty({
    description: 'Phone number of the user',
    example: '+1234567890',
    type: String
  })
  @IsString()
  @IsDefined()
  phone: string;
  @ApiProperty({
    description: 'Message from the user',
    example: 'I have a question about my account.',
    type: String
  })
  @IsString()
  @IsDefined()
  message: string;
}
