<template>
  <div class="card">
    <a :href="getArticleUrl(show.sys.id)">
      <img class="thumbnail" :src="show.thumbnail.url">
      <div class="card-container">
        <div class="title">{{ show.title }}</div>
        <div class="date">
          <p class="published-date">投稿日<time :datetime="show.sys.firstPublishedAt">{{ getDisplayDate(show.sys.firstPublishedAt) }}</time></p>
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
      return date.getFullYear() + "年" + date.getMonth() + "月" + date.getDate() + "日"
    }
  }
}
</script>

<style lang="scss">
@import "@/styles/_variables.scss";
  .card {
    margin: 4% 0;
    height: auto;
    padding-bottom: 4%;
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
      }

      .date {
        display: flex;
        color: #a8abb1;
        font-size: 12px;

        .published-date {
          margin-top: 0px;
          margin-bottom: 8px;
        }

        time {
          margin-left: 4px;
        }
      }

      .description {
        font-size: 14px;
        max-width: 200px;
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