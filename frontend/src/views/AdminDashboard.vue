<template>
  <div class="dashboard-container">
    <h1>Admin Dashboard</h1>
    <p>Welcome, {{ userName }}!</p>
    <button @click="logout">Logout</button>
    <!-- Add your dashboard content here -->
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()
const userName = ref('User')

onMounted(() => {
  // Check token (router guard already does this, but extra safety)
  const token = localStorage.getItem('crm_token')
  if (!token) {
    router.push('/login')
    return
  }
  // You could fetch user name from an API. Here we just set a default.
  userName.value = 'Admin'
})

const logout = () => {
  localStorage.removeItem('crm_token')
  alert('You have been logged out.')
  router.push('/login')
}
</script>

<style scoped>
.dashboard-container { max-width: 800px; margin: 50px auto; padding: 20px; border: 1px solid #ccc; border-radius: 8px; }
button { padding: 8px 16px; background-color: #dc3545; color: white; border: none; border-radius: 4px; cursor: pointer; }
</style>