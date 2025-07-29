import { CanActivate, ExecutionContext, Injectable, UnauthorizedException } from '@nestjs/common';
import { Reflector } from '@nestjs/core';
import { Request } from 'express';
import { UserBaseType } from 'src/app/domain/auth/auth.type';
import { UserRoles, UserRoleType } from 'src/app/domain/user/user.type';

@Injectable()
export class RolesGuard implements CanActivate {
  // eslint-disable-next-line prettier/prettier
  constructor(private reflector: Reflector) { }

  canActivate(context: ExecutionContext): boolean {
    const roles = this.reflector.getAllAndMerge<UserRoles[]>('roles', [context.getClass(), context.getHandler()]) || [];
    console.log('roles', roles);

    if (roles && roles.length === 0) {
      return true;
    }
    const request = context.switchToHttp().getRequest<Request>();
    const user: UserBaseType = request.user as UserBaseType;

    if (user && !user.user_roles) {
      return false;
    }
    const hasRole = () =>
      user.user_roles.some((role: UserRoleType) => roles.find((i) => i === (role.role_name as UserRoles)));
    if (!hasRole()) throw new UnauthorizedException('You not allowed to access this resources');
    return true;
  }
}
