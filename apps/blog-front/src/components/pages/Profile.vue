<template>
  <div class="profile-page">
    <MainHeader page="profile" />
    <div class="profile">
      <Personal />
      <Portfolio />
      <Career />
    </div>
  </div>
</template>

<script>
import MainHeader from '@/components/organisms/MainHeader.vue'
import Personal from '@/components/organisms/Personal.vue'
import Portfolio from '@/components/organisms/Portfolio.vue'
import Career from '@/components/organisms/Career.vue'
export default {
  name: 'Profile',
  components: {
    MainHeader,
    Personal,
    Portfolio,
    Career
  }
}
</script>

<style scoped lang="scss">
  .profile {
    width: 70%;
    margin-left: auto;
    margin-right: auto;
  }
</style>