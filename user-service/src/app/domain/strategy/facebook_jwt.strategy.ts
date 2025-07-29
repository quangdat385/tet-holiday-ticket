/* eslint-disable @typescript-eslint/no-unsafe-assignment */
import { Injectable } from '@nestjs/common';
import { ConfigService } from '@nestjs/config';
import { PassportStrategy } from '@nestjs/passport';
import { Profile, Strategy } from 'passport-facebook';

@Injectable()
export class FacebookOauthStrategy extends PassportStrategy(Strategy, 'facebook') {
  constructor(configService: ConfigService) {
    const { oauth_facebook_id, oauth_facebook_secret, oauth_facebook_callback } = configService.get('facebook') ?? {
      oauth_facebook_id: '',
      oauth_facebook_secret: '',
      oauth_facebook_callback: ''
    };
    super({
      clientID: oauth_facebook_id,
      clientSecret: oauth_facebook_secret,
      callbackURL: oauth_facebook_callback,
      profileFields: ['email', 'photos', 'id', 'displayName']
    });
  }

  validate(_accessToken: string, _refreshToken: string, profile: Profile) {
    const { id, name, emails } = profile;
    return {
      provider: 'facebook',
      providerId: id,
      uid: id,
      name: name?.givenName,
      username: emails && emails.length > 0 ? emails[0].value : undefined,
      email: emails && emails.length > 0 ? emails[0].value : undefined
    };
  }
}
