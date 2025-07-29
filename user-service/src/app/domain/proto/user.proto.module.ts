import { Module } from '@nestjs/common';
import { ClientsModule, Transport } from '@nestjs/microservices';
import { join } from 'path';
import { UserProtoController } from 'src/app/domain/proto/user.proto.controller';
import { UserGrpcService } from 'src/app/domain/proto/user.proto.service';
import { UserModule } from 'src/app/domain/user/user.module';
import { AppLoggerModule } from 'src/logger/logger.module';

@Module({
  imports: [
    ClientsModule.register([
      {
        name: 'TICKET_SERVICE',
        transport: Transport.GRPC,
        options: {
          package: 'ticket',
          protoPath: join(process.cwd(), 'src/app/domain/proto/ticket.proto'),
          url: 'localhost:50051'
        }
      },
      {
        name: 'USER_SERVICE',
        transport: Transport.GRPC,
        options: {
          package: 'user',
          protoPath: join(process.cwd(), 'src/app/domain/proto/user.proto'),
          url: 'localhost:50000'
        }
      }
    ]),
    AppLoggerModule,
    UserModule
  ],
  controllers: [UserProtoController],
  providers: [UserGrpcService]
})
// eslint-disable-next-line prettier/prettier
export class UserProtoModule { }
