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
    meta: { requiresAuth: true } // ආරක්ෂිත පිටුවක් බව හැඟවීමට meta දත්ත දීම
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

// Navigation Guard (Global Auth Guard)
router.beforeEach((to, from, next) => {
  const token = localStorage.getItem('crm_token')
  
  // යමෙක් ලොගින් නොවී /dashboard යන්න හැදුවොත් ඔහුව බ්ලොක් කර /login වෙත හරවා යැවීම
  if (to.matched.some(record => record.meta.requiresAuth) && !token) {
    alert("කරුණාකර පද්ධතියට ප්‍රවේශ වීමට මුලින්ම ලොග් වන්න!")
    next('/login')
  } else {
    next()
  }
})

export default router