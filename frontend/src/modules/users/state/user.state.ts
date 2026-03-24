import { defineStore } from 'pinia';
import { User } from '../../../core/api/user.api';

interface UserState {
  users: User[];
  isLoading: boolean;
  error: string | null;
}

export const useUserStore = defineStore('user', {
  state: (): UserState => ({
    users: [],
    isLoading: false,
    error: null,
  }),
  actions: {
    setUsers(users: User[]) {
      this.users = users;
    },
    setLoading(status: boolean) {
      this.isLoading = status;
    },
    setError(error: string | null) {
      this.error = error;
    }
  }
});
