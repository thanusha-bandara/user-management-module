import { createRouter, createWebHistory } from 'vue-router'
import Login from '../views/Login.vue'
import Register from '../views/Register.vue'
import AdminDashboard from '../views/AdminDashboard.vue'

const routes = [
  { path: '/', redirect: '/login' },
  { path: '/login', component: Login },
  { path: '/register', component: Register },
  { 
    path: '/dashboard', 
    component: AdminDashboard,
    meta: { requiresAuth: true } // Mark this route as protected
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

// Navigation Guard (Global Auth Guard)
router.beforeEach((to, from, next) => {
  const token = localStorage.getItem('crm_token')
  
  // If a user tries to go to /dashboard without logging in, block and redirect to /login
  if (to.matched.some(record => record.meta.requiresAuth) && !token) {
    alert("Please log in first to access the system!")
    next('/login')
  } else {
    next()
  }
})

export default router