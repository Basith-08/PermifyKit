import { api } from './index';

export interface Permission {
  id: number;
  resource: string;
  action: string;
  description?: string;
}

export const permissionApi = {
  getAll: () => api.get<Permission[]>('/permissions'),
};
