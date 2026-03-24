import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router';
import { useAuthStore } from '../../modules/auth/state/auth.state';
import { api } from '../api';

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    redirect: '/login'
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import('../../modules/auth/view/LoginView.vue'),
    meta: { guestOnly: true }
  },
  {
    path: '/dashboard',
    component: () => import('../../modules/dashboard/view/DashboardView.vue'),
    meta: { requiresAuth: true },
    children: [
      {
        path: '',
        name: 'DashboardOverview',
        component: () => import('../../modules/dashboard/view/DashboardOverview.vue'),
      },
      {
        path: 'users',
        name: 'DashboardUsers',
        component: () => import('../../modules/users/view/UserListView.vue'),
      },
      {
        path: 'roles',
        name: 'DashboardRoles',
        component: () => import('../../modules/roles/view/RoleListView.vue'),
      },
      {
        path: 'audit',
        name: 'DashboardAudit',
        component: () => import('../../modules/dashboard/view/AuditLogsView.vue'),
      },
      {
        path: 'settings',
        name: 'DashboardSettings',
        component: () => import('../../modules/dashboard/view/SettingsView.vue'),
      }
    ]
  }
];

export const router = createRouter({
  history: createWebHistory(),
  routes
});

router.beforeEach(async (to, _, next) => {
  const store = useAuthStore();
  const isAuthenticated = store.isAuthenticated;

  if (to.meta.requiresAuth && !isAuthenticated) {
    next('/login');
    return;
  }

  // Session Hydration: if we have a token but lost Pinia state (e.g. page refresh)
  if (isAuthenticated && !store.user) {
    try {
      const [userRes, permRes] = await Promise.all([
        api.get('/user/me'),
        api.get('/user/permissions')
      ]);
      store.setUser(userRes.data);
      store.setPermissions(permRes.data.permissions);
    } catch (err) {
      store.logout();
      next('/login');
      return;
    }
  }

  if (to.meta.guestOnly && isAuthenticated) {
    next('/dashboard');
  } else {
    next();
  }
});
