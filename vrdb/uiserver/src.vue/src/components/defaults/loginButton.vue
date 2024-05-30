<template>
  <div class="flex align-items-center gap-2">
    <Button label="Login" @click="visible = true"></Button>
    <Dialog v-model:visible="visible" modal header="Basic-Auth" :style="{ width: '25rem' }">
      <span class="p-text-secondary block mb-5">Login with your configured basic-auth Credentials.</span>
      <div class="flex align-items-center gap-3 mb-3">
          <label for="username" class="font-semibold w-6rem">Username</label>
          <InputText id="username" class="flex-auto" autocomplete="off" v-model="this.user" />
      </div>
      <div class="flex align-items-center gap-3 mb-5">
          <label for="password" class="font-semibold w-6rem">Password</label>
          <Password id="password" class="flex-auto" autocomplete="off" v-model="this.pass" :feedback="false"/>
      </div>
      <div class="flex justify-content-end gap-2">
          <Button type="button" label="Cancel" severity="secondary" @click="visible = false"></Button>
          <Button type="button" label="Save" @click="visible = false; UpdateUserInfo()"></Button>
      </div>
    </Dialog>

    <!--

    <InputText v-tooltip.bottom="'Enter your username'" placeholder="Basicauth:User" type="text" class="w-8rem sm:w-auto"/>
    <InputText v-tooltip.bottom="'Enter your password'" placeholder="Basicauth:Pass" type="text" class="w-8rem sm:w-auto" />

    -->
  </div>
</template>

<script setup>
const visible = ref(false);

</script>

<script>
import { ref } from "vue";
import Dialog from 'primevue/dialog';
import Button from 'primevue/button';
import Password from 'primevue/password';
import InputText from 'primevue/inputtext';

var user = ref(null);
var pass = ref(null);


export default {
  name: 'LoginButton',
  components: {
    InputText,
    Dialog,
    Button,
    Password,
  },
  beforeMount: {

  },
  methods: {
    UpdateUserInfo() {
      this.$username = this.user;
      this.$password = this.pass;
      console.log(this.user + ' == ' + this.$username, this.pass + ' == ' + this.$password)
    }
  }
}
</script>

<style>
</style>