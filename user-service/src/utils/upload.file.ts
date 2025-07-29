/* eslint-disable @typescript-eslint/no-misused-promises */
import { diskStorage } from 'multer';
import { join } from 'path';
import { existsSync } from 'fs';
import { mkdir } from 'fs/promises';
import { Request } from 'express';

const storage = diskStorage({
  destination: async (req, file, cb) => {
    if (existsSync(join(__dirname, '..', 'public/img/avatar'))) {
      await mkdir(join(__dirname, '..', 'public/img/avatar'));
    }
    cb(null, 'src/public/img/avatar');
  },
  filename: function (req, file, cb) {
    cb(null, 'avatar' + Date.now());
  }
});
const imageFilter = function (
  req: Request,
  file: Express.Multer.File,
  cb: (error: Error | null, acceptFile: boolean) => void
) {
  // Accept images only
  if (!file.originalname.match(/\.(jpg|JPG|jpeg|JPEG|png|PNG|gif|GIF|jfif)$/)) {
    req.fileValidationError = 'Only image files are allowed!';
    return cb(new Error('Only image files are allowed!'), false);
  }
  cb(null, true);
};

export { storage, imageFilter };
