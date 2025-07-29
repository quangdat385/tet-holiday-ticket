import { IsArray, IsDefined, IsOptional, IsString } from 'class-validator';

export class CreateTokenPayLoad {
  @IsOptional()
  user_id?: bigint | string;
  @IsDefined()
  @IsString()
  user_account: string;
  @IsDefined()
  @IsArray()
  user_roles: Role[];
  @IsDefined()
  @IsString()
  public public_key: string;
  @IsDefined()
  @IsString()
  public private_key: string;
  @IsOptional()
  @IsString()
  public public_expires?: string;
  @IsOptional()
  @IsString()
  public private_expires?: string;
  [key: string]: unknown;
}

export interface Role {
  role_name: string;
  role_description: string;
}
