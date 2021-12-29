<template>
  <div class="home">
    <MainHeader page="home" />
    <div class="home-container">
      <ArticleList />
    </div>
  </div>
</template>

<script>
// @ is an alias to /src
import MainHeader from '@/components/organisms/MainHeader.vue'
import ArticleList from '@/components/organisms/ArticleList.vue'

export default {
  name: 'Home',
  components: {
    MainHeader,
    ArticleList,
  },
 }
</script>

<style lang="scss">
@import "@/styles/_variables.scss";
  .home-container {
    display: flex;
  }
</style>