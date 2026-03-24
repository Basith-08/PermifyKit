<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { useRoleViewModel } from '../viewmodel/useRoleViewModel';
import { usePermissions } from '../../../core/permissions/usePermissions';
import { permissionApi, Permission } from '../../../core/api/permission.api';
import Button from 'primevue/button';
import Tag from 'primevue/tag';
import ProgressSpinner from 'primevue/progressspinner';
import Dialog from 'primevue/dialog';
import InputText from 'primevue/inputtext';
import Textarea from 'primevue/textarea';
import MultiSelect from 'primevue/multiselect';
import 'primeicons/primeicons.css';

const { roles, isLoading, error, fetchRoles, deleteRole, createRole, updateRole } = useRoleViewModel();
const { hasPermission } = usePermissions();

const availablePermissions = ref<Permission[]>([]);
const showRoleModal = ref(false);
const isEditing = ref(false);
const isSaving = ref(false);
const roleForm = ref({ id: 0, name: '', description: '', permission_ids: [] as number[] });

onMounted(async () => {
  fetchRoles();
  if (hasPermission('roles', 'write')) {
    try {
      const { data } = await permissionApi.getAll();
      availablePermissions.value = data;
    } catch (err) {
      console.error('Failed to load permissions catalog', err);
    }
  }
});

const openNew = () => {
  roleForm.value = { id: 0, name: '', description: '', permission_ids: [] };
  isEditing.value = false;
  showRoleModal.value = true;
};

const openEdit = (role: any) => {
  roleForm.value = { 
    id: role.id, 
    name: role.name, 
    description: role.description, 
    permission_ids: role.permissions?.map((p: any) => p.id) || [] 
  };
  isEditing.value = true;
  showRoleModal.value = true;
};

const saveRole = async () => {
  isSaving.value = true;
  try {
    if (isEditing.value) {
      await updateRole(roleForm.value.id, roleForm.value);
    } else {
      await createRole(roleForm.value);
    }
    showRoleModal.value = false;
  } finally {
    isSaving.value = false;
  }
};

const handleDelete = async (id: number) => {
  if (confirm('Are you sure you want to delete this role?')) {
    await deleteRole(id);
  }
};
</script>

