<template>
  <div class="register-container">
    <h2>Register</h2>
    <form @submit.prevent="handleRegister">
      <div>
        <label>Name</label>
        <input type="text" v-model="name" required />
      </div>
      <div>
        <label>Email</label>
        <input type="email" v-model="email" required />
      </div>
      <div>
        <label>Password</label>
        <input type="password" v-model="password" required />
      </div>
      <button type="submit">Register</button>
    </form>
    <p>Already have an account? <router-link to="/login">Login</router-link></p>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'

const name = ref('')
const email = ref('')
const password = ref('')
const router = useRouter()

const handleRegister = () => {
  // Simulate API call
  if (name.value && email.value && password.value) {
    // On successful registration, automatically log the user in
    localStorage.setItem('crm_token', 'dummy-jwt-token')
    alert('Registration successful! You are now logged in.')
    router.push('/dashboard')
  } else {
    alert('Please fill in all fields')
  }
}
</script>

<style scoped>
.register-container { max-width: 400px; margin: 50px auto; padding: 20px; border: 1px solid #ccc; border-radius: 8px; }
input, button { width: 100%; padding: 8px; margin: 8px 0; }
button { background-color: #42b983; color: white; border: none; cursor: pointer; }
</style>