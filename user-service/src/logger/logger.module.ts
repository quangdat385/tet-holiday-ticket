import { MiddlewareConsumer, Module, NestModule } from '@nestjs/common';
import { ConfigModule } from '@nestjs/config';
import { LoggerMiddleware } from 'src/logger/logger.middleware';
import { Logger } from './log.logger';
import { FileLogger } from './file.logger';

@Module({
  imports: [ConfigModule],
  controllers: [],
  providers: [Logger, FileLogger],
  exports: [Logger, FileLogger]
})
export class AppLoggerModule implements NestModule {
  public configure(consumer: MiddlewareConsumer) {
    consumer.apply(LoggerMiddleware).forRoutes('*');
  }
}
