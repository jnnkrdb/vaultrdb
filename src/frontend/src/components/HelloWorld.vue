<template>
  <h1>{{ msg }}</h1>

  <p>
    <a href="https://vitejs.dev/guide/features.html" target="_blank">
      Vite Documentation
    </a>
    |
    <a href="https://v3.vuejs.org/" target="_blank">Vue 3 Documentation</a>
  </p>

  <MDBBtn color="primary" class="mb-3" @click="state.count++">
    count is: {{ state.count }}
  </MDBBtn>

  <p>
    Edit components/HelloWorld.vue to test hot module replacement.
  </p>

</template>

<script setup lang="ts">
  import { defineProps, reactive } from "vue";
  import { MDBBtn } from "mdb-vue-ui-kit";

  defineProps({
    msg: String,
  });

  const state = reactive({ count: 10 });
</script>

<style scoped>
  a {
    color: #42b983;
  }
</style>