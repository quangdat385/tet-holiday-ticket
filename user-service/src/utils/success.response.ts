import { Response, Request } from 'express';
import { omit } from 'lodash';
import { Logger } from 'src/logger/log.logger';
import { FileLogger } from 'src/logger/file.logger';
import * as moment from 'moment';
import { HttpStatus } from '@nestjs/common';
import { LogLevel } from 'src/logger/loglevel';

class SuccessResponse<T> {
  message: string;
  status: HttpStatus;
  metadata: T;
  logger: Logger;
  fileLogger = FileLogger.getInstance();
  constructor({ message, statusCode, metadata }: { message: string; statusCode: HttpStatus; metadata: T }) {
    this.message = message;
    this.metadata = metadata;
    this.status = statusCode;
    this.logger = new Logger();
  }

  send = (res: Response, req: Request) => {
    const params = {
      context: req.context,
      requestId: req.correlationId,
      metadata: this.metadata
    };
    this.fileLogger.log(this.generateLogMessage(req, res), params);
    this.logger.log(LogLevel.Info, this.generateLogMessage(req, res));
    return res.status(this.status).json(omit(this, ['logger', 'params', 'fileLogger']));
  };
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
  private generateLogMessage = (req: Request, res: Response): string => {
    const size = this.getResponseSize(res);
    const now = Date.now();
    const timeTaken = now - req.before;
    const terms: { [key: string]: string } = {
      '%h': req.socket.remoteAddress || '-',
      '%l': '-',
      '%x1': `span=${req.span.toString()}`,
      '%x2': `trace=${req.correlationId.toString()}`,
      '%x3': 'type=Incoming request',
      '%u': '-', // todo: parse req.headers.authorization?
      '%t': `date=[${moment().format('DD/MMM/YYYY:HH:mm:ss ZZ')}]`,
      '%r': `request=${req.method} ${req.originalUrl} ${req.httpVersion}`,
      '%>s': `status=${this.status}`,
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

export { SuccessResponse };
