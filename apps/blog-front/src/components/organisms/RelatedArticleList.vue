<template>
  <div class="related-article-section">
    <h2>関連記事</h2>
    <div class="card-list-container">
      <Card v-for="show in shows" :key="show.title" :show="show" />
    </div>
  </div>
</template>

<script>
import Card from '@/components/molecules/Card.vue'
export default {
  name: 'RelatedArticleList',
  components: {
    Card
  },
  props: {
    shows: Object
  }
}
</script>

<style scoped lang="scss">
@import "@/styles/_variables.scss";
  .related-article-section {
    width: 58%;
    margin-right: auto;
    margin-left: auto;
    padding-top: 4%;

    h2 {
      font-size: 1.8em;
      margin: 0;
      // text-align: center;
    }
  }
  .card-list-container {
    display: flex;
    flex-flow: row wrap;
    justify-content: space-between;
  }

  .card-list-container::before{
    content:"";
    display: block;
    width:23%;
    order:1;
  }

  .card-list-container::after {
    display: block;
    content:"";
    width: 23%;
  }
</style>