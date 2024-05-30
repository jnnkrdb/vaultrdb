<template>
  <Menubar :model="items">
    <template #start>
      <img alt="VaultRDB Logo" src="../../assets/logo.png" height="30" width="30" style="margin-left: .5rem; margin-right: .5rem;">
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
      <div class="flex align-items-center gap-2">
        <InputText placeholder="Basicauth:User" type="text" class="w-8rem sm:w-auto"/>
        <InputText placeholder="Basicauth:Pass" type="text" class="w-8rem sm:w-auto" />
      </div>
    </template>
  </Menubar>

  <!--
    <img alt="Vue logo" src="../../assets/logo.png">
  -->
</template>

<script setup>
import { ref } from "vue";
const items = ref([
  { label: "Vault" },
  { label: "InternalDB", disabled: true },
  { 
    label: "Help", 
    items: [
      { label: "GitHub", target: "_blank", url: "https://github.com/jnnkrdb/vaultrdb" },
      { label: "Wiki", target: "_blank", url: "https://github.com/jnnkrdb/vaultrdb/wiki", disabled: true },
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
import InputText from 'primevue/inputtext';

export default {
  name: 'FrontHeader',
  components: {
    Menubar,
    Badge,
    InputText,
  },
  mounted() {
    this.$axios.get('/storageapi/v1/internaldb/buckets')
      .then((response) => {
        console.log(response.data)
      })
  }
}
</script>

<style>
    
</style>