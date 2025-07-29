import { Injectable } from '@nestjs/common';
import { SendEmailCommand, SendEmailCommandOutput } from '@aws-sdk/client-ses';
import { SeSService } from '../provider/ses.provider';
import { ConfigService } from '@nestjs/config';

@Injectable()
export class SendMailService {
  // eslint-disable-next-line prettier/prettier
  constructor(private configService: ConfigService) {
  }
  private createSendEmailHtmlCommand({
    fromAddress,
    toAddresses,
    // ccAddresses = [],
    body,
    subject
    // replyToAddresses = []
  }: {
    fromAddress: string;
    toAddresses: string | string[];
    // ccAddresses?: string | string[];
    body: string;
    subject: string;
    // replyToAddresses?: string | string[];
  }): SendEmailCommand {
    return new SendEmailCommand({
      Destination: {
        /* required */
        // CcAddresses: ccAddresses instanceof Array ? ccAddresses : [ccAddresses],
        ToAddresses: toAddresses instanceof Array ? toAddresses : [toAddresses]
      },
      Message: {
        /* required */
        Body: {
          /* required */
          Html: {
            Charset: 'UTF-8',
            Data: body
          }
        },
        Subject: {
          Charset: 'UTF-8',
          Data: subject
        }
      },
      Source: fromAddress
      // ReplyToAddresses: replyToAddresses instanceof Array ? replyToAddresses : [replyToAddresses]
    });
  }
  private createSendEmailCommand({
    fromAddress,
    toAddresses,
    ccAddresses = [],
    body,
    subject,
    replyToAddresses = []
  }: {
    fromAddress: string;
    toAddresses: string | string[];
    ccAddresses?: string | string[];
    body: string;
    subject: string;
    replyToAddresses?: string | string[];
  }): SendEmailCommand {
    return new SendEmailCommand({
      Destination: {
        /* required */
        CcAddresses: ccAddresses instanceof Array ? ccAddresses : [ccAddresses],
        ToAddresses: toAddresses instanceof Array ? toAddresses : [toAddresses]
      },
      Message: {
        /* required */
        Body: {
          /* required */
          Text: {
            Charset: 'UTF-8',
            Data: body
          }
        },
        Subject: {
          Charset: 'UTF-8',
          Data: subject
        }
      },
      Source: fromAddress,
      ReplyToAddresses: replyToAddresses instanceof Array ? replyToAddresses : [replyToAddresses]
    });
  }
  async sendVerifyEmail(
    fromAdress: string,
    toAddress: string,
    subject: string,
    body: string
  ): Promise<SendEmailCommandOutput> {
    const sesService = new SeSService(this.configService);
    const sendEmailCommand = this.createSendEmailCommand({
      fromAddress: fromAdress,
      toAddresses: toAddress,
      body,
      subject
    });
    return await sesService.createSeSConfig().send(sendEmailCommand);
  }
  async sendVerifyEmailHtml(
    fromAdress: string,
    toAddress: string,
    subject: string,
    body: string
  ): Promise<SendEmailCommandOutput> {
    const sesService = new SeSService(this.configService);
    const sendEmailCommand = this.createSendEmailHtmlCommand({
      fromAddress: fromAdress,
      toAddresses: toAddress,
      body,
      subject
    });
    return await sesService.createSeSConfig().send(sendEmailCommand);
  }
}
