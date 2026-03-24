<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { useUserViewModel } from '../viewmodel/useUserViewModel';
import { useRoleViewModel } from '../../roles/viewmodel/useRoleViewModel';
import { usePermissions } from '../../../core/permissions/usePermissions';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
import Button from 'primevue/button';
import Tag from 'primevue/tag';
import ProgressSpinner from 'primevue/progressspinner';
import Dialog from 'primevue/dialog';
import InputText from 'primevue/inputtext';
import MultiSelect from 'primevue/multiselect';
import 'primeicons/primeicons.css';

const { users, isLoading, error, fetchUsers, deleteUser, createUser, updateUser } = useUserViewModel();
const { roles, fetchRoles } = useRoleViewModel();
const { hasPermission } = usePermissions();

const showUserModal = ref(false);
const isEditing = ref(false);
const isSaving = ref(false);
const userForm = ref({ id: 0, username: '', email: '', password: '', role_ids: [] as number[] });

onMounted(() => {
  fetchUsers();
  if (hasPermission('users', 'write')) {
    fetchRoles();
  }
});

const openNew = () => {
  userForm.value = { id: 0, username: '', email: '', password: '', role_ids: [] };
  isEditing.value = false;
  showUserModal.value = true;
};

const openEdit = (user: any) => {
  userForm.value = { 
    id: user.id, 
    username: user.username, 
    email: user.email, 
    password: '', 
    role_ids: user.roles?.map((r: any) => r.id) || [] 
  };
  isEditing.value = true;
  showUserModal.value = true;
};

const saveUser = async () => {
  isSaving.value = true;
  try {
    if (isEditing.value) {
      await updateUser(userForm.value.id, userForm.value);
    } else {
      await createUser(userForm.value);
    }
    showUserModal.value = false;
  } finally {
    isSaving.value = false;
  }
};

const handleDelete = async (id: number) => {
  if (confirm('Are you sure you want to delete this user?')) {
    await deleteUser(id);
  }
};
</script>

