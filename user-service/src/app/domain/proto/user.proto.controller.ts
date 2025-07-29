import { Metadata } from '@grpc/grpc-js';
import { Controller, Get, HttpStatus, Inject, OnModuleInit, Param, Req, Res } from '@nestjs/common';
import { ClientGrpc, GrpcMethod } from '@nestjs/microservices';
import { ApiExcludeController } from '@nestjs/swagger';
import { Request, Response } from 'express';
import { firstValueFrom } from 'rxjs';
import { PreGoTicket99999, TICKET_SERVICE_NAME, TicketServiceClient } from 'src/app/domain/proto/type/ticket';
import { User, USER_SERVICE_GRPC_SERVICE_NAME, UserById } from 'src/app/domain/proto/type/user';
import { UserGrpcService } from 'src/app/domain/proto/user.proto.service';
import { UserService } from 'src/app/domain/user/service/impl/user.service';
import { FileLogger } from 'src/logger/file.logger';
import { SuccessResponse } from 'src/utils/success.response';

@Controller('proto')
@ApiExcludeController()
export class UserProtoController implements OnModuleInit {
  private ticketServiceGrpc: TicketServiceClient;
  private userServiceGrpc: UserGrpcService;
  constructor(
    @Inject('TICKET_SERVICE') private client: ClientGrpc,
    @Inject('USER_SERVICE') private userServiceClient: ClientGrpc,
    private readonly logger: FileLogger,
    private readonly userService: UserService
    // eslint-disable-next-line prettier/prettier
  ) { }
  onModuleInit() {
    this.ticketServiceGrpc = this.client.getService<TicketServiceClient>(TICKET_SERVICE_NAME);
    this.userServiceGrpc = this.userServiceClient.getService<UserGrpcService>(USER_SERVICE_GRPC_SERVICE_NAME);
    console.log('gRPC service initialized:', Object.keys(this.ticketServiceGrpc), Object.keys(this.userServiceGrpc));
  }
  @Get('get-ticket/:ticket_id')
  async getTicketGrpc(@Param('ticket_id') ticket_id: number, @Res() res: Response, @Req() req: Request) {
    const context = { param: { ticket_id } };
    req['context'] = context;
    const data: PreGoTicket99999 = await firstValueFrom(
      this.ticketServiceGrpc.getTicket({ id: ticket_id }, new Metadata())
    );
    new SuccessResponse<PreGoTicket99999>({
      message: 'Register successfully',
      metadata: data,
      statusCode: HttpStatus.OK
    }).send(res, req);
  }
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
