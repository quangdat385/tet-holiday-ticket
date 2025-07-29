import { forwardRef, Module } from '@nestjs/common';
import { ConfigModule } from '@nestjs/config';
import { JwtModule } from '@nestjs/jwt';
import { AuthModule } from 'src/app/domain/auth/auth.module';
import { CacheModule } from 'src/app/domain/cache/cache.module';
import { DatabaseModule } from 'src/app/domain/database/database.module';
import { UserController } from 'src/app/domain/user/controller/user.controller';
import { UserService } from 'src/app/domain/user/service/impl/user.service';
import { UtilsModule } from 'src/app/domain/utils/utils.module';
import { AppLoggerModule } from 'src/logger/logger.module';

@Module({
  imports: [
    CacheModule,
    DatabaseModule,
    forwardRef(() => AuthModule),
    UtilsModule,
    ConfigModule,
    JwtModule.register({}),
    AppLoggerModule
  ],
  controllers: [UserController],
  providers: [UserService],
  exports: [UserService]
})
// eslint-disable-next-line prettier/prettier
export class UserModule { }
