<template>
  <div class="article-section">
    <h2>新着記事</h2>
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
  mounted() {
    var div = document.getElementsByClassName("card-list-container")
    var card = document.createElement("div")
    card.setAttribute("class", "card empty")
    div[0].appendChild(card)
    var card2 = document.createElement("div")
    card2.setAttribute("class", "card empty")
    div[0].appendChild(card2)
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
              publishedAt
              firstPublishedAt
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

    .empty {
      height: 0;
      padding-top: 0;
      padding-bottom: 0;
      margin-top: 0;
      margin-bottom: 0;
    }
  }
</style>