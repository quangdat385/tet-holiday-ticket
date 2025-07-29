import { Module } from '@nestjs/common';
import { SendMailService } from 'src/app/domain/utils/service/send-email.service';
import { UploadService } from 'src/app/domain/utils/service/upload.service';

@Module({
  providers: [UploadService, SendMailService],
  exports: [UploadService, SendMailService]
})
// eslint-disable-next-line prettier/prettier
export class UtilsModule { }
