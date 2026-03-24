<script setup lang="ts">
import { computed } from 'vue';
import { useAuthViewModel } from '../../auth/viewmodel/useAuthViewModel';
import Button from 'primevue/button';
import InputText from 'primevue/inputtext';
import Divider from 'primevue/divider';
import 'primeicons/primeicons.css';

const { user } = useAuthViewModel();
const currentUser = computed(() => user());

// removed unused settings ref
// const settings = ref({ ... });

const languages = ['English', 'Indonesian', 'Japanese'];
// removed unused timezones
</script>

<template>
  <div class="space-y-6 animate-fade-in max-w-4xl">
    <div class="bg-white p-6 rounded-md shadow-sm border border-slate-200">
      <h2 class="text-2xl font-extrabold text-slate-900 tracking-tight">Account Settings</h2>
      <p class="text-slate-500 font-medium mt-1">Manage your personal preferences and security configurations.</p>
    </div>

    <!-- Profile Section -->
    <div class="bg-white rounded-md border border-slate-200 shadow-sm p-8">
      <div class="flex items-center gap-6 mb-8">
        <div class="w-20 h-20 rounded-full bg-indigo-100 text-indigo-700 flex items-center justify-center border-4 border-white shadow-md">
           <i class="pi pi-user text-4xl"></i>
        </div>
        <div>
           <h3 class="text-xl font-bold text-slate-900">Personal Profile</h3>
           <p class="text-slate-500 text-sm">Update your public information and avatar.</p>
        </div>
      </div>

      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div class="flex flex-col gap-2">
          <label class="text-sm font-bold text-slate-700">Display Name</label>
          <InputText :value="currentUser?.username" disabled class="p-3 bg-slate-50 border-slate-200 text-slate-500" />
        </div>
        <div class="flex flex-col gap-2">
          <label class="text-sm font-bold text-slate-700">Email Address</label>
          <InputText :value="currentUser?.email" disabled class="p-3 bg-slate-50 border-slate-200 text-slate-500" />
        </div>
      </div>
      
      <Divider />
      
      <div class="flex flex-col gap-4 mt-6">
        <h4 class="font-bold text-slate-900">Login Security</h4>
        <div class="flex items-center justify-between p-4 bg-slate-50 rounded-md border border-slate-200">
          <div class="flex items-center gap-3">
            <i class="pi pi-lock text-slate-400"></i>
             <div>
               <p class="text-sm font-bold text-slate-700">Password</p>
               <p class="text-xs text-slate-500">Last changed 2 months ago</p>
             </div>
          </div>
          <Button label="Change" text class="font-bold text-indigo-600" />
        </div>
      </div>
    </div>

    <!-- Preferences -->
    <div class="bg-white rounded-md border border-slate-200 shadow-sm p-8">
      <h3 class="text-xl font-bold text-slate-900 mb-6">System Preferences</h3>
      
      <div class="space-y-6">
        <div class="flex items-center justify-between">
          <div>
            <p class="font-bold text-slate-700">Interface Language</p>
            <p class="text-xs text-slate-500">Pick which language you want to see.</p>
          </div>
          <select class="p-2 border border-slate-200 rounded-md bg-white text-sm font-medium w-40 outline-none focus:border-indigo-500">
            <option v-for="lang in languages" :key="lang" :value="lang">{{ lang }}</option>
          </select>
        </div>

        <div class="flex items-center justify-between">
          <div>
            <p class="font-bold text-slate-700">Email Notifications</p>
            <p class="text-xs text-slate-500">Receive weekly security digest emails.</p>
          </div>
          <div class="w-12 h-6 bg-indigo-600 rounded-full cursor-pointer relative transition-colors">
            <div class="absolute right-1 top-1 w-4 h-4 bg-white rounded-full shadow-sm"></div>
          </div>
        </div>
      </div>

      <div class="mt-10 flex justify-end">
         <Button label="Save Preferences" icon="pi pi-check" class="bg-indigo-600 hover:bg-indigo-700 border-none font-bold py-3 px-8 shadow-md shadow-indigo-600/20" />
      </div>
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
