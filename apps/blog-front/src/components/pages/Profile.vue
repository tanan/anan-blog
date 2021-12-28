<template>
  <div class="profile-page">
    <MainHeader page="profile" />
    <Personal />
    <Portfolio />
    <Career />
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