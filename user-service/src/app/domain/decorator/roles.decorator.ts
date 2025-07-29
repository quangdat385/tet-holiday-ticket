import { SetMetadata } from '@nestjs/common';
import { UserRoles } from 'src/app/domain/user/user.type';

export const RoleAllowed = (...role: UserRoles[]) => SetMetadata('roles', role);
