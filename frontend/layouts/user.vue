<template>
  <div class="flex flex-col min-h-screen">

    <div class="flex flex-grow ">     
       <aside class=" w-20 md:w-64 bg-slate-50 text-white min-h-screen border-r-2 fixed left-0 right-0">
      <UserSidebar/>
       </aside>
     
         <div class="flex flex-col w-full">
          <header  class="bg-slate-50 text-white  w-full fixed  h-16">
          <UserNavbar/>
        </header>
    <hr/>
    
        <main class="flex-grow p-4 mt-4 ml-20 md:ml-64">
      <NuxtPage />
    </main>

     </div>
     
    </div>
  
    
    <footer class="w-full z-50">
      <Footer/>
    </footer>
  </div>
</template>
<script setup>
import UserNavbar from "../components/userlayout/UserNavbar.vue";
import UserSidebar from "../components/userlayout/UserSidebar.vue";
</script>

