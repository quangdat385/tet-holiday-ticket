import { CreateTokenPayLoad } from 'src/app/domain/auth/dto/request/CreateToken';

declare module 'express' {
  interface Request {
    requestId?: string | string[];
    before: number;
    correlationId: string | string[];
    headers: {
      x_device_id: string | string[];
      cookie: string;
      [key: string]: string | string[] | undefined;
    };
    parentSpan: string | string[];
    span: string | string[];
    origin: string;
    user: {
      [key: string]: any;
    };
    token_id: bigint;
    token_payload?: CreateTokenPayLoad;
    access_token?: string;
    refresh_token?: string;
    exprire_at?: Date;
    cookies: {
      [key: string]: any;
    };
    context?: {
      [key: string]: any;
    };
    fileValidationError?: string;
    device: string;
  }
}
