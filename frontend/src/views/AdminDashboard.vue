<template>
  <div class="dashboard-layout">
    <div class="main-content">
      <header>
        <h2>{{ isAdmin ? 'Control Panel Dashboard' : 'User Profile' }}</h2>
        <div class="user-badge">Welcome, {{ current_user }}! (Role: {{ current_role_name }})</div>
      </header>

      <div class="content-body" v-if="isAdmin">
        <h3>System Registered Users List</h3>
        <table class="user-table">
          <thead>
            <tr>
              <th>User ID</th>
              <th>Full Name</th>
              <th>Username</th>
              <th>Email Address</th>
              <th>Account Status</th>
              <th>Actions</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="user in users" :key="user.user_id">
              <td>#{{ user.user_id }}</td>
              <td><strong>{{ user.full_name }}</strong></td>
              <td>{{ user.username }}</td>
              <td>{{ user.email }}</td>
              <td><span class="status-tag" :class="(user.status || 'active').toLowerCase()">{{ user.status || 'Active' }}</span></td>
              <td v-if="isAdmin">
                <button class="btn-edit" @click="editUser(user)">Edit</button>
                <button class="btn-delete" @click="deleteUser(user.user_id)">Delete</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <div class="content-body" v-else>
        <h3>Your Profile Information</h3>
        <p><strong>Full Name:</strong> {{ users[0]?.full_name }}</p>
        <p><strong>Username:</strong> {{ users[0]?.username }}</p>
        <p><strong>Email:</strong> {{ users[0]?.email }}</p>
        <p><strong>Status:</strong> <span class="status-tag" :class="(users[0]?.status || 'active').toLowerCase()">{{ users[0]?.status || 'Active' }}</span></p>
        
        <div style="margin-top: 20px;">
          <button class="btn-primary" @click="editProfile" style="margin-right: 10px;">Edit Profile</button>
          <button class="btn-secondary" @click="changePassword">Change Password</button>
        </div>

        <p style="margin-top: 15px; color: #6b7280; font-size: 14px;">
          * Only administrators have full access to manage other users in the CRM system.
        </p>
      </div>
    </div>

    <!-- Edit User Modal (Admin) -->
    <div v-if="showModal" class="modal-overlay" @click.self="showModal = false">
      <div class="modal-content">
        <h3>Edit User</h3>
        <div class="form-group">
          <label>Full Name</label>
          <input type="text" v-model="editData.full_name" class="form-control" />
        </div>
        <div class="form-group">
          <label>Username</label>
          <input type="text" v-model="editData.username" class="form-control" />
        </div>
        <div class="form-group">
          <label>Email</label>
          <input type="email" v-model="editData.email" class="form-control" />
        </div>
        <div class="form-group">
          <label>Role</label>
          <select v-model="editData.role_id" class="form-control">
            <option value="1">Admin</option>
            <option value="2">Sales Agent</option>
            <option value="3">Manager</option>
            <option value="4">Customer</option>
          </select>
        </div>
        <div class="form-group">
          <label>Status</label>
          <select v-model="editData.status" class="form-control">
            <option value="active">Active</option>
            <option value="inactive">Inactive</option>
          </select>
        </div>
        <div class="modal-actions">
          <button class="btn-secondary" @click="showModal = false">Cancel</button>
          <button class="btn-primary" @click="saveUser">Save Changes</button>
        </div>
      </div>
    </div>

    <!-- Edit Own Profile Modal -->
    <div v-if="showProfileModal" class="modal-overlay" @click.self="showProfileModal = false">
      <div class="modal-content">
        <h3>Edit Profile</h3>
        <div class="form-group">
          <label>Full Name</label>
          <input type="text" v-model="profileData.full_name" class="form-control" />
        </div>
        <div class="form-group">
          <label>Username</label>
          <input type="text" v-model="profileData.username" class="form-control" />
        </div>
        <div class="form-group">
          <label>Email</label>
          <input type="email" v-model="profileData.email" class="form-control" />
        </div>
        <div class="modal-actions">
          <button class="btn-secondary" @click="showProfileModal = false">Cancel</button>
          <button class="btn-primary" @click="saveProfile">Save Changes</button>
        </div>
      </div>
    </div>

    <!-- Change Password Modal -->
    <div v-if="showPasswordModal" class="modal-overlay" @click.self="showPasswordModal = false">
      <div class="modal-content">
        <h3>Change Password</h3>
        <div class="form-group">
          <label>Current Password</label>
          <input type="password" v-model="passwordData.old_password" class="form-control" />
        </div>
        <div class="form-group">
          <label>New Password</label>
          <input type="password" v-model="passwordData.new_password" class="form-control" />
        </div>
        <div class="modal-actions">
          <button class="btn-secondary" @click="showPasswordModal = false">Cancel</button>
          <button class="btn-primary" @click="savePassword">Change Password</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'

