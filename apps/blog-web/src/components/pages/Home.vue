<template>
  <div class="home">
    <div class="home-container">
      <ArticleList />
      <Category />
    </div>
  </div>
</template>

<script>
// @ is an alias to /src
import ArticleList from '@/components/organisms/ArticleList.vue'
import Category from '@/components/organisms/Category.vue'

export default {
  name: 'Home',
  components: {
    ArticleList,
    Category
  },
 }
</script>

<style lang="scss">
@import "@/styles/_variables.scss";
  .home-container {
    display: flex;
  }
</style>