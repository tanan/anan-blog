<template>
  <div class="health-care">
    <MainHeader page="health-care" />
    <div class="health-care">
      <HealthData />
    </div>
  </div>
</template>

<script>
import MainHeader from '@/components/organisms/MainHeader.vue'
import HealthData from '@/components/organisms/HealthData.vue'
export default {
  name: 'Profile',
  components: {
    MainHeader,
    HealthData,
  }
}
</script>

<style scoped lang="scss">
  .profile {
    width: 76%;
    margin-left: auto;
    margin-right: auto;
  }
</style>