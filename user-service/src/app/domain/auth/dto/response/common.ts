import { ApiResponseProperty } from '@nestjs/swagger';

export class BadAuthRequetResponse {
  @ApiResponseProperty({
    example: {
      statusCode: 400,
      timestamp: '2024-05-24T01:29:46.687Z',
      path: '/api/v1/auth/*path',
      response: {
        message: 'bad requet',
        error: 'bad requet',
        statusCode: 400
      }
    }
  })
  public body: object;
}
export class NotFoundResponse {
  @ApiResponseProperty({
    example: {
      statusCode: 404,
      timestamp: '2024-05-24T01:29:46.687Z',
      path: '/api/v1/auth/*path',
      response: {
        message: 'token not found',
        error: 'Not Found',
        statusCode: 404
      }
    }
  })
  public body: object;
}
export class UnauthorizedAuthResponse {
  @ApiResponseProperty({
    example: {
      statusCode: 401,
      timestamp: '2024-05-24T01:40:35.358Z',
      path: '/api/v1/auth/*path',
      response: {
        message: 'Unauthorized',
        statusCode: 401
      }
    }
  })
  public body: object;
}
