<template>
  <div class="login-container">
    <h2>Login</h2>
    <form @submit.prevent="handleLogin">
      <div>
        <label>Email</label>
        <input type="email" v-model="email" required />
      </div>
      <div>
        <label>Password</label>
        <input type="password" v-model="password" required />
      </div>
      <button type="submit">Login</button>
    </form>
    <p>Don't have an account? <router-link to="/register">Register</router-link></p>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'

const email = ref('')
const password = ref('')
const router = useRouter()

const handleLogin = () => {
  // Simple simulation – replace with actual API call
  if (email.value && password.value) {
    // On successful login, store a token in localStorage
    localStorage.setItem('crm_token', 'dummy-jwt-token')
    alert('Login successful!')
    router.push('/dashboard')
  } else {
    alert('Please enter email and password')
  }
}
</script>

<style scoped>
.login-container { max-width: 400px; margin: 50px auto; padding: 20px; border: 1px solid #ccc; border-radius: 8px; }
input, button { width: 100%; padding: 8px; margin: 8px 0; }
button { background-color: #42b983; color: white; border: none; cursor: pointer; }
</style>