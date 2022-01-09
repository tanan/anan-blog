<template>
  <div class="article-page">
    <MainHeader page="about" />
    <ArticleText :article="article" />
  </div>
</template>

<script>
import MainHeader from '@/components/organisms/MainHeader.vue'
import ArticleText from '@/components/organisms/ArticleText.vue'
import { MARKS, BLOCKS } from '@contentful/rich-text-types'
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
    console.log(item)
    const options = this.getRenderOptions(item.content.links)
    this.article = {
      title: item.title,
      thumbnail: item.thumbnail.url,
      publishedAt: item.sys.publishedAt,
      firstPublishedAt: item.sys.firstPublishedAt,
      description: item.description,
      content: documentToHtmlString(item.content.json, options)
    }
  },
  methods: {
    getRenderOptions(links) {
      const assetMap = new Map()

      for (const asset of links.assets.block) {
        assetMap.set(asset.sys.id, asset)
      }

      return {
        renderMark: {
          [MARKS.CODE]: text => `<pre>${text}</pre>`
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
      }`
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