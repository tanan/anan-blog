<template>
  <div class="personal">
    <div class="title">プロフィール</div>
    <PersonalContent />
  </div>
</template>

<script>
import PersonalContent from '@/components/molecules/PersonalContent.vue'
export default {
  name: 'Personal',
  components: {
    PersonalContent
  },
}
</script>

<style lang="scss">
  .personal {
    margin-left: auto;
    margin-right: auto;

    .title {
      padding-top: 48px;
      text-align: center;
      font-size: 1.6em;
      font-weight: 600;
    }
  }
</style>