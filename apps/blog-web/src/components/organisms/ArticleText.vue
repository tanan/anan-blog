<template>
  <div class="article">
    <h1>{{ article.title }}</h1>
    <p>{{ article.description }}</p>
    <p v-html="article.content"></p>
  </div>
</template>

<script>
export default {
  name: 'ArticleText',
  props: {
    article: Object
  }
}
</script>

<style lang="scss">
  
</style>