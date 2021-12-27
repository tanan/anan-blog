<template>
  <div class="article-page">
    <MainHeader page="article" />
    <ArticleText :article="article" />
  </div>
</template>

<script>
import MainHeader from '@/components/organisms/MainHeader.vue'
import ArticleText from '@/components/organisms/ArticleText.vue'
import { MARKS } from '@contentful/rich-text-types';
import { documentToHtmlString } from '@contentful/rich-text-html-renderer'
export default {
  name: 'Article',
  components: {
    MainHeader,
    ArticleText
  },
  data() {
    return {
      article: {},
    }
  },
  async created() {
    let item = await this.getArticle(this.$route.params.id);
    console.log(item.content.json)
    const options = {
      renderMark: {
        [MARKS.CODE]: text => `<pre>${text}</pre>`
      }
    }
    this.article = {
      title: item.title,
      thumbnail: item.thumbnail.url,
      description: item.description,
      content: documentToHtmlString(item.content.json, options)
    }
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