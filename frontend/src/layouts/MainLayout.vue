<template>
  <div class="main-layout">
    <!-- Navbar -->
    <nav class="navbar">
      <div class="logo">CRM System</div>
      <div class="user-info">
        <span>{{ username }}</span>
        <button @click="logout" class="logout-btn">Logout</button>
      </div>
    </nav>

    <div class="layout-body">
      <!-- Sidebar -->
      <aside class="sidebar">
        <ul>
          <li><router-link to="/dashboard">Dashboard</router-link></li>
          <li><router-link to="/users">Users</router-link></li>
          <li><router-link to="/profile">Profile</router-link></li>
        </ul>
      </aside>

      <!-- Dynamic content -->
      <main class="main-content">
        <slot />
      </main>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()
const username = ref('')

onMounted(() => {
  username.value = localStorage.getItem('crm_username') || 'User'
})

const logout = () => {
  localStorage.removeItem('crm_token')
  localStorage.removeItem('crm_username')
  localStorage.removeItem('crm_role')
  router.push('/login')
}
</script>

<style scoped>
.main-layout {
  display: flex;
  flex-direction: column;
  height: 100vh;
}
.navbar {
  background: #1f2937;
  color: white;
  padding: 1rem;
  display: flex;
  justify-content: space-between;
  align-items: center;
}
.logout-btn {
  background: #ef4444;
  border: none;
  padding: 0.3rem 0.8rem;
  border-radius: 5px;
  color: white;
  cursor: pointer;
}
.layout-body {
  display: flex;
  flex: 1;
}
.sidebar {
  width: 250px;
  background: #f3f4f6;
  padding: 1rem;
}
.sidebar ul {
  list-style: none;
  padding: 0;
}
.sidebar li {
  margin: 1rem 0;
}
.sidebar a {
  text-decoration: none;
  color: #374151;
}
.main-content {
  flex: 1;
  padding: 1rem;
  overflow-y: auto;
}
</style>