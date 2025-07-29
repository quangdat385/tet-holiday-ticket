import { RedisModule } from '@nestjs-modules/ioredis';
import { Module } from '@nestjs/common';
import { ConfigModule, ConfigService } from '@nestjs/config';
import { AuthModule } from 'src/app/domain/auth/auth.module';
import { DatabaseModule } from 'src/app/domain/database/database.module';
import { UserModule } from 'src/app/domain/user/user.module';
import { AppLoggerModule } from 'src/logger/logger.module';
import { CacheModule } from './cache/cache.module';
import { CacheModule as MemCache } from '@nestjs/cache-manager';
import { UtilsModule } from './utils/utils.module';
import { UserProtoModule } from 'src/app/domain/proto/user.proto.module';

@Module({
  imports: [
    ConfigModule,
    UserModule,
    AppLoggerModule,
    DatabaseModule,
    AuthModule,
    RedisModule.forRootAsync({
      imports: [ConfigModule],
      useFactory: (configService: ConfigService) => ({
        type: 'single',
        url: configService.get<string>('redis_db'),
        maxRetriesPerRequest: 100
      }),
      inject: [ConfigService]
    }),
    CacheModule,
    MemCache.register({
      isGlobal: true
    }),
    UtilsModule,
    UserProtoModule
  ],
  controllers: [],
  providers: [],
  exports: []
})
// eslint-disable-next-line prettier/prettier
export class DomainModule { }
