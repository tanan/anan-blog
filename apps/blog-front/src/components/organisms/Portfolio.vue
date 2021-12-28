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
    width: 800px;
    margin-left: auto;
    margin-right: auto;

    .title {
      padding-top: 128px;
      text-align: center;
      font-size: 28px;
      font-weight: 600;
    }
  }
</style>