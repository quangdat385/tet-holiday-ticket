import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { ConfigModule } from '@nestjs/config';
import { AppLoggerModule } from 'src/logger/logger.module';
import configuration from 'src/config/configuration';
import { DomainModule } from 'src/app/domain/domain.module';
import { TerminusModule } from '@nestjs/terminus';
import { HttpModule } from '@nestjs/axios';
import { ThrottlerModule } from '@nestjs/throttler';
@Module({
  imports: [
    ConfigModule.forRoot({
      load: [configuration],
      isGlobal: true
    }),
    AppLoggerModule,
    DomainModule,
    TerminusModule,
    HttpModule,
    ThrottlerModule.forRoot({
      throttlers: [
        {
          ttl: 1000,
          limit: 5
        }
      ]
    })
  ],
  controllers: [AppController],
  providers: []
})
// eslint-disable-next-line prettier/prettier
export class AppModule { }
