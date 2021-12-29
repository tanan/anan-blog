<template>
  <div class="portfolio">
    <div class="title">登壇資料・Blog・インタビュー等</div>
    <PortfolioContent />
  </div>
</template>

<script>
import PortfolioContent from '@/components/molecules/PortfolioContent.vue'
export default {
  name: 'Portfolio',
  components: {
    PortfolioContent
  },
}
</script>

<style lang="scss">
  .portfolio {

    .title {
      padding-top: 8%;
      text-align: center;
      font-size: 1.6em;
      font-weight: 600;
    }
  }
</style>