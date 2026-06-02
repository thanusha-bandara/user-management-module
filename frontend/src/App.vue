<template>
  <div id="app-root">
    <router-view />
  </div>
</template>

<style>
@import url('https://fonts.googleapis.com/css2?family=Inter:wght@400;500;600;700&display=swap');

body {
  margin: 0;
  padding: 0;
  background-color: #0b0f19;
  background-image: 
    radial-gradient(circle at 15% 50%, rgba(37, 99, 235, 0.12), transparent 25%),
    radial-gradient(circle at 85% 30%, rgba(147, 51, 234, 0.12), transparent 25%);
  background-attachment: fixed;
  color: #e2e8f0;
  font-family: 'Inter', sans-serif;
  min-height: 100vh;
}
</style>