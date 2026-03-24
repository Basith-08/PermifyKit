import { defineStore } from 'pinia';
import { Role } from '../../../core/api/role.api';

interface RoleState {
  roles: Role[];
  isLoading: boolean;
  error: string | null;
}

export const useRoleStore = defineStore('role', {
  state: (): RoleState => ({
    roles: [],
    isLoading: false,
    error: null,
  }),
  actions: {
    setRoles(roles: Role[]) {
      this.roles = roles;
    },
    setLoading(status: boolean) {
      this.isLoading = status;
    },
    setError(error: string | null) {
      this.error = error;
    }
  }
});
