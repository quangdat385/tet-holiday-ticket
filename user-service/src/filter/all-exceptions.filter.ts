import { Catch, ArgumentsHost, HttpStatus, HttpException } from '@nestjs/common';
import { BaseExceptionFilter } from '@nestjs/core';
import { Response } from 'express';
import { omit } from 'lodash';
import { PrismaClientValidationError } from '@prisma/client/runtime/library';
import * as moment from 'moment';
import { JsonWebTokenError } from '@nestjs/jwt';
import { WsException } from '@nestjs/websockets';
import { Logger } from 'src/logger/log.logger';
import { FileLogger } from 'src/logger/file.logger';
import { Request } from 'express';

type MyResponseObj = {
  statusCode: number;
  timestamp: string;
  path: string;
  response: string | object;
};

@Catch()
export class AllExceptionsFilter extends BaseExceptionFilter {
  private readonly logger: Logger = new Logger();
  private readonly fileLogger: FileLogger = FileLogger.getInstance();

  catch(exception: unknown, host: ArgumentsHost) {
    const ctx = host.switchToHttp();
    const response = ctx.getResponse<Response>();
    const request = ctx.getRequest<Request>();

    const myResponseObj: MyResponseObj = {
      statusCode: 500,
      timestamp: new Date().toISOString(),
      path: request.url,
      response: ''
    };
    // Add more Prisma Error Types if you want
    if (exception instanceof HttpException) {
      myResponseObj.statusCode = exception.getStatus();
      myResponseObj.response = exception.getResponse();
    } else if (exception instanceof PrismaClientValidationError) {
      myResponseObj.statusCode = 422;
      myResponseObj.response = omit(exception, ['stack']);
    } else if (exception instanceof JsonWebTokenError) {
      myResponseObj.statusCode = 401;
      myResponseObj.response = omit(exception, ['stack']);
    } else if (exception instanceof WsException) {
      myResponseObj.statusCode = 401;
      myResponseObj.response = HttpStatus.UNAUTHORIZED as unknown as string;
    } else {
      myResponseObj.statusCode = HttpStatus.INTERNAL_SERVER_ERROR;
      myResponseObj.response = 'Internal Server Error';
    }
    const params = {
      context: request.context,
      requestId: request.correlationId,
      metadata: myResponseObj.response
    };
    this.fileLogger.error(this.generateLogMessage(request, response, myResponseObj.statusCode), params);
    this.logger.error(this.generateLogMessage(request, response, myResponseObj.statusCode));
    response.status(myResponseObj.statusCode).json(myResponseObj);

    super.catch(exception, host);
  }
  private getResponseSize(res: Response): number {
    const sizeRaw = res.getHeader('Content-Length');
    if (typeof sizeRaw === 'number') {
      return sizeRaw;
    }
    if (typeof sizeRaw === 'string') {
      const parsed = parseInt(sizeRaw, 10);
      if (isNaN(parsed)) {
        return 0;
      }
      return parsed;
    }
    return 0;
  }
  /*
  date=${moment().format('DD/MMM/YYYY:HH:mm:ss ZZ')} trace=${id} type=IncomingRequest endpoint=${req.originalUrl} duration=${duration} span=${span} status=${res.statusCode}
   */
  private generateLogMessage = (req: Request, res: Response, status: number): string => {
    const size = this.getResponseSize(res);
    const now = Date.now();
    const timeTaken = now - req.before;
    const terms: { [key: string]: string } = {
      '%h': req.socket.remoteAddress || '-',
      '%l': '-',
      '%x1': `span=${req.span ? req.span.toString() : '-'}`,
      '%x2': `trace=${req.correlationId ? req.correlationId.toString() : '-'}`,
      '%x3': 'type=Incoming request',
      '%u': '-', // todo: parse req.headers.authorization?
      '%t': `date=[${moment().format('DD/MMM/YYYY:HH:mm:ss ZZ')}]`,
      '%r': `request=${req.method} ${req.originalUrl} ${req.httpVersion}`,
      '%>s': `status=${status}`,
      '%b': size === 0 ? 'size=-' : `size=${size}`,
      '%tt': `duration=${timeTaken}`
    };
    let str = '%t %x2 %x3 "%r" %x1 %>s %b %tt';
    for (const term in terms) {
      if (term in terms) {
        str = str.replace(term, terms[term]);
      }
    }
    // eslint-disable-next-line no-useless-escape
    str = str.replace(/%\{([a-zA-Z\-]+)\}i/g, (_match, p1) => {
      const header = req.headers[`${p1}`.toLowerCase()];
      if (header == null) {
        return '-';
      }
      if (Array.isArray(header)) {
        return `"${header.join(',')}"`;
      }
      return `"${header}"`;
    });
    return str;
  };
}
