import { defineStore } from 'pinia';

interface AuthState {
  token: string | null;
  user: any | null;
  permissions: string[];
}

export const useAuthStore = defineStore('auth', {
  state: (): AuthState => ({
    token: localStorage.getItem('token') || null,
    user: null,
    permissions: [],
  }),
  getters: {
    isAuthenticated: (state) => !!state.token,
  },
  actions: {
    setToken(token: string) {
      this.token = token;
      localStorage.setItem('token', token);
    },
    logout() {
      this.token = null;
      this.user = null;
      this.permissions = [];
      localStorage.removeItem('token');
    },
    setUser(user: any) {
      this.user = user;
    },
    setPermissions(perms: string[]) {
      this.permissions = perms;
    }
  }
});
