/* eslint-disable @typescript-eslint/no-unsafe-assignment */
import { Injectable } from '@nestjs/common';
import { ConfigService } from '@nestjs/config';
import { PassportStrategy } from '@nestjs/passport';
import { Profile, Strategy } from 'passport-google-oauth20';

@Injectable()
export class GoogleOauthStrategy extends PassportStrategy(Strategy, 'google') {
  constructor(configService: ConfigService) {
    const { oauth_google_id, oauth_google_secret, oauth_google_callback } = configService.get('google') ?? {
      oauth_google_id: '',
      oauth_google_secret: '',
      oauth_google_callback: ''
    };
    super({
      clientID: oauth_google_id,
      clientSecret: oauth_google_secret,
      callbackURL: oauth_google_callback,
      passReqToCallback: true,
      scope: ['profile', 'email']
    });
  }
  validate(_request: any, _accessToken: string, _refreshToken: string, profile: Profile) {
    const { id, name, emails } = profile;
    return {
      provider: 'google',
      providerId: id,
      uid: id,
      name: name?.givenName,
      username: emails && emails.length > 0 ? emails[0].value : undefined,
      email: emails && emails.length > 0 ? emails[0].value : undefined
    };
  }
}