<template>
  <div class="space-y-6 animate-fade-in">
    <!-- Header Section -->
    <div class="flex flex-col md:flex-row justify-between items-start md:items-center gap-4 bg-white p-6 rounded-md shadow-sm border border-slate-200">
      <div>
        <h2 class="text-2xl font-extrabold text-slate-900 tracking-tight">Access Roles</h2>
        <p class="text-slate-500 font-medium mt-1">Configure composite permissions and control policies.</p>
      </div>
      <Button v-if="hasPermission('roles', 'write')" @click="openNew" label="Create Role" icon="pi pi-plus" class="bg-purple-600 hover:bg-purple-700 border-none rounded-md font-bold py-3 shadow-md shadow-purple-600/20 transition-all" />
    </div>

    <!-- Error State -->
    <div v-if="error" class="bg-red-50 text-red-700 p-4 rounded-md border border-red-200 flex items-center font-semibold">
      <i class="pi pi-exclamation-circle mr-3 text-xl"></i>
      {{ error }}
    </div>

    <!-- Loader -->
    <div v-if="isLoading && roles.length === 0" class="flex justify-center items-center py-20">
       <ProgressSpinner strokeWidth="4" class="w-12 h-12" />
    </div>

    <!-- Roles Grid View -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6" v-if="!isLoading || roles.length > 0">
      
      <!-- Role Card -->
      <div v-for="role in roles" :key="role.id" class="bg-white rounded-md p-6 border border-slate-200 shadow-sm relative overflow-hidden group hover:shadow-2xl hover:shadow-purple-500/10 hover:border-purple-300 transition-all duration-300">
         <div class="absolute top-0 right-0 p-6 opacity-5 group-hover:opacity-10 transition-opacity transform group-hover:scale-110 duration-500">
           <i class="pi pi-shield text-6xl text-purple-900"></i>
        </div>

        <div class="flex justify-between items-start mb-4 relative z-10">
          <div class="flex items-center gap-3">
             <div class="w-12 h-12 rounded-lg bg-gradient-to-br from-purple-100 to-fuchsia-100 text-purple-700 flex items-center justify-center border border-purple-200 shadow-inner">
               <i class="pi pi-shield text-xl"></i>
             </div>
             <div>
               <h3 class="text-lg font-extrabold text-slate-900 capitalize">{{ role.name }}</h3>
               <span class="text-xs font-bold text-slate-400">ID: {{ role.id }}</span>
             </div>
          </div>
          
          <div class="flex gap-1" v-if="hasPermission('roles', 'write')">
             <Button icon="pi pi-pencil" text rounded class="text-purple-600 hover:bg-purple-50 w-8 h-8 p-0" @click="openEdit(role)" />
             <Button icon="pi pi-trash" text rounded class="text-red-500 hover:bg-red-50 w-8 h-8 p-0" @click="handleDelete(role.id)" />
          </div>
        </div>

        <p class="text-slate-500 text-sm font-medium mb-6 relative z-10 min-h-[40px]">
          {{ role.description || 'No description provided for this role.' }}
        </p>

        <div class="relative z-10">
          <p class="text-xs font-bold text-slate-400 uppercase tracking-widest mb-3">Assigned Permissions</p>
          <div class="flex flex-wrap gap-2">
            <Tag v-for="perm in role.permissions" :key="perm.id" :value="`${perm.resource}.${perm.action}`" 
                 class="font-extrabold text-[10px] px-2 py-1 tracking-wider bg-slate-100 text-slate-700 border border-slate-200" rounded />
            <span v-if="!role.permissions || role.permissions.length === 0" class="text-slate-400 italic text-sm">No permissions configured</span>
          </div>
        </div>
      </div>
      
    </div>

    <!-- Role Modal -->
    <Dialog v-model:visible="showRoleModal" :modal="true" :header="isEditing ? 'Edit Role' : 'Create New Role'" :style="{ width: '450px' }" class="p-fluid">
      <form @submit.prevent="saveRole" class="space-y-4 mt-2">
        <div class="flex flex-col gap-2">
          <label for="name" class="text-sm font-bold text-slate-700">Role Name</label>
          <InputText id="name" v-model.trim="roleForm.name" required autofocus class="p-3 border rounded-md focus:border-purple-500" placeholder="e.g. manager" />
        </div>
        
        <div class="flex flex-col gap-2">
          <label for="description" class="text-sm font-bold text-slate-700">Description</label>
          <Textarea id="description" v-model.trim="roleForm.description" rows="3" class="p-3 border rounded-md focus:border-purple-500" placeholder="Brief explanation of this role's purpose..." />
        </div>

        <div class="flex flex-col gap-2">
          <label for="permissions" class="text-sm font-bold text-slate-700">Assign Permissions <span class="text-xs font-normal text-slate-400">({{availablePermissions.length}} available)</span></label>
          <MultiSelect id="permissions" v-model="roleForm.permission_ids" :options="availablePermissions" optionLabel="description" optionValue="id" placeholder="Select System Permissions"
            :maxSelectedLabels="2" filter class="border rounded-md focus:border-purple-500">
            <template #option="slotProps">
              <div class="flex flex-col">
                <span class="font-bold text-slate-800">{{ slotProps.option.description }}</span>
                <span class="text-xs text-slate-500 font-mono mt-0.5">{{ slotProps.option.resource }}.{{ slotProps.option.action }}</span>
              </div>
            </template>
          </MultiSelect>
        </div>

        <div class="pt-6 flex justify-end gap-3 border-t border-slate-100 mt-6">
          <Button label="Cancel" icon="pi pi-times" text class="text-slate-500 hover:bg-slate-100 font-bold" @click="showRoleModal = false"/>
          <Button type="submit" :label="isEditing ? 'Save Changes' : 'Create Role'" icon="pi pi-check" :loading="isSaving" class="bg-purple-600 hover:bg-purple-700 border-none font-bold shadow-md shadow-purple-500/20" />
        </div>
      </form>
    </Dialog>
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
