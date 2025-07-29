import { ApiProperty } from '@nestjs/swagger';

export class UpdateAvatarRequest {
  @ApiProperty({
    description: 'upate avatar',
    type: 'array',
    items: {
      type: 'string',
      format: 'binary'
    },
    required: true
  })
  avatar: Express.Multer.File[];
}
