<script setup lang="ts">
import { computed } from 'vue';
import { useAuthViewModel } from '../../auth/viewmodel/useAuthViewModel';
import { usePermissions } from '../../../core/permissions/usePermissions';
import Button from 'primevue/button';

const { user } = useAuthViewModel();
const { hasPermission } = usePermissions();
const currentUser = computed(() => user());
</script>

<template>
  <div class="animate-fade-in">
    <!-- Welcome Banner -->
    <div class="bg-slate-900 rounded-md p-10 text-white shadow-2xl shadow-slate-900/20 mb-12 relative overflow-hidden flex items-center justify-between">
      <div class="relative z-10 max-w-2xl">
        <span class="inline-block py-1 px-3 rounded-full bg-white/10 text-xs font-bold tracking-widest uppercase mb-4 text-indigo-200 border border-white/10">Workspace Status</span>
        <h2 class="text-4xl font-extrabold mb-4 tracking-tight" v-if="currentUser">Welcome back, {{ currentUser.username }}! 👋</h2>
        <p class="text-slate-400 text-lg leading-relaxed font-medium">Your permissions are fully synchronized. You have visibility over your assigned resources according to the strict RBAC policies.</p>
      </div>
      <div class="hidden md:flex relative z-10">
         <div class="w-32 h-32 bg-indigo-500 rounded-full blur-3xl opacity-50 absolute -left-10 top-0"></div>
         <div class="w-32 h-32 bg-purple-500 rounded-full blur-3xl opacity-50 absolute right-0 bottom-0"></div>
         <i class="pi pi-verified text-8xl text-indigo-400 drop-shadow-2xl relative z-20"></i>
      </div>
      <!-- Abstract Background Texture -->
      <div class="absolute inset-0 bg-[url('https://www.transparenttextures.com/patterns/carbon-fibre.png')] opacity-20 Mix-blend-overlay"></div>
    </div>

    <div class="flex items-center justify-between mb-8">
      <h3 class="text-2xl font-extrabold text-slate-800 tracking-tight">Active Modules</h3>
    </div>

    <!-- Metric Cards Grid -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">

      <!-- Users Module Card -->
      <router-link to="/dashboard/users" v-if="hasPermission('users', 'read')" class="group bg-white rounded-md p-8 border border-slate-200 shadow-sm hover:shadow-2xl hover:shadow-indigo-500/10 hover:border-indigo-300 transition-all duration-300 transform hover:-translate-y-1 relative overflow-hidden block">
        <div class="absolute top-0 right-0 p-8 opacity-5 group-hover:opacity-10 transition-opacity transform group-hover:scale-110 duration-500">
           <i class="pi pi-users text-8xl text-indigo-900"></i>
        </div>
        
        <div class="w-14 h-14 bg-indigo-50 border border-indigo-100 text-indigo-600 rounded-lg flex items-center justify-center mb-6 group-hover:bg-indigo-600 group-hover:text-white transition-colors duration-300 relative z-10 shadow-sm">
          <i class="pi pi-users text-2xl"></i>
        </div>
        <h4 class="text-xl font-extrabold text-slate-900 mb-3 relative z-10 tracking-tight">User Management</h4>
        <p class="text-slate-500 mb-8 text-sm font-medium leading-relaxed relative z-10">Manage user identities, oversee accounts, and control system-wide access configurations.</p>
        
        <div class="flex relative z-10">
           <Button v-if="hasPermission('users', 'write')" label="Manage Users" icon="pi pi-arrow-right" iconPos="right" class="font-bold py-3 rounded-md bg-slate-900 border-none hover:bg-indigo-600 text-white shadow-lg w-full justify-center transition-colors" />
        </div>
      </router-link>

      <!-- Roles Module Card -->
      <router-link to="/dashboard/roles" v-if="hasPermission('roles', 'read')" class="group bg-white rounded-md p-8 border border-slate-200 shadow-sm hover:shadow-2xl hover:shadow-purple-500/10 hover:border-purple-300 transition-all duration-300 transform hover:-translate-y-1 relative overflow-hidden block">
         <div class="absolute top-0 right-0 p-8 opacity-5 group-hover:opacity-10 transition-opacity transform group-hover:scale-110 duration-500">
           <i class="pi pi-key text-8xl text-purple-900"></i>
        </div>
        
        <div class="w-14 h-14 bg-purple-50 border border-purple-100 text-purple-600 rounded-lg flex items-center justify-center mb-6 group-hover:bg-purple-600 group-hover:text-white transition-colors duration-300 relative z-10 shadow-sm">
          <i class="pi pi-key text-2xl"></i>
        </div>
        <h4 class="text-xl font-extrabold text-slate-900 mb-3 relative z-10 tracking-tight">Role Configurations</h4>
        <p class="text-slate-500 mb-8 text-sm font-medium leading-relaxed relative z-10">Define custom roles, aggregate permissions via composite policies, and assign globally.</p>
        
        <div class="flex relative z-10">
           <Button v-if="hasPermission('roles', 'write')" label="Configure Roles" severity="secondary" icon="pi pi-arrow-right" iconPos="right" class="bg-white border-2 border-slate-200 text-slate-800 hover:border-slate-300 hover:bg-slate-50 font-bold py-3 rounded-md w-full justify-center shadow-sm" />
        </div>
      </router-link>

      <!-- Analytics (Functional) -->
      <router-link to="/dashboard/audit" class="group bg-white rounded-md p-8 border border-slate-200 shadow-sm hover:shadow-2xl hover:shadow-cyan-500/10 hover:border-cyan-300 transition-all duration-300 transform hover:-translate-y-1 relative overflow-hidden block">
        <div class="absolute top-0 right-0 p-8 opacity-5 group-hover:opacity-10 transition-opacity transform group-hover:scale-110 duration-500">
           <i class="pi pi-chart-bar text-8xl text-cyan-900"></i>
        </div>
        
        <div class="w-14 h-14 bg-cyan-50 border border-cyan-100 text-cyan-600 rounded-lg flex items-center justify-center mb-6 group-hover:bg-cyan-600 group-hover:text-white transition-colors duration-300 relative z-10 shadow-sm">
          <i class="pi pi-chart-bar text-2xl"></i>
        </div>
        <h4 class="text-xl font-extrabold text-slate-900 mb-3 relative z-10 tracking-tight">Audit Logs & Insights</h4>
        <p class="text-slate-500 mb-8 text-sm font-medium leading-relaxed relative z-10">Monitor real-time role delegations, access attempts, and strict compliance logs.</p>
        
        <div class="flex relative z-10">
           <Button label="View System Activity" icon="pi pi-arrow-right" iconPos="right" class="font-bold py-3 rounded-md bg-slate-100 border-none text-slate-700 hover:bg-cyan-50 w-full justify-center transition-colors" />
        </div>
      </router-link>
      
    </div>
  </div>
</template>

<style scoped>
.animate-fade-in {
  animation: fadeIn 0.4s ease-out;
}
@keyframes fadeIn {
  from { opacity: 0; transform: translateY(10px); }
  to { opacity: 1; transform: translateY(0); }
}
</style>
