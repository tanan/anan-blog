<template>
  <div class="article-section">
    <h2>記事一覧</h2>
    <div class="card-list-container">
      <Card v-for="show in shows" :key="show.title" :show="show" />
    </div>
  </div>    
</template>

<script>
import Card from '@/components/molecules/Card.vue'
export default {
  name: 'ArticleList',
  components: {
    Card
  },
  data() {
    return {
      shows: []
    };
  },
  async created() {
   this.shows = await this.getShows();
  },
  methods: {
    getShows: async () => {
      const query = `{
        articlesCollection(limit: 20) {
          items {
            sys {
              id
              publishedAt
              firstPublishedAt
            }
            title
            description
            category
            thumbnail {
              title
              fileName
              url
            }
            content {
              json
            }
          }
        }
      }`;
      const fetchUrl = `https://graphql.contentful.com/content/v1/spaces/${process.env.VUE_APP_CONTENTFUL_SPACE_ID}`;
      const fetchOptions = {
        method: "POST",
        headers: {
          Authorization: `Bearer ${process.env.VUE_APP_CONTENTFUL_ACCESS_TOKEN}`,
          "Content-Type": "application/json"
        },
        body: JSON.stringify({ query })
      };

      try {
        const response = await fetch(fetchUrl, fetchOptions).then(response =>
          response.json()
        );
        return response.data.articlesCollection.items;
      } catch (error) {
        throw new Error("Could not receive the data from Contentful!");
      }
    }
  }
}
</script>

<style lang="scss">
@import "@/styles/_variables.scss";
  .article-section {
    width: 58%;
    margin-right: auto;
    margin-left: auto;

    h2 {
      font-size: 1.8em;
      text-align: center;
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