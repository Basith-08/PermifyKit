<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { auditApi, AuditLog } from '../../../core/api/audit.api';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
import Tag from 'primevue/tag';
import ProgressSpinner from 'primevue/progressspinner';
import 'primeicons/primeicons.css';

const logs = ref<AuditLog[]>([]);
const isLoading = ref(false);
const error = ref<string | null>(null);

const fetchLogs = async () => {
  isLoading.value = true;
  error.value = null;
  try {
    const { data } = await auditApi.getLogs();
    logs.value = data;
  } catch (err: any) {
    error.value = err.response?.data?.error || 'Failed to load audit logs';
  } finally {
    isLoading.value = false;
  }
};

onMounted(() => {
  fetchLogs();
});

const getActionSeverity = (action: string) => {
  switch (action.toLowerCase()) {
    case 'create': return 'success';
    case 'update': return 'info';
    case 'delete': return 'danger';
    case 'login': return 'warn';
    default: return 'secondary';
  }
};

const formatDate = (dateStr: string) => {
  return new Date(dateStr).toLocaleString();
};
</script>

<template>
  <div class="space-y-6 animate-fade-in">
    <div class="bg-white p-6 rounded-md shadow-sm border border-slate-200">
      <h2 class="text-2xl font-extrabold text-slate-900 tracking-tight">Audit Logs & Insights</h2>
      <p class="text-slate-500 font-medium mt-1">Monitor all system-wide activities and security events.</p>
    </div>

    <div v-if="error" class="bg-red-50 text-red-700 p-4 rounded-md border border-red-200 font-semibold">
      {{ error }}
    </div>

    <div class="bg-white rounded-md border border-slate-200 shadow-sm overflow-hidden">
      <DataTable :value="logs" :loading="isLoading" paginator :rows="10" 
        class="p-datatable-sm overflow-hidden" 
        responsiveLayout="stack" breakpoint="960px">
        
        <template #empty>
          <div class="py-10 text-center text-slate-400 font-medium">No system activities recorded yet.</div>
        </template>
        
        <template #loading>
           <div class="flex justify-center p-10"><ProgressSpinner /></div>
        </template>

        <Column field="timestamp" header="Timestamp" sortable style="width: 20%">
          <template #body="{ data }">{{ formatDate(data.timestamp) }}</template>
        </Column>
        
        <Column field="username" header="Operator" sortable style="width: 15%">
          <template #body="{ data }">
            <span class="font-bold text-slate-700">{{ data.username }}</span>
          </template>
        </Column>

        <Column field="action" header="Action" style="width: 12%">
          <template #body="{ data }">
            <Tag :value="data.action.toUpperCase()" :severity="getActionSeverity(data.action)" class="font-extrabold text-[10px]" />
          </template>
        </Column>

        <Column field="resource" header="Resource" style="width: 15%">
          <template #body="{ data }">
             <span class="text-slate-500 font-mono text-xs capitalize">{{ data.resource }}</span>
          </template>
        </Column>

        <Column field="resource_id" header="Target ID" style="width: 10%"></Column>
        
        <Column field="ip_address" header="IP Address" style="width: 15%"></Column>
      </DataTable>
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
