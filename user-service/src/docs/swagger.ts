import { INestApplication } from '@nestjs/common';
import { DocumentBuilder, SwaggerModule } from '@nestjs/swagger';
import * as basicAuth from 'express-basic-auth';
import { ConfigService } from '@nestjs/config';
import { SWAGGER_CONFIG } from './swagger.config';

/**
 * Creates an OpenAPI document for an application, via swagger.
 * @param app the nestjs application
 * @returns the OpenAPI document
 */
const SWAGGER_ENVS = ['local', 'development', 'production'];

export function createDocument(app: INestApplication) {
  const builder = new DocumentBuilder()
    .setTitle(SWAGGER_CONFIG.title)
    .addBearerAuth({ type: 'http', scheme: 'bearer', bearerFormat: 'JWT' }, 'authorization')
    .setDescription(SWAGGER_CONFIG.description)
    .setVersion(SWAGGER_CONFIG.version)
    .addCookieAuth('refresh_token');
  for (const tag of SWAGGER_CONFIG.tags) {
    builder.addTag(tag);
  }
  const options = builder.build();
  const env: string = app.get(ConfigService).get<string>('env') || 'development';
  const { username, password }: { username: string; password: string } = app
    .get(ConfigService)
    .get<{ username: string; password: string }>('swagger') || {
    username: '',
    password: ''
  };
  if (SWAGGER_ENVS.includes(env)) {
    app.use(
      'ticket-user/api/v1/docs',
      basicAuth({
        challenge: true,
        users: {
          [username]: password
        }
      })
    );
    const document = SwaggerModule.createDocument(app, options);
    SwaggerModule.setup('ticket-user/api/v1/docs', app, document);
  }
}
