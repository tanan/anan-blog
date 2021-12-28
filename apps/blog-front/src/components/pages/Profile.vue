<template>
  <div class="profile-page">
    <MainHeader page="profile" />
    <Personal />
    <Career />
  </div>
</template>

<script>
import MainHeader from '@/components/organisms/MainHeader.vue'
import Career from '@/components/organisms/Career.vue'
import Personal from '@/components/organisms/Personal.vue'
export default {
  name: 'Profile',
  components: {
    MainHeader,
    Personal,
    Career
  }
}
</script>