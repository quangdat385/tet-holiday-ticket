import { BadRequestException, CanActivate, ExecutionContext, ForbiddenException, Injectable } from '@nestjs/common';
import { JwtService } from '@nestjs/jwt';
import { AuthService } from 'src/app/domain/auth/service/impl/auth.service';
import { DistributeCacheService } from 'src/app/domain/cache/service/impl/cache.distribute';
import { UserService } from 'src/app/domain/user/service/impl/user.service';
import { Request } from 'express';
import { generateHash, getUserHashKey } from 'src/utils/crypto';
import { UserKeyToken } from '@prisma/client';
import { JwtPayload } from 'jsonwebtoken';
import { Role } from 'src/app/domain/auth/dto/request/CreateToken';
import { generateOtp } from 'src/utils/otp';

@Injectable()
export class RefreshTokenGuard implements CanActivate {
  constructor(
    private jwtService: JwtService,
    private authService: AuthService,
    private userService: UserService,
    private distributedCacheService: DistributeCacheService
    // eslint-disable-next-line prettier/prettier
  ) { }

  async canActivate(context: ExecutionContext): Promise<boolean> {
    const request = context.switchToHttp().getRequest<Request>();
    const refreshToken: string | null = this.extractRefreshTokenFromCookies(request);
    const user_id: bigint | null = this.extractClientIdFromHeader(request);
    const device = request.headers['x-device-id']
      ? (request.headers['x-device-id'] as string)
      : generateOtp().toString();
    console.log('device', device);
    console.log('user_id', user_id);
    console.log('refreshToken', refreshToken);
    if (!refreshToken || !user_id) {
      throw new ForbiddenException();
    }
    const hashKey = generateHash(user_id.toString(), this.authService.getSalt(), 'sha256');
    const userKey = getUserHashKey(hashKey, device);
    const keyToken: UserKeyToken | null = await this.distributedCacheService.getObject('user_key_token', userKey);
    if (!keyToken) {
      throw new ForbiddenException();
    }
    try {
      const payload: JwtPayload = await this.jwtService.verifyAsync(refreshToken, {
        secret: keyToken.private_key
      });
      if (!payload) {
        throw new ForbiddenException();
      }
    } catch (err: any) {
      throw new BadRequestException(err);
    }

    const user = await this.authService.validateJwtPayload(user_id);
    if (!user) {
      throw new ForbiddenException();
    }
    const roles = user.user_roles.map((role) => {
      return {
        role_name: role.role_name,
        role_description: role.role_description
      };
    });
    request['token_payload'] = {
      id: keyToken.id,
      user_id: user.user_id,
      user_account: user.user_account,
      user_roles: roles as Role[],
      public_key: keyToken.public_key,
      private_key: keyToken.private_key
    };
    request['refresh_token'] = refreshToken;
    request['exprire_at'] = keyToken.expireAt;
    return true;
  }
  private extractRefreshTokenFromCookies(request: Request): string | null {
    // eslint-disable-next-line @typescript-eslint/no-unsafe-assignment
    const cookies = request.cookies['refresh_token'];
    if (!cookies) {
      return null;
    }
    return cookies as string;
  }

  private extractClientIdFromHeader(request: Request): bigint | null {
    const client_id = request.headers['x-client-id'];
    if (!client_id) {
      return null;
    }
    return typeof client_id === 'string' ? BigInt(client_id) : null;
  }
}