const users = ref([])
const current_user = localStorage.getItem('crm_username') || 'User'
const roleId = localStorage.getItem('crm_role')
const isAdmin = ref(roleId === '1')

const roleNames = { '1': 'Admin', '2': 'Sales Agent', '3': 'Manager', '4': 'Customer' }
const current_role_name = roleNames[roleId] || 'Staff'

const showModal = ref(false)
const editData = ref({ user_id: null, full_name: '', username: '', email: '', role_id: 1, status: '' })

const showProfileModal = ref(false)
const profileData = ref({ full_name: '', username: '', email: '' })

const showPasswordModal = ref(false)
const passwordData = ref({ old_password: '', new_password: '' })

const fetchUsers = async () => {
  try {
    const token = localStorage.getItem('crm_token')
    
    if (isAdmin.value) {
      const res = await axios.get('http://localhost:8080/api/v1/users', {
        headers: { Authorization: `Bearer ${token}` }
      })
      users.value = res.data || []
    } else {
      const payload = JSON.parse(atob(token.split('.')[1]))
      const userId = payload.user_id
      
      const res = await axios.get(`http://localhost:8080/api/v1/users/${userId}`, {
        headers: { Authorization: `Bearer ${token}` }
      })
      users.value = [res.data]
    }
  } catch (err) {
    console.error("Failed to fetch users", err)
    if(err.response?.status === 401) {
      alert("Session expired, please login again.")
    }
  }
}

const editUser = (user) => {
  editData.value = { ...user }
  showModal.value = true
}

const saveUser = async () => {
  try {
    const token = localStorage.getItem('crm_token')
    // Parse role_id to int to match backend Go struct expectation
    editData.value.role_id = parseInt(editData.value.role_id)
    await axios.put(`http://localhost:8080/api/v1/users/${editData.value.user_id}`, editData.value, {
      headers: { Authorization: `Bearer ${token}` }
    })
    alert('User updated successfully')
    showModal.value = false
    fetchUsers()
  } catch (err) {
    console.error(err)
    alert(err.response?.data?.error || 'Failed to update user')
  }
}

const deleteUser = async (id) => {
  if (confirm(`Are you sure you want to delete user #${id}?`)) {
    try {
      const token = localStorage.getItem('crm_token')
      await axios.delete(`http://localhost:8080/api/v1/users/${id}`, {
        headers: { Authorization: `Bearer ${token}` }
      })
      alert('User deleted successfully')
      fetchUsers()
    } catch (err) {
      console.error(err)
      alert('Failed to delete user')
    }
  }
}

const editProfile = () => {
  profileData.value = { full_name: users.value[0]?.full_name, username: users.value[0]?.username, email: users.value[0]?.email }
  showProfileModal.value = true
}

const saveProfile = async () => {
  try {
    const token = localStorage.getItem('crm_token')
    await axios.put('http://localhost:8080/api/v1/users/me/profile', profileData.value, {
      headers: { Authorization: `Bearer ${token}` }
    })
    alert('Profile updated successfully')
    showProfileModal.value = false
    fetchUsers()
  } catch (err) {
    console.error(err)
    alert('Failed to update profile')
  }
}

const changePassword = () => {
  passwordData.value = { old_password: '', new_password: '' }
  showPasswordModal.value = true
}

const savePassword = async () => {
  if (passwordData.value.new_password.length < 6) {
    alert("New password must be at least 6 characters.")
    return
  }
  try {
    const token = localStorage.getItem('crm_token')
    await axios.put('http://localhost:8080/api/v1/users/me/password', passwordData.value, {
      headers: { Authorization: `Bearer ${token}` }
    })
    alert('Password changed successfully')
    showPasswordModal.value = false
  } catch (err) {
    console.error(err)
    alert(err.response?.data?.error || 'Failed to change password')
  }
}

onMounted(() => fetchUsers())
</script>

