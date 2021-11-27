<template>
  <div class="home">
    <MainLogo />
    <HelloWorld msg="Welcome to Your Vue.js App"/>
    <div class="card-list wrapper">
      <div class="cards-container">
        <div v-for="show in shows" :key="show.title">
          <a :href="getArticleUrl(show.sys.id)">
            <div class="card">
              <img class="thumbnail" :src="show.thumbnail.url">
              <div class="title">{{ show.title }}</div>
              <div class="description">{{ show.description }}</div>
            </div>
          </a>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
// @ is an alias to /src
import HelloWorld from '@/components/HelloWorld.vue'
import MainLogo from '@/components/atoms/MainLogo.vue'

export default {
  name: 'Home',
  components: {
    HelloWorld,
    MainLogo
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
   getArticleUrl (id) {
     return "/articles/" + id
   },
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
@import "../styles/_variables.scss";
  .card {

  width: 256px;
  height: 320px;
  box-shadow: 0 4px 8px 0 rgba(40, 40, 40, 0.2);
  margin: 32px 32px;
  margin-left: 0px;

  &:hover {
    box-shadow: 0 8px 16px 0 rgba(0,0,0,0.2);
    transition: 0.3s;
  }
  
  .thumbnail {
    width: 100%;
    height: 152px;
    border: 1px solid $color-border-gray;
  }

  img {
    width: 100%;
    max-height: 151px;
  }

  .card-container {
    padding: 16px;
    font-size: 16px;
    color: $color-font-black;

    .title {
      font-size: 16px;
      margin-bottom: 8px;
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

    .description {
      margin-top: 8px;
      font-size: 14px;
    }
  }
}
</style>