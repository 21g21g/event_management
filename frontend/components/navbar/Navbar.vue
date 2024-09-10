<template>
<!--
    minabevent
    search
    create event
    login
    signup

    the search by location functionality must be done at the backeend

 -->
    <nav class="bg-slate-300 z-1000  p-4">
    <div class="container mx-auto flex items-center justify-between">
      <NuxtLink to="/" exact class="text-xl font-bold text-blue-600 hover:text-blue-800">
        minabevent
      </NuxtLink>

    
      <div class="flex items-center space-x-4">
        
        <NuxtLink to="/auth/login" exact class=" text-black px-4 py-2 rounded hover:bg-amber-500 transition duration-300">
          Login
        </NuxtLink>
        <NuxtLink to="/auth/signup" exact class=" text-black px-4 py-2 rounded hover:bg-amber-500 transition duration-300">
          Signup
        </NuxtLink>
      </div>
    </div>
  </nav>
</template>
<script setup>

</script>
<style scoped>
.router-link-exact-active{
  color: blueviolet !important;
}

</style>