import { Injectable, NestMiddleware } from '@nestjs/common';
import { Request, Response } from 'express';
import { v4 as uuidv4 } from 'uuid';
@Injectable()
export class LoggerMiddleware implements NestMiddleware<Request, Response> {
  public use(req: Request, _res: Response, next: () => void): any {
    const before = Date.now();
    const id = req.headers['x-request-id'] ? req.headers['x-request-id'] : uuidv4();
    const span = req.headers['x-span'] || '0';
    req.correlationId = id;
    req.requestId = id;
    req.parentSpan = span;
    req.span = span;
    req.before = before;
    next();
  }
}
