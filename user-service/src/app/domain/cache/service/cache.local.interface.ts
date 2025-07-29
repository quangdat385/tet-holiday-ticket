export interface LocalCacheInterface {
  setString(key: string, value: string, ttl?: number): Promise<string>;
  setNumber(key: string, value: number, ttl?: number): Promise<number>;
  setObject(key: string, value: object, ttl?: number): Promise<object>;
  getString(key: string): Promise<string>;
  getNumber(key: string): Promise<number>;
  getObject(key: string): Promise<object>;
  del(key: string): Promise<boolean>;
}
