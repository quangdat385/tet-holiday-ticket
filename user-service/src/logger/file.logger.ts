import { createLogger, format, transports, Logger as WinstonLogger } from 'winston';
import 'winston-daily-rotate-file';
import { Injectable } from '@nestjs/common';

/**
 * @Author :quangdat385
 *
 */
@Injectable()
export class FileLogger {
  private static instance: FileLogger;
  logger: WinstonLogger;
  constructor() {
    const formatPrint = format.printf(({ level, message, context, requestId, timestamp, metadata }) => {
      const data = JSON.stringify(metadata, (key: string, value: unknown) => {
        if (typeof value === 'bigint') {
          return value.toString();
        }
        return value;
      });
      // eslint-disable-next-line @typescript-eslint/restrict-template-expressions
      return `${level}::${message}::${JSON.stringify(context)}::${requestId}::${timestamp}::${data}`;
    });
    this.logger = createLogger({
      format: format.combine(format.timestamp({ format: 'YYYY-MM-DD HH:mm:ss' }), formatPrint),
      transports: [
        new transports.Console(),
        new transports.DailyRotateFile({
          dirname: 'src/logs',
          filename: 'application-%DATE%.info.log',
          datePattern: 'YYYY-MM-DD-HH-mm',
          zippedArchive: true,
          maxSize: '20m',
          maxFiles: '14d',
          format: format.combine(format.timestamp({ format: 'YYYY-MM-DD HH:mm:ss' }), formatPrint),
          level: 'info'
        }),
        new transports.DailyRotateFile({
          dirname: 'src/logs',
          filename: 'application-%DATE%.error.log',
          datePattern: 'YYYY-MM-DD-HH-mm',
          zippedArchive: true,
          maxSize: '20m',
          maxFiles: '14d',
          format: format.combine(format.timestamp({ format: 'YYYY-MM-DD HH:mm:ss' }), formatPrint),
          level: 'error'
        })
      ]
    });
  }
  commonParams(params: { [key: string]: unknown }) {
    return {
      ...params
    };
  }
  log(message: string, params: { [key: string]: unknown }) {
    const paramsLog = this.commonParams(params);
    const logObject = Object.assign({ message }, paramsLog);
    this.logger.info(logObject);
  }
  error(message: string, params: { [key: string]: unknown }) {
    const paramsLog = this.commonParams(params);
    const logObject = Object.assign({ message }, paramsLog);
    this.logger.error(logObject);
  }
  public static getInstance(): FileLogger {
    if (this.instance == null) {
      this.instance = new FileLogger();
    }
    return this.instance;
  }
}
