<template>
  <div class="header">
    <MainLogo />
  </div>    
</template>

<script>
import MainLogo from '@/components/atoms/MainLogo.vue'
// import HeaderContent from '@/components/molecues/HeaderContent.vue'
// import AccountInfo from '@/components/molecues/AccountInfo.vue'

export default {
  name: 'MainHeader',
  components: {
    MainLogo,
    // HeaderContent,
    // AccountInfo
  }
}
</script>

<style lang="scss">
@import "@/styles/_variables.scss";
.header {
  height: 96px;
  background-color: green;
}
</style>