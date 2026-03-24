import { computed } from 'vue';
import { useRoleStore } from '../state/role.state';
import { roleApi } from '../../../core/api/role.api';

export function useRoleViewModel() {
  const store = useRoleStore();

  const fetchRoles = async () => {
    store.setLoading(true);
    store.setError(null);
    try {
      const response = await roleApi.getAll();
      store.setRoles(response.data);
    } catch (err: any) {
      store.setError(err.response?.data?.error || 'Failed to fetch roles');
    } finally {
      store.setLoading(false);
    }
  };

  const createRole = async (data: any) => {
    try {
      await roleApi.create(data);
      await fetchRoles();
    } catch (err: any) {
      const msg = err.response?.data?.error || 'Failed to create role';
      store.setError(msg);
      throw new Error(msg);
    }
  };

  const updateRole = async (id: number, data: any) => {
    try {
      await roleApi.update(id, data);
      await fetchRoles();
    } catch (err: any) {
      const msg = err.response?.data?.error || 'Failed to update role';
      store.setError(msg);
      throw new Error(msg);
    }
  };

  const deleteRole = async (id: number) => {
    try {
      await roleApi.delete(id);
      await fetchRoles();
    } catch (err: any) {
      store.setError(err.response?.data?.error || 'Failed to delete role');
      throw err;
    }
  };

  return {
    roles: computed(() => store.roles),
    isLoading: computed(() => store.isLoading),
    error: computed(() => store.error),
    fetchRoles,
    createRole,
    updateRole,
    deleteRole
  };
}
