<script setup>


</script>
<template>
     <div class="flex flex-col min-h-screen"> 
     <header class="fixed top-0 left-0 right-0  z-50">
     <Navbar />
    </header>
      <main class="flex-grow"> 
      <NuxtPage/>  
       </main>
        <footer class="z-10">
        <Footer/>
        </footer>
    </div>
   
</template>