<style scoped>
.dashboard-layout { font-family: 'Inter', sans-serif; background: transparent; min-height: 100vh; padding: 40px 20px; color: #e2e8f0; }
.main-content { max-width: 1200px; margin: 0 auto; }
header { display: flex; flex-direction: column; align-items: center; justify-content: center; padding-bottom: 25px; margin-bottom: 35px; text-align: center; }
header h2 { font-size: 28px; color: #f8fafc; margin-bottom: 15px; font-weight: 700; }
.user-badge { background: rgba(99, 102, 241, 0.1); color: #a5b4fc; padding: 8px 16px; border-radius: 20px; font-weight: 500; font-size: 14px; border: 1px solid rgba(99, 102, 241, 0.2); }
.content-body { background: rgba(30, 41, 59, 0.5); backdrop-filter: blur(12px); padding: 30px; border-radius: 16px; border: 1px solid rgba(255, 255, 255, 0.05); box-shadow: 0 10px 30px rgba(0,0,0,0.2); margin-bottom: 30px; }
.content-body h3 { color: #f8fafc; margin-top: 0; margin-bottom: 25px; font-size: 20px; border-bottom: 1px solid rgba(255, 255, 255, 0.05); padding-bottom: 15px; }

/* Table Styles */
.user-table { width: 100%; border-collapse: separate; border-spacing: 0; text-align: left; }
.user-table th { background: rgba(15, 23, 42, 0.6); padding: 15px; color: #94a3b8; font-weight: 600; font-size: 13px; text-transform: uppercase; letter-spacing: 0.5px; border-bottom: 1px solid rgba(255, 255, 255, 0.05); }
.user-table th:first-child { border-top-left-radius: 8px; }
.user-table th:last-child { border-top-right-radius: 8px; }
.user-table td { padding: 16px 15px; border-bottom: 1px solid rgba(255, 255, 255, 0.05); color: #cbd5e1; font-size: 14px; }
.user-table tbody tr { transition: background 0.2s ease; }
.user-table tbody tr:hover { background: rgba(255, 255, 255, 0.02); }
.user-table tbody tr:last-child td { border-bottom: none; }
.user-table td strong { color: #f8fafc; font-weight: 600; }

/* Status Tags */
.status-tag { padding: 6px 12px; border-radius: 20px; font-size: 12px; font-weight: 600; text-transform: uppercase; letter-spacing: 0.5px; border: 1px solid transparent; }
.status-tag.active { background: rgba(16, 185, 129, 0.1); color: #34d399; border-color: rgba(52, 211, 153, 0.2); }
.status-tag.inactive { background: rgba(239, 68, 68, 0.1); color: #f87171; border-color: rgba(248, 113, 113, 0.2); }
.status-tag:not(.active):not(.inactive) { background: rgba(99, 102, 241, 0.1); color: #818cf8; border-color: rgba(129, 140, 248, 0.2); }

/* Buttons */
.btn-edit { background: rgba(59, 130, 246, 0.1); color: #60a5fa; border: 1px solid rgba(96, 165, 250, 0.3); padding: 8px 14px; border-radius: 6px; cursor: pointer; margin-right: 8px; transition: all 0.2s; font-size: 13px; font-weight: 500; }
.btn-edit:hover { background: rgba(59, 130, 246, 0.2); color: #93c5fd; }
.btn-delete { background: rgba(239, 68, 68, 0.1); color: #f87171; border: 1px solid rgba(248, 113, 113, 0.3); padding: 8px 14px; border-radius: 6px; cursor: pointer; transition: all 0.2s; font-size: 13px; font-weight: 500; }
.btn-delete:hover { background: rgba(239, 68, 68, 0.2); color: #fca5a5; }

.btn-primary { background: linear-gradient(to right, #6366f1, #8b5cf6); color: white; border: none; padding: 10px 20px; border-radius: 8px; cursor: pointer; font-weight: 600; transition: all 0.3s ease; box-shadow: 0 4px 6px rgba(99, 102, 241, 0.2); }
.btn-primary:hover { box-shadow: 0 6px 8px rgba(99, 102, 241, 0.3); transform: translateY(-1px); }
.btn-secondary { background: rgba(255, 255, 255, 0.05); color: #e2e8f0; border: 1px solid rgba(255, 255, 255, 0.1); padding: 10px 20px; border-radius: 8px; cursor: pointer; font-weight: 500; transition: all 0.2s; }
.btn-secondary:hover { background: rgba(255, 255, 255, 0.1); }

/* Modal Styles */
.modal-overlay { position: fixed; top: 0; left: 0; right: 0; bottom: 0; background: rgba(0,0,0,0.6); backdrop-filter: blur(5px); display: flex; align-items: center; justify-content: center; z-index: 1000; }
.modal-content { background: #1e293b; padding: 35px; border-radius: 16px; width: 420px; max-width: 90%; box-shadow: 0 20px 25px -5px rgba(0,0,0,0.5), 0 10px 10px -5px rgba(0,0,0,0.3); border: 1px solid rgba(255, 255, 255, 0.1); }
.modal-content h3 { margin-top: 0; margin-bottom: 25px; color: #f8fafc; font-size: 22px; font-weight: 700; }
.form-group { margin-bottom: 18px; }
.form-group label { display: block; margin-bottom: 8px; font-weight: 600; color: #cbd5e1; font-size: 13px; text-transform: uppercase; letter-spacing: 0.5px; }
.form-control { width: 100%; padding: 12px 15px; background: rgba(15, 23, 42, 0.6); color: #f8fafc; border: 1px solid rgba(255, 255, 255, 0.05); border-radius: 8px; font-size: 14px; box-sizing: border-box; outline: none; transition: all 0.3s ease; }
.form-control:focus { border-color: #6366f1; box-shadow: 0 0 0 3px rgba(99, 102, 241, 0.2); }
select.form-control option { background: #1e293b; color: #f8fafc; }
.modal-actions { display: flex; justify-content: flex-end; gap: 12px; margin-top: 30px; }
</style>