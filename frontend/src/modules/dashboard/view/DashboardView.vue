<script setup lang="ts">
import { computed } from 'vue';
import { useAuthViewModel } from '../../auth/viewmodel/useAuthViewModel';
import { usePermissions } from '../../../core/permissions/usePermissions';
import Button from 'primevue/button';
import Avatar from 'primevue/avatar';
import 'primeicons/primeicons.css';

const { user, logout } = useAuthViewModel();
const { hasPermission } = usePermissions();

const currentUser = computed(() => user());

const handleLogout = () => {
  logout();
};
</script>

<template>
  <div class="min-h-screen bg-slate-50 font-sans flex text-slate-800 selection:bg-indigo-100 selection:text-indigo-900">
    
    <!-- Sidebar -->
    <aside class="w-72 bg-white border-r border-slate-200 flex flex-col hidden lg:flex fixed h-full shadow-sm z-10 transition-all">
      <div class="h-24 flex items-center px-8 border-b border-slate-100">
        <div class="w-12 h-12 bg-gradient-to-br from-indigo-500 to-purple-600 rounded-lg flex items-center justify-center shadow-lg shadow-indigo-200/50 text-white mr-4">
          <i class="pi pi-shield text-2xl"></i>
        </div>
        <span class="text-2xl font-extrabold tracking-tight text-slate-900">Permify.</span>
      </div>
      
      <nav class="flex-1 px-5 py-8 space-y-2 overflow-y-auto">
        <span class="text-xs font-bold text-slate-400 uppercase tracking-widest px-4 mb-4 block">Main Menu</span>
        
        <router-link to="/dashboard" active-class="bg-indigo-50 text-indigo-700 shadow-sm" class="flex items-center px-4 py-3 text-slate-600 hover:bg-slate-50 hover:text-slate-900 rounded-md font-bold transition-all">
          <i class="pi pi-home mr-3 text-lg opacity-80"></i> Overview
        </router-link>
        <router-link to="/dashboard/users" active-class="bg-indigo-50 text-indigo-700 shadow-sm" v-if="hasPermission('users', 'read')" class="flex items-center px-4 py-3 text-slate-600 hover:bg-slate-50 hover:text-slate-900 rounded-md font-semibold transition-all group">
          <i class="pi pi-users mr-3 text-lg opacity-50 group-hover:opacity-100 group-hover:text-indigo-600 transition-all"></i> Users Directory
        </router-link>
        <router-link to="/dashboard/roles" active-class="bg-indigo-50 text-indigo-700 shadow-sm" v-if="hasPermission('roles', 'read')" class="flex items-center px-4 py-3 text-slate-600 hover:bg-slate-50 hover:text-slate-900 rounded-md font-semibold transition-all group">
          <i class="pi pi-key mr-3 text-lg opacity-50 group-hover:opacity-100 group-hover:text-amber-600 transition-all"></i> Access Roles
        </router-link>
        
        <div class="pt-6 pb-2">
          <span class="text-xs font-bold text-slate-400 uppercase tracking-widest px-4 block">Preferences</span>
        </div>
        <router-link to="/dashboard/settings" active-class="bg-indigo-50 text-indigo-700 shadow-sm" class="flex items-center px-4 py-3 text-slate-600 hover:bg-slate-50 hover:text-slate-900 rounded-md font-semibold transition-all group">
          <i class="pi pi-cog mr-3 text-lg opacity-50 group-hover:opacity-100 transition-all"></i> Settings
        </router-link>
      </nav>
      
      <div class="p-5 border-t border-slate-100 bg-slate-50/50">
        <div class="bg-white border border-slate-200 shadow-sm rounded-lg p-4 flex items-center gap-4 hover:border-indigo-300 transition-colors cursor-pointer">
          <Avatar icon="pi pi-user" class="bg-indigo-100 text-indigo-700 font-bold" shape="circle" size="large" />
          <div class="flex-1 min-w-0">
            <p class="text-sm font-extrabold text-slate-900 truncate" v-if="currentUser">{{ currentUser.username }}</p>
            <p class="text-xs font-medium text-slate-500 truncate" v-if="currentUser">{{ currentUser.email }}</p>
          </div>
          <i class="pi pi-ellipsis-v text-slate-400"></i>
        </div>
      </div>
    </aside>

    <!-- Main Content -->
    <div class="flex-1 flex flex-col lg:ml-72 bg-slate-50/50 min-h-screen">
      
      <!-- Top Header -->
      <header class="h-24 bg-white/80 backdrop-blur-md border-b border-slate-200 flex items-center justify-between px-10 sticky top-0 z-20">
        <div class="flex items-center">
          <!-- Mobile Menu Button -->
          <button class="lg:hidden mr-4 text-slate-500 hover:text-slate-800">
            <i class="pi pi-bars text-2xl"></i>
          </button>
          <h1 class="text-2xl font-extrabold text-slate-800 tracking-tight">System Dashboard</h1>
        </div>
        
        <div class="flex items-center gap-5">
          <button class="relative p-2 text-slate-400 hover:text-indigo-600 transition-colors">
            <i class="pi pi-bell text-xl"></i>
            <span class="absolute top-1 right-1 w-2.5 h-2.5 bg-red-500 border-2 border-white rounded-full"></span>
          </button>
          <div class="w-px h-8 bg-slate-200"></div>
          <Button @click="handleLogout" icon="pi pi-sign-out" label="Sign Out" severity="danger" text class="font-bold border border-transparent hover:bg-red-50 hover:border-red-100 rounded-md px-4" />
        </div>
      </header>

      <!-- Dashboard Nested View Routing -->
      <main class="flex-1 p-10 max-w-7xl mx-auto w-full">
        <router-view v-slot="{ Component }">
          <transition name="fade" mode="out-in">
            <component :is="Component" />
          </transition>
        </router-view>
      </main>
    </div>
  </div>
</template>

<style scoped>
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.2s ease, transform 0.2s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
  transform: translateY(5px);
}
</style>
