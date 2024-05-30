<template>
  <Menubar :model="navbarItems">
    <template #start>
      <img alt="VaultRDB Logo" src="./assets/logo.png" height="30" width="30" style="margin-left: .5rem; margin-right: .5rem;">
    </template>
    <template #item="{ item, props, hasSubmenu, root }">
      <a v-ripple class="flex align-items-center" :href="item.url" :target="item.target" v-bind="props.action">
        <span :class="item.icon" />
        <span class="ml-2">{{ item.label }}</span>
        <Badge v-if="item.badge" :class="{ 'ml-auto': !root, 'ml-2': root }" :value="item.badge"/>
        <i v-if="hasSubmenu" :class="['pi pi-angle-down', { 'pi-angle-down ml-2': root, 'pi-angle-right ml-auto': !root }]"></i>
      </a>
    </template>
    <template #end>
      <LoginButton />
    </template>
  </Menubar>


  <main>
    <div class="content">
      <FrontContent />
    </div>
  </main>
  
  
  <FrontFooter />
</template>

<script setup>
import { ref } from "vue";
const navbarItems = ref([
    { label: "Vault" },
    { label: "InternalDB" },
    { 
      label: "Help", 
      items: [
        { label: "GitHub", target: "_blank", url: "https://github.com/jnnkrdb/vaultrdb" },
        { label: "Wiki", target: "_blank", url: "https://github.com/jnnkrdb/vaultrdb/wiki" },
        { label: "Swagger", target: "_blank", url: "/swagger/", disabled: true }
      ]
    }
  ])
</script>

<script>
import "primeflex/primeflex.css";
import "primevue/resources/themes/aura-light-green/theme.css";
import "primeicons/primeicons.css";
import Menubar from 'primevue/menubar'
import Badge from 'primevue/badge';

import LoginButton from './components/defaults/loginButton.vue'
import FrontContent from './components/content/FrontContent.vue'
import FrontFooter from './components/content/FrontFooter.vue'

export default {
  name: 'App',
  components: {
    FrontContent,
    FrontFooter,
    Menubar,
    Badge,
    LoginButton,
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
