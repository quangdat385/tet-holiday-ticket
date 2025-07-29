// src/app/domain/user/service/impl/user.grpc.service.ts
import { Metadata } from '@grpc/grpc-js';
import { Injectable } from '@nestjs/common';
import { GrpcMethod } from '@nestjs/microservices';
import {
  User,
  UserById,
  UserServiceGrpcController,
  UserServiceGrpcControllerMethods
} from 'src/app/domain/proto/type/user';
import { UserService } from 'src/app/domain/user/service/impl/user.service';

@Injectable()
@UserServiceGrpcControllerMethods()
export class UserGrpcService implements UserServiceGrpcController {
  // eslint-disable-next-line prettier/prettier
  constructor(private readonly userService: UserService) { }
  @GrpcMethod('UserServiceGrpc', 'FindOne')
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  async findOne(request: UserById, metadata: Metadata, ...rest: any): Promise<User> {
    const { userId } = request;
    const message = { ...request, metadata };
    console.log('findOne', message);
    const user = await this.userService.findUserRolesByUserId(BigInt(userId));
    if (!user) {
      throw new Error(`User with ID ${userId} not found`);
    }
    // Map UserBaseType to User
    const grpcUser: User = {
      userId: Number(user.user_id),
      userAccount: user.user_account,
      userSalt: user.user_salt,
      userRoles: user.user_roles.map((role) => ({
        roleId: Number(role.role_id),
        roleName: role.role_name,
        roleDescription: role.role_description ?? '',
        roleMenus: role.role_menus.map((menu) => ({
          menuId: Number(menu.menu_id),
          menuName: menu.menu_name,
          menuUrl: menu.menu_url,
          menuPrefix: menu.menu_prefix,
          menuPid: menu.menu_pid
        }))
      }))
    };
    return grpcUser;
  }
}
