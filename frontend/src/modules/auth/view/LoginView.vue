<script setup lang="ts">
import { ref } from 'vue';
import { useAuthViewModel } from '../viewmodel/useAuthViewModel';
import Message from 'primevue/message';
import 'primeicons/primeicons.css';

const showPassword = ref(false);
import 'primeicons/primeicons.css';

const username = ref('');
const password = ref('');
const { login, isLoading, error } = useAuthViewModel();

const handleLogin = async () => {
  await login(username.value, password.value);
};
</script>

<template>
  <div class="flex min-h-screen bg-slate-50 font-sans">
    
    <!-- Left: Branding & Graphic -->
    <div class="hidden lg:flex lg:w-1/2 relative bg-gradient-to-br from-indigo-700 via-purple-700 to-black overflow-hidden justify-center items-center">
      <div class="absolute inset-0 bg-[url('https://www.transparenttextures.com/patterns/cubes.png')] opacity-10 mix-blend-overlay"></div>
      
      <div class="relative z-10 text-center text-white px-12 max-w-xl">
        <div class="w-20 h-20 bg-white/10 rounded-md flex items-center justify-center backdrop-blur-xl mx-auto mb-10 border border-white/20 shadow-2xl">
          <i class="pi pi-shield text-4xl text-indigo-100"></i>
        </div>
        <h1 class="text-5xl font-extrabold tracking-tight mb-6">PermifyKit</h1>
        <p class="text-xl text-indigo-100/90 font-light leading-relaxed">
          The ultimate Role-Based Access Control starter kit.<br/>Enterprise-grade security meets stunning aesthetics.
        </p>
      </div>

      <!-- Decorative abstract blobs -->
      <div class="absolute -top-32 -left-32 w-[30rem] h-[30rem] bg-purple-500 rounded-full mix-blend-screen filter blur-[100px] opacity-40 animate-blob"></div>
      <div class="absolute -bottom-32 -right-32 w-[30rem] h-[30rem] bg-indigo-500 rounded-full mix-blend-screen filter blur-[100px] opacity-40 animate-blob animation-delay-2000"></div>
    </div>

    <!-- Right: Login Form -->
    <div class="flex flex-col justify-center w-full lg:w-1/2 p-8 sm:p-16 xl:p-24 bg-white relative">
      <div class="w-full max-w-md mx-auto">
        <div class="lg:hidden w-16 h-16 bg-indigo-600 rounded-lg flex items-center justify-center mb-8 shadow-lg shadow-indigo-200">
          <i class="pi pi-shield text-3xl text-white"></i>
        </div>
        
        <h2 class="text-4xl font-extrabold text-slate-900 mb-2 tracking-tight">Welcome back</h2>
        <p class="text-slate-500 mb-10 text-lg">Please enter your details to access the system.</p>
        
        <Message v-if="error" severity="error" :closable="false" class="mb-6 rounded-md">{{ error }}</Message>
        
        <form @submit.prevent="handleLogin" class="flex flex-col gap-6">
          <div class="flex flex-col gap-2">
            <label for="username" class="text-sm font-bold text-slate-700 ml-1">Username</label>
            <input type="text" id="username" v-model="username" required placeholder="Enter admin or johndoe" class="w-full px-4 py-3.5 border border-slate-200 rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 shadow-sm transition-all text-slate-900 bg-white" />
          </div>

          <div class="flex flex-col gap-2">
            <div class="flex justify-between items-center ml-1">
              <label for="password" class="text-sm font-bold text-slate-700">Password</label>
              <a href="#" class="text-sm font-semibold text-indigo-600 hover:text-indigo-500 transition-colors">Forgot password?</a>
            </div>
            
            <div class="relative w-full">
              <input :type="showPassword ? 'text' : 'password'" id="password" v-model="password" required placeholder="••••••••" 
                class="w-full pl-4 pr-12 py-3.5 border border-slate-200 rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 shadow-sm transition-all text-slate-900 bg-white" />
              <button type="button" @click="showPassword = !showPassword" class="absolute inset-y-0 right-0 pr-4 flex items-center justify-center text-slate-400 hover:text-indigo-600 outline-none transition-colors">
                <i :class="['pi', showPassword ? 'pi-eye-slash' : 'pi-eye', 'text-lg']"></i>
              </button>
            </div>
          </div>

          <button type="submit" :disabled="isLoading" class="w-full mt-4 py-4 rounded-md bg-slate-900 hover:bg-slate-800 disabled:opacity-75 disabled:cursor-not-allowed border-none shadow-xl shadow-slate-900/20 transition-all duration-300 font-bold text-white text-lg flex justify-center items-center group">
            <i v-if="isLoading" class="pi pi-spin pi-spinner mr-2"></i>
            <i v-else class="pi pi-sign-in mr-2 group-hover:translate-x-1 transition-transform"></i>
            Sign In
          </button>

        </form>

        <div class="mt-8 pt-8 border-t border-slate-100">
          <p class="text-center text-sm text-slate-500">
            Need an account? <a href="#" class="font-bold text-indigo-600 hover:text-indigo-500 transition-colors">Contact Administrator</a>
          </p>
        </div>
      </div>
    </div>

  </div>
</template>

<style>
@keyframes blob {
  0% { transform: translate(0px, 0px) scale(1); }
  33% { transform: translate(30px, -50px) scale(1.1); }
  66% { transform: translate(-20px, 20px) scale(0.9); }
  100% { transform: translate(0px, 0px) scale(1); }
}
.animate-blob {
  animation: blob 7s infinite;
}
.animation-delay-2000 {
  animation-delay: 2s;
}
</style>
