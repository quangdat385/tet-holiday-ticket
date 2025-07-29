import { Module } from '@nestjs/common';
import { DistributeCacheService } from 'src/app/domain/cache/service/impl/cache.distribute';
import { LocalCacheService } from 'src/app/domain/cache/service/impl/cache.local';

@Module({
  providers: [LocalCacheService, DistributeCacheService],
  exports: [LocalCacheService, DistributeCacheService]
})
// eslint-disable-next-line prettier/prettier
export class CacheModule { }
