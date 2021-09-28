<template>
  <div class="home">
    <MainLogo />
    <HelloWorld msg="Welcome to Your Vue.js App"/>
  </div>
</template>

<script>
// @ is an alias to /src
import HelloWorld from '@/components/HelloWorld.vue'
import MainLogo from '@/components/atoms/MainLogo.vue'

export default {
  name: 'Home',
  components: {
    HelloWorld,
    MainLogo
  }
}
</script>
