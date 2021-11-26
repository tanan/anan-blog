<template>
  <div class="home">
    <MainLogo />
    <HelloWorld msg="Welcome to Your Vue.js App"/>
    <ul class="shows">
    <li
      v-for="show in shows"
      v-bind:key="show.title"
      class="show"
    > 
     {{ show.title }}
    </li>
  </ul>
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
 },
 methods: {
   getShows: async () => {
     const query = `{
       articlesCollection(limit: 20) {
         items {
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
