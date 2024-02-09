

<script setup>
import TheWelcome from './components/TheWelcome.vue'
</script>

<template>
  <main>
    <MDBNavbar expand="lg" light bg="light" container>
      <MDBNavbarBrand href="#">
        <img alt="Vue logo" src="./assets/logo.svg" height="25" width="25"/>
      </MDBNavbarBrand>
      <MDBNavbarToggler @click="collapse1 = !collapse1" target="#navbarSupportedContent"></MDBNavbarToggler>
      <MDBCollapse v-model="collapse1" id="navbarSupportedContent">
        <MDBNavbarNav class="mb-2 mb-lg-0">
          <MDBNavbarItem to="#" active>Home</MDBNavbarItem>
          <MDBNavbarItem to="#" >Link1</MDBNavbarItem>
          <MDBNavbarItem to="#" >Link2</MDBNavbarItem>
          <MDBNavbarItem to="#" disabled>Disabled</MDBNavbarItem>
        </MDBNavbarNav>
        <form class="d-flex input-group w-auto">
          <MDBBtn outline="primary">Login</MDBBtn>
        </form>
      </MDBCollapse>
    </MDBNavbar>

    <TheWelcome />
  </main>
</template>

<style scoped>
#app {
  font-family: Roboto, Helvetica, Arial, sans-serif;
}
</style>

<script>
  import {
    MDBBtn,
    MDBNavbar,
    MDBNavbarToggler,
    MDBNavbarBrand,
    MDBNavbarNav,
    MDBNavbarItem,
    MDBCollapse,
    MDBDropdown,
    MDBDropdownToggle,
    MDBDropdownMenu,
    MDBDropdownItem
  } from 'mdb-vue-ui-kit';
  import { ref } from 'vue';

  export default {
    components: {
      MDBBtn,
      MDBNavbar,
      MDBNavbarToggler,
      MDBNavbarBrand,
      MDBNavbarNav,
      MDBNavbarItem,
      MDBCollapse,
      MDBDropdown,
      MDBDropdownToggle,
      MDBDropdownMenu,
      MDBDropdownItem
    },
    setup() {
      const collapse1 = ref(false);
      const dropdown1 = ref(false);
      return {
        collapse1,
        dropdown1
      }
    }
  };
</script>