<template>
  <div class="home">
    <MainHeader />
    <div class="home-container">
      <ArticleList />
      <Category />
    </div>
  </div>
</template>

<script>
// @ is an alias to /src
import MainHeader from '@/components/organisms/MainHeader.vue'
import ArticleList from '@/components/organisms/ArticleList.vue'
import Category from '@/components/organisms/Category.vue'

export default {
  name: 'Home',
  components: {
    MainHeader,
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