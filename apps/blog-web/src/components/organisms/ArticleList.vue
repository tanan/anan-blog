<template>
  <div class="article-section">
    <h2>新着記事</h2>
    <div class="card-list-container">
      <div v-for="show in shows" :key="show.title">
        <Card :show="show" />
      </div>
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
   console.log(this.shows)
  },
  methods: {
    getShows: async () => {
      const query = `{
        articlesCollection(limit: 20) {
          items {
            sys {
              id
            }
            title
            description
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
    width: 56%;
    margin-left: 8%;
  }
  .card-list-container {
    // background-color: #dfdfdf;
    display: flex;
    flex-wrap: wrap;
    padding: 0 12px;
  }
</style>