<template>
  <div class="article">
    <img class="thumbnail" :src="article.thumbnail" alt="thumbnail">
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
  .article {
    width: 620px;
    margin-left: auto;
    margin-right: auto;

    font-family: -apple-system,BlinkMacSystemFont,Helvetica Neue,Segoe UI,Hiragino Kaku Gothic ProN,Hiragino Sans,ヒラギノ角ゴ ProN W3,Arial,メイリオ,Meiryo,sans-serif;

    h1 {
      font-size: 32px;
    }

    p {
      letter-spacing: 1px;
      line-height: 36px;
    }

    .thumbnail {
      width: 640px;
      height: 360px;
      margin-top: 24px;
    }

    ul {
      margin-left: 12px;
      p {
        width: 80%;
        margin-right: auto;
      }
    }
  }
</style>