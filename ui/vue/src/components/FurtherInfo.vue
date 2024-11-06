<script setup>
import WelcomeItem from './WelcomeItem.vue'
import GitIcon from './icons/IconGit.vue'
import LicenseIcon from './icons/IconLicense.vue'
</script>

<template>
  <WelcomeItem>
    <template #icon>
      <GitIcon />
    </template>
    <template #heading>Version</template>

    Currently rolling under <code>v0.0.1</code>

    <br>

    See other versions under 
    <a href="https://github.com/jnnkrdb/vaultrdb/pkgs/container/vaultrdb" target="_blank" rel="noopener"><code>Packages</code></a>.
  </WelcomeItem>

  
  <WelcomeItem>
    <template #icon>
      <LicenseIcon />
    </template>
    <template #heading>License</template>

    Licensed under <a href="https://github.com/jnnkrdb/vaultrdb/blob/master/LICENSE" target="_blank" rel="noopener">MIT-License</a>

  </WelcomeItem>

</template>
