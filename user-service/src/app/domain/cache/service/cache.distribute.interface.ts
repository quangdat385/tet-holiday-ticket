export interface DistributeCacheInterface {
  setString(key: string, value: string): Promise<string>;
  setNumber(key: string, value: number): Promise<string>;
  setObject(key: string, value: object): Promise<number>;
  setStringWithTtl(key: string, value: string, ttl: number): Promise<string>;
  setNumberWithTtl(key: string, value: number, ttl: number): Promise<string>;
  setObjectWithTtl(key: string, field: string, value: object, ttl: number): Promise<number>;
  getString(key: string): Promise<string | null>;
  getNumber(key: string): Promise<number | null>;
  getObject(key: string, field: string): Promise<object | null>;
  getAllObject(key: string): Promise<object | null>;
  del(key: string): Promise<void>;
  delObject(key: string, field: string): Promise<void>;
}
