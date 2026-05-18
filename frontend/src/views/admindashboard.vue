<template>
  <div class="dashboard-layout">
    <div class="main-content">
      <header>
        <h2>Control Panel Dashboard</h2>
        <div class="user-badge">Welcome, {{ current_user }}! (Role: 
{{ current_role }})</div>
      </header>
      <div class="content-body">
        <h3>System Registered Users List</h3>
        <table class="user-table">
          <thead>
            <tr>
              <th>User ID</th>
              <th>Username</th>
              <th>Email Address</th>
              <th>Account Status</th>
              <th>Actions</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="user in users" :key="user.user_id">
              <td>#{{ user.user_id }}</td>
              <td>{{ user.username }}</td>
              <td>{{ user.email }}</td>
              <td><span class="status-tag">{{ user.status || 'Active' }}</span></td>
              <td><button class="btn-edit" @click="editUser(user.user_id)">Edit</
button> <button class="btn-delete" @click="deleteUser(user.user_id)">Delete</
button></td>
            </tr>
          </tbody>