<template>
  <div class="article">
    <h1>{{ item.title }}</h1>
    <p>{{ item.description }}</p>
    <p>{{ item.content }}</p>
  </div>
</template>

<script>
export default {
  name: 'Article',
  data() {
    return {
      item: {}
    }
  },
  async created() {
   this.item = await this.getArticle(this.$route.params.id);
   console.log(this.item)
 },
  methods: {
    getArticle: async (id) => {
      const query = `query {
        articles(id: "${id}") {
          title
          description
          thumbnail {
            title
            url
          }
          content {
            json
          }
        }
      }`;
      console.log(query)
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
        return response.data.articles;
      } catch (error) {
        throw new Error("Could not receive the data from Contentful!");
      }
    }
  }
}
</script>