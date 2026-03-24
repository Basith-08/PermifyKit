import { api } from './index';

export interface AuditLog {
  id: number;
  user_id: number;
  username: string;
  action: string;
  resource: string;
  resource_id: string;
  timestamp: string;
  ip_address: string;
}

export const auditApi = {
  getLogs: () => api.get<AuditLog[]>('/audit'),
};
