import { S3Client, S3ClientConfig } from '@aws-sdk/client-s3';
import { ConfigService } from '@nestjs/config';

export class S3CloudService {
  // eslint-disable-next-line prettier/prettier
  constructor(private configService: ConfigService) {
  }
  createS3Config(): S3Client {
    const accessKeyId = this.configService.get<string>('aws_key');
    const secretAccessKey = this.configService.get<string>('aws_secret');

    if (!accessKeyId || !secretAccessKey) {
      throw new Error('AWS credentials are not properly configured.');
    }

    const config: S3ClientConfig = {
      region: 'ap-southeast-1',
      credentials: {
        accessKeyId,
        secretAccessKey
      }
    };
    return new S3Client(config);
  }
}
