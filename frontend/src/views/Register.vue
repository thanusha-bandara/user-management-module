<template>
  <div class="auth-wrapper">
    <div class="auth-card">
      <h2>Create an Account</h2>
      <p class="subtitle">Join the CRM platform today.</p>
      
      <form @submit.prevent="handleRegister">
        <div class="form-group">
          <label>Full Name</label>
          <input v-model="form.full_name" type="text" placeholder="John Doe" required />
        </div>

        <div class="form-group">
          <label>Username</label>
          <input v-model="form.username" type="text" placeholder="johndoe123" required />
        </div>

        <div class="form-group">
          <label>Email Address</label>
          <input v-model="form.email" type="email" placeholder="example@mail.com" required />
        </div>
        
        <div class="form-group">
          <label>Password</label>
          <input v-model="form.password" type="password" placeholder="min 6 characters" required />
        </div>


        <button type="submit" :disabled="loading" class="btn-auth">
          {{ loading ? 'Creating...' : 'Register User' }}
        </button>
      </form>
      <p class="switch-text">Already have an account? <router-link to="/login">Sign In</router-link></p>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import axios from 'axios'

const router = useRouter()
const loading = ref(false)
const form = ref({ full_name: '', username: '', email: '', password: '', role_id: 4 })

const handleRegister = async () => {
  loading.value = true
  try {
    form.value.role_id = parseInt(form.value.role_id)
    const res = await axios.post('http://localhost:8080/api/v1/auth/register', form.value)
    alert(res.data.message)
    router.push('/login')
  } catch (err) {
    alert(err.response?.data?.error || 'Registration Failed.')
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
.auth-wrapper { display: flex; justify-content: center; align-items: center; min-height: 100vh; padding: 40px 20px; box-sizing: border-box; background: transparent; font-family: 'Inter', sans-serif; }
.auth-card { background: rgba(30, 41, 59, 0.7); backdrop-filter: blur(12px); padding: 40px; border-radius: 16px; border: 1px solid rgba(255, 255, 255, 0.1); box-shadow: 0 10px 30px rgba(0,0,0,0.3); width: 100%; max-width: 420px; transition: transform 0.3s ease; }
.auth-card:hover { transform: translateY(-5px); }
h2 { margin-bottom: 8px; color: #f8fafc; text-align: center; font-weight: 700; font-size: 26px; }
.subtitle { text-align: center; color: #94a3b8; font-size: 14px; margin-bottom: 30px; }
.form-group { margin-bottom: 20px; }
label { display: block; margin-bottom: 8px; font-size: 12px; color: #cbd5e1; font-weight: 600; text-transform: uppercase; letter-spacing: 0.5px; }
input { width: 100%; padding: 14px 16px; background: rgba(15, 23, 42, 0.6); color: #f8fafc; border: 1px solid rgba(255, 255, 255, 0.08); border-radius: 8px; box-sizing: border-box; font-size: 15px; transition: all 0.3s ease; outline: none; }
input:focus { border-color: #6366f1; box-shadow: 0 0 0 3px rgba(99, 102, 241, 0.2); }
input::placeholder { color: #475569; }
.btn-auth { width: 100%; padding: 15px; background: linear-gradient(to right, #6366f1, #8b5cf6); color: white; border: none; border-radius: 8px; font-weight: 600; font-size: 16px; cursor: pointer; margin-top: 10px; transition: all 0.3s ease; box-shadow: 0 4px 12px rgba(99, 102, 241, 0.3); }
.btn-auth:hover { background: linear-gradient(to right, #4f46e5, #7c3aed); box-shadow: 0 6px 15px rgba(99, 102, 241, 0.4); transform: translateY(-1px); }
.btn-auth:disabled { opacity: 0.7; cursor: not-allowed; transform: none; box-shadow: none; }
.switch-text { text-align: center; font-size: 14px; margin-top: 25px; color: #94a3b8; }
.switch-text a { color: #818cf8; text-decoration: none; font-weight: 600; transition: color 0.2s ease; }
.switch-text a:hover { color: #a5b4fc; text-decoration: underline; }
</style>