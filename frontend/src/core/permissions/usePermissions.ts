import { useAuthStore } from '../../modules/auth/state/auth.state';

export function usePermissions() {
  const store = useAuthStore();

  const hasPermission = (resource: string, action: string): boolean => {
    const required = `${resource}.${action}`;
    return store.permissions.includes(required);
  };

  const hasAnyPermission = (requirements: Array<{resource: string, action: string}>): boolean => {
    return requirements.some(req => hasPermission(req.resource, req.action));
  };

  return {
    hasPermission,
    hasAnyPermission
  };
}
