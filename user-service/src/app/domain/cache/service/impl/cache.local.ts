import { CACHE_MANAGER } from '@nestjs/cache-manager';
import { Inject, Injectable } from '@nestjs/common';
import { Cache } from 'cache-manager';
import { LocalCacheInterface } from 'src/app/domain/cache/service/cache.local.interface';

@Injectable()
export class LocalCacheService implements LocalCacheInterface {
  // eslint-disable-next-line prettier/prettier
  constructor(@Inject(CACHE_MANAGER) private cacheManager: Cache) { }
  async setString(key: string, value: string, ttl?: number): Promise<string> {
    return await this.cacheManager.set(key, value, ttl);
  }
  async setNumber(key: string, value: number, ttl?: number): Promise<number> {
    return await this.cacheManager.set(key, value, ttl);
  }
  async setObject(key: string, value: object, ttl?: number): Promise<object> {
    return await this.cacheManager.set(key, value, ttl);
  }
  async getString(key: string): Promise<string> {
    return this.cacheManager.get(key) as Promise<string>;
  }
  async getNumber(key: string): Promise<number> {
    return this.cacheManager.get(key) as Promise<number>;
  }
  async getObject(key: string): Promise<object> {
    return (await this.cacheManager.get(key)) as Promise<object>;
  }
  async del(key: string): Promise<boolean> {
    return await this.cacheManager
      .del(key)
      .then(() => true)
      .catch(() => false);
  }
}