<template>
  <div class="space-y-6 animate-fade-in">
    <!-- Header Section -->
    <div class="flex flex-col md:flex-row justify-between items-start md:items-center gap-4 bg-white p-6 rounded-md shadow-sm border border-slate-200">
      <div>
        <h2 class="text-2xl font-extrabold text-slate-900 tracking-tight">User Directory</h2>
        <p class="text-slate-500 font-medium mt-1">Manage system accounts and access identities.</p>
      </div>
      <Button v-if="hasPermission('users', 'write')" @click="openNew" label="Add New User" icon="pi pi-plus" class="bg-indigo-600 hover:bg-indigo-700 border-none rounded-md font-bold py-3 shadow-md shadow-indigo-600/20 transition-all" />
    </div>

    <!-- Error State -->
    <div v-if="error" class="bg-red-50 text-red-700 p-4 rounded-md border border-red-200 flex items-center font-semibold">
      <i class="pi pi-exclamation-circle mr-3 text-xl"></i>
      {{ error }}
    </div>

    <!-- Data Table Section -->
    <div class="bg-white rounded-md shadow-sm border border-slate-200 overflow-hidden relative">
      <div v-if="isLoading" class="absolute inset-0 bg-white/50 backdrop-blur-sm z-10 flex items-center justify-center">
        <ProgressSpinner strokeWidth="4" class="w-12 h-12" />
      </div>

      <DataTable :value="users" :paginator="true" :rows="10" 
         class="p-datatable-sm w-full"
         pt:headerRow:class="bg-slate-50 text-slate-600 uppercase text-xs tracking-wider"
         pt:bodyRow:class="border-b border-slate-100 hover:bg-slate-50/50 transition-colors"
         responsiveLayout="scroll">
         
        <template #empty>
          <div class="text-center p-8 text-slate-500 font-medium">No users found.</div>
        </template>

        <Column field="username" header="Username" :sortable="true" style="min-width: 15rem">
          <template #body="{ data }">
            <div class="flex items-center gap-3 py-2">
              <div class="w-10 h-10 rounded-full bg-gradient-to-br from-indigo-100 to-purple-100 text-indigo-700 font-bold flex items-center justify-center border border-indigo-200">
                {{ data.username.charAt(0).toUpperCase() }}
              </div>
              <span class="font-bold text-slate-900">{{ data.username }}</span>
            </div>
          </template>
        </Column>

        <Column field="email" header="Email" :sortable="true" style="min-width: 15rem">
          <template #body="{ data }">
            <span class="text-slate-600">{{ data.email }}</span>
          </template>
        </Column>

        <Column header="Roles" style="min-width: 10rem">
          <template #body="{ data }">
            <div class="flex flex-wrap gap-2">
              <Tag v-for="role in data.roles" :key="role.id" :value="role.name" 
                 class="font-extrabold uppercase text-[10px] px-2 py-1 tracking-wider" 
                 :class="role.name === 'admin' ? 'bg-indigo-100 text-indigo-800 border border-indigo-200' : 'bg-slate-100 text-slate-700 border border-slate-200'" rounded />
              <span v-if="!data.roles || data.roles.length === 0" class="text-slate-400 italic text-sm">No roles assigned</span>
            </div>
          </template>
        </Column>

        <Column header="Actions" :exportable="false" style="min-width: 8rem; text-align: center">
          <template #body="{ data }">
            <div class="flex items-center justify-center gap-2" v-if="hasPermission('users', 'write')">
               <Button icon="pi pi-pencil" text rounded class="text-indigo-600 hover:bg-indigo-50 w-10 h-10" @click="openEdit(data)" />
               <Button icon="pi pi-trash" text rounded class="text-red-500 hover:bg-red-50 w-10 h-10" @click="handleDelete(data.id)" />
            </div>
          </template>
        </Column>
      </DataTable>
    </div>

    <!-- User Modal -->
    <Dialog v-model:visible="showUserModal" :modal="true" :header="isEditing ? 'Edit User' : 'Create New User'" :style="{ width: '450px' }" class="p-fluid">
      <form @submit.prevent="saveUser" class="space-y-4 mt-2">
        <div class="flex flex-col gap-2">
          <label for="username" class="text-sm font-bold text-slate-700">Username</label>
          <InputText id="username" v-model.trim="userForm.username" required autofocus class="p-3 border rounded-md focus:border-indigo-500" />
        </div>
        
        <div class="flex flex-col gap-2">
          <label for="email" class="text-sm font-bold text-slate-700">Email Address</label>
          <InputText id="email" type="email" v-model.trim="userForm.email" required class="p-3 border rounded-md focus:border-indigo-500" />
        </div>

        <div class="flex flex-col gap-2">
          <label for="password" class="text-sm font-bold text-slate-700">
            Password <span v-if="isEditing" class="text-xs text-slate-400 font-normal">(Leave blank to keep current)</span>
          </label>
          <InputText id="password" type="password" v-model="userForm.password" :required="!isEditing" class="p-3 border rounded-md focus:border-indigo-500" />
        </div>

        <div class="flex flex-col gap-2">
          <label for="roles" class="text-sm font-bold text-slate-700">Assign Roles</label>
          <MultiSelect id="roles" v-model="userForm.role_ids" :options="roles" optionLabel="name" optionValue="id" placeholder="Select Roles"
            :maxSelectedLabels="3" class="border rounded-md focus:border-indigo-500" />
        </div>

        <div class="pt-6 flex justify-end gap-3 border-t border-slate-100 mt-6">
          <Button label="Cancel" icon="pi pi-times" text class="text-slate-500 hover:bg-slate-100 font-bold" @click="showUserModal = false"/>
          <Button type="submit" :label="isEditing ? 'Save Changes' : 'Create Account'" icon="pi pi-check" :loading="isSaving" class="bg-indigo-600 hover:bg-indigo-700 border-none font-bold shadow-md shadow-indigo-500/20" />
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
