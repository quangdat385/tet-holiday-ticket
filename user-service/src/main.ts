import { HttpAdapterHost, NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';
import { INestApplication, ValidationPipe } from '@nestjs/common';
import { Transport, GrpcOptions } from '@nestjs/microservices';
import { ConfigService } from '@nestjs/config';
import { createDocument } from 'src/docs/swagger';
import { join } from 'path';
import { ReflectionService } from '@grpc/reflection';
import * as cookieParser from 'cookie-parser';
import { AllExceptionsFilter } from 'src/filter/all-exceptions.filter';
import helmet from 'helmet';

async function bootstrap(): Promise<void> {
  const app: INestApplication = await NestFactory.create<INestApplication>(AppModule);
  app.enableCors({
    origin: '*', // Thay bằng URL frontend của bạn
    methods: 'GET,HEAD,PUT,PATCH,POST,DELETE',
    credentials: true // Nếu cần gửi cookie
  });
  app.useGlobalPipes(new ValidationPipe());
  app.connectMicroservice<GrpcOptions>({
    transport: Transport.GRPC,
    options: {
      package: 'user',
      protoPath: join(process.cwd(), 'src/app/domain/proto/user.proto'),
      url: '127.0.0.1:50050',
      onLoadPackageDefinition: (pkg, server) => {
        // eslint-disable-next-line @typescript-eslint/no-unsafe-argument
        new ReflectionService(pkg).addToServer(server);
      }
    }
  });
  await app.startAllMicroservices();
  // app.connectMicroservice<MicroserviceOptions>({
  //   transport: Transport.NATS,
  //   options: {
  //     servers: [app.get(ConfigService).get<string>('nats_uri') || 'nats://localhost:4222'],
  //     pass: app.get(ConfigService).get<string>('nats_pass') || '',
  //     user: app.get(ConfigService).get<string>('nats_user') || ''
  //   }
  // });
  const globalPrefix = 'ticket-user/api/v1';
  const { httpAdapter } = app.get(HttpAdapterHost);
  app.useGlobalFilters(new AllExceptionsFilter(httpAdapter));
  app.setGlobalPrefix(globalPrefix);
  app.use(cookieParser());
  app.use(helmet());
  createDocument(app);
  const configService = app.get(ConfigService);
  const PORT: number = configService.get<number>('port') || 8080;
  await app.listen(PORT, () => {
    console.log(`User service is running on port ${PORT}`);
  });
}
bootstrap().catch((error) => {
  console.error('Error starting the application:', error);
  process.exit(1); // Exit the process with a failure code
});
