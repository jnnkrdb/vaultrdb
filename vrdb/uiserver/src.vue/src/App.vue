<template>
  <FrontHeader />
  <FrontContent />
  <FrontFooter />
</template>

<script>
import FrontHeader from './components/content/FrontHeader.vue'
import FrontContent from './components/content/FrontContent.vue'
import FrontFooter from './components/content/FrontFooter.vue'

export default {
  name: 'App',
  components: {
    FrontHeader,
    FrontContent,
    FrontFooter
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 0px;
}
</style>
