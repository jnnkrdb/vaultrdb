<script setup>
import WelcomeItem from './WelcomeItem.vue'
import DocumentationIcon from './icons/IconDocumentation.vue'
import EcosystemIcon from './icons/IconEcosystem.vue'
import CommunityIcon from './icons/IconCommunity.vue'
</script>

<template>
  <WelcomeItem>
    <template #icon>
      <DocumentationIcon />
    </template>
    <template #heading>Documentation</template>

    VaultRDB´s
    <a href="https://github.com/jnnkrdb/vaultrdb/wiki" target="_blank" rel="noopener">official documentation</a>
    provides you with all information you need to get started.
  </WelcomeItem>

  <WelcomeItem>
    <template #icon>
      <EcosystemIcon />
    </template>
    <template #heading>Environment</template>

    This project is deployable on 
    <a href="https://kubernetes.io/" target="_blank" rel="noopener">Kubernetes</a>. The 
    deployment manifests can be found under 
    <a href="https://github.com/jnnkrdb/vaultrdb/tree/master/.docs/deployment" target="_blank" rel="noopener"><code>.docs/deployment/</code></a>.

    <br /> 

    More instructions are available in <code>README.md</code> or <code>Wiki</code>.
  </WelcomeItem>

  <WelcomeItem>
    <template #icon>
      <CommunityIcon />
    </template>
    <template #heading>Issues</template>

    Got stuck? Ask your question on 
    <a href="https://github.com/jnnkrdb/vaultrdb/discussions" target="_blank" rel="noopener">VaultRDB Discussions</a> or
    submit an issue to 
    <a href="https://github.com/jnnkrdb/vaultrdb/issues" target="_blank" rel="noopener">VaultRDB Issues</a>.
  </WelcomeItem>
</template>
