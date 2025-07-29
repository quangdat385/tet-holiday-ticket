import { Module } from '@nestjs/common';
import { DatabaseService } from './database.service';

@Module({
  providers: [DatabaseService],
  exports: [DatabaseService]
})
// eslint-disable-next-line prettier/prettier
export class DatabaseModule { }
