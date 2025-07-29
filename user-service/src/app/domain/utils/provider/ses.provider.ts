import { SESClient, SESClientConfig } from '@aws-sdk/client-ses';
import { ConfigService } from '@nestjs/config';

export class SeSService {
  // eslint-disable-next-line prettier/prettier
  constructor(private configService: ConfigService) {
  }
  createSeSConfig(): SESClient {
    const accessKeyId = this.configService.get<string>('aws_key');
    const secretAccessKey = this.configService.get<string>('aws_secret');
    if (!accessKeyId || !secretAccessKey) {
      throw new Error('AWS credentials are not properly configured.');
    }
    const config: SESClientConfig = {
      region: 'ap-southeast-1',
      credentials: {
        accessKeyId,
        secretAccessKey
      }
    };
    return new SESClient(config);
  }
}
