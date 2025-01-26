<script setup>
defineProps({
  msg: {
    type: String,
    required: true,
  },
})
</script>

<template>
  <div class="greetings">
    <h1 class="vaultrdb-color">{{ msg }}</h1>
    <h3>
      This project is maintained in  
      <a href="https://github.com/jnnkrdb/vaultrdb" target="_blank" rel="noopener">GitHub</a> under the 
      <a href="https://github.com/jnnkrdb/vaultrdb?tab=MIT-1-ov-file#readme" target="_blank" rel="noopener">MIT License</a>.
    </h3>
  </div>
</template>

<style scoped>
h1 {
  font-weight: 500;
  font-size: 2.6rem;
  position: relative;
  top: -10px;
}

h3 {
  font-size: 1.2rem;
}

.greetings h1,
.greetings h3 {
  text-align: center;
}

@media (min-width: 1024px) {
  .greetings h1,
  .greetings h3 {
    text-align: left;
  }
}
</style>
