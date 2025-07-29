import { forwardRef, Module } from '@nestjs/common';
import { AuthService } from './service/impl/auth.service';
import { AuthController } from './controller/auth.controller';
import { UserModule } from 'src/app/domain/user/user.module';
import { JwtModule } from '@nestjs/jwt';
import { CacheModule } from 'src/app/domain/cache/cache.module';
import { AppLoggerModule } from 'src/logger/logger.module';
import { FacebookController } from 'src/app/domain/auth/controller/facebook.controller';
import { GoogleController } from 'src/app/domain/auth/controller/google.controller';
import { FacebookOauthStrategy } from 'src/app/domain/strategy/facebook_jwt.strategy';
import { GoogleOauthStrategy } from 'src/app/domain/strategy/google_jwt.strategy.ts';

@Module({
  imports: [forwardRef(() => UserModule), JwtModule.register({}), CacheModule, AppLoggerModule],
  controllers: [AuthController, FacebookController, GoogleController],
  providers: [AuthService, FacebookOauthStrategy, GoogleOauthStrategy],
  exports: [AuthService, FacebookOauthStrategy, GoogleOauthStrategy]
})
// eslint-disable-next-line prettier/prettier
export class AuthModule { }
