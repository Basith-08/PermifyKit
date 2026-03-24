import { ref } from 'vue';
import { useAuthStore } from '../state/auth.state';
import { api } from '../../../core/api';
import { useRouter } from 'vue-router';

export function useAuthViewModel() {
  const store = useAuthStore();
  const router = useRouter();
  const isLoading = ref(false);
  const error = ref<string | null>(null);

  const login = async (username: string, password: string) => {
    isLoading.value = true;
    error.value = null;
    try {
      const { data } = await api.post('/auth/login', { username, password });
      store.setToken(data.token);
      
      await fetchUserData();      
      router.push('/dashboard');
    } catch (err: any) {
      error.value = err.response?.data?.error || 'Login failed';
    } finally {
      isLoading.value = false;
    }
  };

  const fetchUserData = async () => {
    try {
      const [userRes, permRes] = await Promise.all([
        api.get('/user/me'),
        api.get('/user/permissions')
      ]);
      store.setUser(userRes.data);
      store.setPermissions(permRes.data.permissions);
    } catch (err) {
      console.error('Failed to fetch user data', err);
      store.logout();
      router.push('/login');
      throw err;
    }
  };

  const logout = () => {
    store.logout();
    router.push('/login');
  };

  return {
    login,
    fetchUserData,
    logout,
    isLoading,
    error,
    isAuthenticated: () => store.isAuthenticated,
    user: () => store.user,
  };
}
