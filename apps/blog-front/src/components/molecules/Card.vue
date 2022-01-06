<template>
  <div class="card">
    <a :href="getArticleUrl(show.sys.id)">
      <img class="thumbnail" :src="show.thumbnail.url">
      <div class="card-container">
        <div class="title">{{ show.title }}</div>
        <div class="date">
          <p class="published-date">更新日<time :datetime="show.sys.publishedAt">{{ getDisplayDate(show.sys.publishedAt) }}</time></p>
        </div>
        <div class="category">
          <div v-for="name in getCategories(show.category)" :key="name" class="name">#{{ name }}</div>
        </div>
        <div class="description">{{ show.description }}</div>
      </div>
    </a>
  </div>
</template>

<script>
export default {
  name: 'Card',
  props: {
    show: Object
  },
  methods: {
    getArticleUrl (id) {
      return "/article/" + id
    },
    getDisplayDate (d) {
      let date = new Date(d)
      return date.getFullYear() + "年" + (date.getMonth() + 1) + "月" + date.getDate() + "日"
    },
    getCategories (category) {
      console.log(category)
      if (!category) {
        return null
      }
      return category.split(',')
    }
  }
}
</script>

<style lang="scss">
@import "@/styles/_variables.scss";
  .card {
    margin: 4% 0;
    height: auto;
    padding-bottom: 2%;
    box-shadow: 0 4px 8px 0 rgba(40, 40, 40, 0.2);

    &:hover {
      box-shadow: 0 8px 16px 0 rgba(0,0,0,0.2);
      transition: 0.3s;
    }

    a {
      text-decoration: none;
    }
    
    .thumbnail {
      width: 230px;
      height: 152px;
      border: 1px solid $color-border-gray;
    }

    img {
      max-width: 230px;
      max-height: 151px;
    }

    .card-container {
      padding: 16px;
      font-size: 16px;
      color: $color-font-black;

      .title {
        font-size: 16px;
        font-weight: 600;
        margin-bottom: 8px;
        max-width: 200px;
      }

      .date {
        display: flex;
        color: #a8abb1;
        font-size: 12px;

        .published-date {
          margin-top: 0px;
          margin-bottom: 4px;
        }

        time {
          margin-left: 4px;
        }
      }

      .description {
        font-size: 14px;
        max-width: 200px;
        margin-top: 8px;
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

      .editor {
        position: relative;
        display: flex;
        margin-top: 8px;
        color: rgba(0, 0, 0, 0.54);

        .name {
          position: absolute;
          top: 0px;
          line-height: 1;
          font-size: 14px;
        }

        .announcement-date {
          font-size: 12px;
          position: absolute;
          bottom: 0;
          line-height: 1;
          color: rgba(0, 0, 0, 0.54);
        }

        .icon {
          border-radius: 32px;
          max-width: 32px;
          max-height: 32px;
          margin-right: 8px;
        }
      }
    }
  }
</style>