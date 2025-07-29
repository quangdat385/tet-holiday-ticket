import { ApiResponseProperty } from '@nestjs/swagger';
import { JwtPayload } from 'jsonwebtoken';

export class LoginResponse {
  user: JwtPayload;
  access_token: string;
  refresh_token: string;
}

export class LoginResponseDto {
  @ApiResponseProperty({
    example: 'Create User Successfully'
  })
  public message: string;
  @ApiResponseProperty({
    example: {
      user: {
        user_id: 123,
        user_email: 'test123@gmail.com',
        user_roles: ['SYS-ADMIN', 'ADMIN', 'SHOP', 'USER'],
        user_slug: '6f3c053b-d752-4c34-aae0-b0b6888888',
        user_name: 'test123'
      },
      access_token:
        'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo3LCJ1c2VyX2VtYWlsIjoiY2RkZDRmQHlhaG9vLmNvbSIsInVzZXJfcm9sZXMiOlt7InVzZXJfaWQiOjcsInJvbGVfaWQiOjMsInJvbGVzIjp7InJvbGVfbmFtZSI6IlVTRVIifX1dLCJ1c2VyX3NsdWciOiI2ZjNjMDUzYi1kNzUyLTRjMzQtYWFlMC1iMGI2MWY5OWRhNzciLCJpYXQiOjE3MTU0MTM5MTEsImV4cCI6MTcxNTQxMzkxMX0.NVuNOtgesPIac0rrj5BivgHK3cJX7H8rYprq2dxRI6s',
      refresh_token:
        'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo3LCJ1c2VyX2VtYWlsIjoiY2RkZDRmQHlhaG9vLmNvbSIsInVzZXJfcm9sZXMiOlt7InVzZXJfaWQiOjcsInJvbGVfaWQiOjMsInJvbGVzIjp7InJvbGVfbmFtZSI6IlVTRVIifX1dLCJ1c2VyX3NsdWciOiI2ZjNjMDUzYi1kNzUyLTRjMzQtYWFlMC1iMGI2MWY5OWRhNzciLCJpYXQiOjE3MTU0MTM5MTEsImV4cCI6MTcxNTUwMDMxMX0.z6ocxvclMUDIjnWRhPDLs0FLBmOsCw9CNKwCaNuh8yE'
    }
  })
  public metadata: LoginResponse;
  @ApiResponseProperty({
    example: 201
  })
  public status: number;
}
