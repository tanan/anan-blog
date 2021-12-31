<template>
  <div class="article">
    <div class="thumbnail">
      <img :src="article.thumbnail" alt="thumbnail">
    </div>
    <h1>{{ article.title }}</h1>
    <div class="author">
      <div class="date">
        <p class="published-date">投稿日<time :datetime="article.firstPublishedAt">{{ getDisplayDate(article.firstPublishedAt) }}</time></p>
        <p class="updated-date">更新日<time :datetime="article.publishedAt">{{ getDisplayDate(article.firstPublishedAt) }}</time></p>
      </div>
    </div>
    <div class="category">
      <div v-for="name in getCategories(article.category)" :key="name" class="name">#{{ name }}</div>
    </div>
    <p>{{ article.description }}</p>
    <p v-html="article.content"></p>
  </div>
</template>

<script>
export default {
  name: 'ArticleText',
  props: {
    article: Object
  },
  methods: {
    getDisplayDate (d) {
      let date = new Date(d)
      return date.getFullYear() + "年" + date.getMonth() + "月" + date.getDate() + "日"
    },
    getCategories (category) {
      console.log(category)
      if (!category) {
        return null
      }
      return category.split(',')
    }
  },
}
</script>

<style lang="scss">
@import "@/styles/_variables.scss";
  .article {
    width: 58%;
    margin-left: auto;
    margin-right: auto;

    font-family: -apple-system,BlinkMacSystemFont,Helvetica Neue,Segoe UI,Hiragino Kaku Gothic ProN,Hiragino Sans,ヒラギノ角ゴ ProN W3,Arial,メイリオ,Meiryo,sans-serif;

    h1 {
      font-size: 1.8em;
      margin-bottom: 4px;
    }

    .author {
      font-size: 0.8em;

      p {
        line-height: 12px;
        margin-top: 4px;
      }

      .date {
        display: flex;
        color: #a8abb1;

        .updated-date {
          margin-left: 16px;
        }

        time {
          margin-left: 8px;
        }
      }
    }

    p {
      letter-spacing: 1px;
      line-height: 200%;
    }

    .thumbnail {
      height: auto;
      text-align: center;
      margin-top: 24px;
      img {
        width: 80%;
      }
    }

    .category {
      display: flex;

      .name {
        font-size: 0.8em;
        margin-right: 6px;
        padding: 4px 4px;
        border: 1px solid $color-border-gray;
        border-radius: 2px;
      }
    }

    ul {
      p {
        width: 80%;
        margin-right: auto;
        margin-top: 0px;
        margin-bottom: 0px;
        line-height: 24px;
        font-size: 1em;
      }
    }

    pre {
      line-height: 20px;
      font-size: 0.9em;
      background-color: whitesmoke;
      padding: 16px;
      white-space: pre-wrap;
    }
  }
</style>