export interface UserRoleType {
  role_id: bigint;
  role_name: string;
  role_description: string | null;
  role_menus: {
    menu_id: bigint;
    menu_name: string;
    menu_url: string;
    menu_prefix: string;
    menu_pid: string;
  }[];
}

export enum UserRoles {
  ADMIN = 'Admin',
  MANAGER = 'Manager',
  STAFF = 'Staff',
  USER = 'User'
}
