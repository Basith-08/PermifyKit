import { computed } from 'vue';
import { useUserStore } from '../state/user.state';
import { userApi } from '../../../core/api/user.api';

export function useUserViewModel() {
  const store = useUserStore();

  const fetchUsers = async () => {
    store.setLoading(true);
    store.setError(null);
    try {
      const response = await userApi.getAll();
      store.setUsers(response.data);
    } catch (err: any) {
      store.setError(err.response?.data?.error || 'Failed to fetch users');
    } finally {
      store.setLoading(false);
    }
  };

  const createUser = async (data: any) => {
    try {
      await userApi.create(data);
      await fetchUsers();
    } catch (err: any) {
      const msg = err.response?.data?.error || 'Failed to create user';
      store.setError(msg);
      throw new Error(msg);
    }
  };

  const updateUser = async (id: number, data: any) => {
    try {
      await userApi.update(id, data);
      await fetchUsers();
    } catch (err: any) {
      const msg = err.response?.data?.error || 'Failed to update user';
      store.setError(msg);
      throw new Error(msg);
    }
  };

  const deleteUser = async (id: number) => {
    try {
      await userApi.delete(id);
      await fetchUsers(); // Re-fetch after deletion
    } catch (err: any) {
      store.setError(err.response?.data?.error || 'Failed to delete user');
      throw err;
    }
  };

  return {
    users: computed(() => store.users),
    isLoading: computed(() => store.isLoading),
    error: computed(() => store.error),
    fetchUsers,
    createUser,
    updateUser,
    deleteUser
  };
}
