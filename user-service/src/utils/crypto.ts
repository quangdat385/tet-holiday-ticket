import { createHmac } from 'crypto';

/**
 * Generates a hash for the given input using the specified algorithm and salt.
 * @param input - The input string to hash.
 * @param salt - The salt to add to the input before hashing.
 * @param algorithm - The hashing algorithm to use (default: 'sha256').
 * @returns The resulting hash as a hexadecimal string.
 */
export function generateHash(input: string, salt: string, algorithm: string = 'sha256'): string {
  const hash = createHmac(algorithm, salt);
  hash.update(input);
  return hash.digest('hex');
}
export function getUserHashKey(hashKey: string, device: string): string {
  return `user:${device}:${hashKey}`;
}
export function getUserBaseKey(hashKey: string): string {
  return `user_base:${hashKey}`;
}
