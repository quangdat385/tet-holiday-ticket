import { Controller, Get } from '@nestjs/common';
import { ApiExcludeEndpoint, ApiOkResponse, ApiTags } from '@nestjs/swagger';
import { HealthCheck, HealthCheckService, HttpHealthIndicator } from '@nestjs/terminus';
import { generateHash } from 'src/utils/crypto';

@Controller('/health')
@ApiTags('health')
export class AppController {
  constructor(
    private readonly health: HealthCheckService,
    private http: HttpHealthIndicator
  ) {
    console.log('AppController initialized');
  }
  @ApiExcludeEndpoint()
  @Get('hello')
  getHello(): string {
    return generateHash('1', '123456');
  }
  @Get('health')
  @ApiOkResponse({ description: 'returns database the health check ' })
  @HealthCheck()
  getApplicationHealth() {
    return this.health.check([
      () => this.http.pingCheck('holiday-ticket', 'http://localhost:8080/ticket-user/api/v1/health/hello')
    ]);
  }
}
