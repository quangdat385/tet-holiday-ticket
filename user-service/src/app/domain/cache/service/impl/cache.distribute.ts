import { InjectRedis } from '@nestjs-modules/ioredis';
import { Injectable } from '@nestjs/common';
import Redis from 'ioredis';
import { DistributeCacheInterface } from 'src/app/domain/cache/service/cache.distribute.interface';

@Injectable()
export class DistributeCacheService implements DistributeCacheInterface {
  // eslint-disable-next-line prettier/prettier
  constructor(@InjectRedis() private readonly redis: Redis) { }
  async setString(key: string, value: string): Promise<string> {
    return this.redis.set(key, value);
  }
  async setNumber(key: string, value: number): Promise<string> {
    return await this.redis.set(key, value);
  }
  async setObject(key: string, value: object): Promise<number> {
    return await this.redis.hset(key, value);
  }
  async setStringWithTtl(key: string, value: string, ttl: number): Promise<string> {
    return await this.redis.set(key, value, 'EX', ttl);
  }
  async setNumberWithTtl(key: string, value: number, ttl: number): Promise<string> {
    return await this.redis.set(key, value, 'EX', ttl);
  }
  async setObjectWithTtl(key: string, field: string, value: object, ttl: number): Promise<number> {
    const data = JSON.stringify(value, (key: string, value: unknown) => {
      if (typeof value === 'bigint') {
        return value.toString();
      }
      return value;
    });
    const result = await this.redis.hset(key, field, data);
    await this.redis.expire(key, ttl);
    return result;
  }
  async getString(key: string): Promise<string | null> {
    return await this.redis.get(key);
  }
  async getNumber(key: string): Promise<number | null> {
    return await this.redis.get(key).then((value: string) => {
      if (value) {
        return Number(value);
      }
      return null;
    });
  }
  async getObject<T>(key: string, field: string): Promise<T | null> {
    const value = await this.redis.hget(key, field);
    if (value) {
      return JSON.parse(value) as T;
    }
    return null;
  }
  async getAllObject(key: string): Promise<object | null> {
    const value = await this.redis.hgetall(key);
    if (value) {
      const result: { [key: string]: object } = {};
      for (const [field, val] of Object.entries(value)) {
        result[field] = JSON.parse(val) as object;
      }
      return result;
    }
    return null;
  }
  async del(key: string): Promise<void> {
    return this.redis.del(key).then(() => {
      return;
    });
  }
  async delObject(key: string, field: string): Promise<void> {
    return this.redis.hdel(key, field).then(() => {
      return;
    });
  }
}
