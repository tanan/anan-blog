<template>
  <div class="header">
    <div class="top"><MainLogo /></div>
    <div class="content"><HeaderContent :currentPage="currentPage" /></div>
  </div>    
</template>

<script>
import MainLogo from '@/components/atoms/MainLogo.vue'
import HeaderContent from '@/components/molecules/HeaderContent.vue'
// import AccountInfo from '@/components/molecules/AccountInfo.vue'

export default {
  name: 'MainHeader',
  components: {
    MainLogo,
    HeaderContent,
    // AccountInfo
  },
  data() {
    return {
      currentPage: 'home'
    }
  }
}
</script>

<style lang="scss">
@import "@/styles/_variables.scss";
.header {
  height: 96px;
  background-color: #55c500;

  .top {
    height: 48px;
  }

  .content {
    height: 48px;
  }
}
</style>