<template>
  <div class="article-page">
    <MainHeader page="article" />
    <ArticleText :article="article" />
    <RelatedArticleList :shows="relatedArticles" />
  </div>
</template>

<script>
import katex from 'katex'
import MainHeader from '@/components/organisms/MainHeader.vue'
import ArticleText from '@/components/organisms/ArticleText.vue'
import RelatedArticleList from '@/components/organisms/RelatedArticleList.vue'
import { MARKS, BLOCKS } from '@contentful/rich-text-types'
import { documentToHtmlString } from '@contentful/rich-text-html-renderer'
export default {
  name: 'Article',
  components: {
    MainHeader,
    ArticleText,
    RelatedArticleList
  },
  data() {
    return {
      article: {},
      relatedArticles: []
    }
  },
  async created() {
    let item = await this.getArticle(this.$route.params.id);
    const options = this.getRenderOptions(item.content.links)
    this.article = {
      title: item.title,
      thumbnail: item.thumbnail.url,
      publishedAt: item.sys.publishedAt,
      firstPublishedAt: item.sys.firstPublishedAt,
      description: item.description,
      category: this.getCategories(item.category),
      content: documentToHtmlString(item.content.json, options)
    }
    this.relatedArticles = await this.getRelatedArticle(this.article.category)
  },
  methods: {
    getCategories (category) {
      if (!category) {
        return null
      }
      return category.split(',')
    },
    getRenderOptions(links) {
      const assetMap = new Map()

      for (const asset of links.assets.block) {
        assetMap.set(asset.sys.id, asset)
      }

      return {
        renderMark: {
          [MARKS.CODE]: text => {
            const doc = new DOMParser().parseFromString(text, 'text/html')
            if (doc.documentElement.textContent.startsWith('\\')) {
              console.log(katex.renderToString(doc.documentElement.textContent))
              return `<pre>${katex.renderToString(doc.documentElement.textContent)}</pre>`
            }
            return `<pre>${text}</pre>`
          }
        },
        renderNode: {
          [BLOCKS.EMBEDDED_ASSET]: (node) => `<img src=${assetMap.get(node.data.target.sys.id).url} alt=${assetMap.get(node.data.target.sys.id).fileName} />`
        }
      }
    },
    getArticle: async (id) => {
      const query = `query {
        articles(id: "${id}") {
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
            url
          }
          content {
            json
            links {
              assets {
                hyperlink {
                  sys {
                    id
                  }
                  contentType
                  fileName
                  width
                  height
                  url
                }
                block {
                  sys {
                    id
                  }
                  contentType
                  fileName
                  width
                  height
                  url
                }
              }
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
        return response.data.articles;
      } catch (error) {
        throw new Error("Could not receive the data from Contentful!");
      }
    },
    getRelatedArticle: async (category) => {
      const query = `query {
        articlesCollection(where: { category_in: "${category[0]}" }, limit: 20) {
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
      }`
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
