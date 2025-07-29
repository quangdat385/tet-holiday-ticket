import { Injectable, BadRequestException } from '@nestjs/common';
import { ConfigService } from '@nestjs/config';
import { PutObjectCommand, PutObjectCommandOutput } from '@aws-sdk/client-s3';
import { S3CloudService } from '../provider/s3.provider';
import { randomBytes } from 'crypto';
import { getSignedUrl } from '@aws-sdk/cloudfront-signer';

@Injectable()
export class UploadService {
  // eslint-disable-next-line prettier/prettier
  constructor(private configService: ConfigService) { }
  async uploadImageFromLocalS3(
    file: Express.Multer.File,
    type: string
  ): Promise<{
    url: string;
    result: PutObjectCommandOutput;
    fileName: string;
  }> {
    try {
      const randomKey = () => randomBytes(16).toString('hex');
      const imageName = randomKey();
      const command = new PutObjectCommand({
        Bucket: this.configService.get('bucket_name') as string,
        Key: `${type}/${imageName}`,
        Body: file.buffer,
        ContentType: type
      });

      const s3CloudService = new S3CloudService(this.configService);
      const result = await s3CloudService.createS3Config().send(command);
      console.log(result);
      const expiresIn = 1000 * 60 * 60 * 24;
      const url = this.getCloudfrontImage(imageName, expiresIn);
      return {
        url: url,
        result: result,
        fileName: imageName
      };
    } catch (error: any) {
      throw new BadRequestException(error);
    }
  }
  getCloudfrontImage(imageName: string, expiresIn: number) {
    const urlImagePuclic = this.configService.get<string>('cloudfront_url');
    return getSignedUrl({
      url: `${urlImagePuclic}/${imageName}`,
      keyPairId: this.configService.get<string>('cloudfront_public') as string,
      dateLessThan: new Date(Date.now() + expiresIn).toUTCString(),
      privateKey: this.configService.get<string>('cloudfrin_private') as string
    });
  }
}
