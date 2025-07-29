import { CanActivate, ExecutionContext, Injectable, UnauthorizedException } from '@nestjs/common';
import { JwtService } from '@nestjs/jwt';
import { AuthService } from 'src/app/domain/auth/service/impl/auth.service';
import { DistributeCacheService } from 'src/app/domain/cache/service/impl/cache.distribute';
import { UserService } from 'src/app/domain/user/service/impl/user.service';
import { Request } from 'express';
import { UserKeyToken } from '@prisma/client';
import { JwtPayload } from 'jsonwebtoken';
import { generateHash, getUserHashKey } from 'src/utils/crypto';
import { generateOtp } from 'src/utils/otp';

@Injectable()
export class AccessTokenGuard implements CanActivate {
  constructor(
    private jwtService: JwtService,
    private authService: AuthService,
    private userService: UserService,
    private distributedCacheService: DistributeCacheService
    // eslint-disable-next-line prettier/prettier
  ) { }
  async canActivate(context: ExecutionContext): Promise<boolean> {
    const request = context.switchToHttp().getRequest<Request>();
    const device = request.headers['x-device-id']
      ? (request.headers['x-device-id'] as string)
      : generateOtp().toString();
    const accessToken: string | null = this.extractTokenFromHeader(request);
    const user_id: bigint | null = this.extractClientIdFromHeader(request);
    console.log('accessToken', accessToken);
    console.log('user_id', user_id);
    if (!accessToken || !user_id) {
      throw new UnauthorizedException('Unauthorized access token or user_id');
    }
    const hashKey = generateHash(user_id.toString(), this.authService.getSalt(), 'sha256');
    const userKey = getUserHashKey(hashKey, device);
    const keyToken: UserKeyToken | null = await this.distributedCacheService.getObject('user_key_token', userKey);
    if (!keyToken) {
      throw new UnauthorizedException('Key token not found');
    }
    try {
      const payload: JwtPayload = await this.jwtService.verifyAsync(accessToken, {
        secret: keyToken.public_key
      });
      if (!payload) {
        throw new UnauthorizedException();
      }
    } catch (error: any) {
      throw new UnauthorizedException(error);
    }
    const user = await this.authService.validateJwtPayload(user_id);
    if (!user) {
      throw new UnauthorizedException('User not found');
    }
    request['user'] = user;
    request['token_id'] = keyToken.id;
    request['device'] = device;
    return true;
  }
  private extractTokenFromHeader(request: Request): string | null {
    const [type, token] = (request.headers['authorization'] as string)?.split(' ') ?? [];
    return type === 'Bearer' ? token : null;
  }
  private extractClientIdFromHeader(request: Request): bigint | null {
    const client_id = request.headers['x-client-id'];
    console.log('client_id', client_id);
    if (!client_id) {
      return null;
    }
    return typeof client_id === 'string' ? BigInt(client_id) : null;
  }
}
