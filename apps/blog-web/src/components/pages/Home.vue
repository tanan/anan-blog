<template>
  <div class="home">
    <div class="main-logo"><MainLogo /></div>
    <div class="home-container">
      <ArticleList />
      <Category />
    </div>
  </div>
</template>

<script>
// @ is an alias to /src
import MainLogo from '@/components/atoms/MainLogo.vue'
import ArticleList from '@/components/organisms/ArticleList.vue'
import Category from '@/components/organisms/Category.vue'

export default {
  name: 'Home',
  components: {
    MainLogo,
    ArticleList,
    Category
  },
 }
</script>

<style lang="scss">
@import "@/styles/_variables.scss";
  .main-logo {
    margin-top: 96px;
    margin-bottom: 48px;
  }
  .home-container {
    display: flex;
  }
</style>