export interface CreateUserUserRole {
  user_id: bigint;
  role_id: bigint;
}

export const createPrismaUserUserRole = (createUserUserRole: CreateUserUserRole) => {
  return {
    ...createUserUserRole
  };
};
