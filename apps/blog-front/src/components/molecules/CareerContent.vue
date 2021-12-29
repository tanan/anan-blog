<template>
  <div class="career-content">
      <div class="period">{{ career.period }}</div>
      <!-- <div class="icon"><OfficeBuilding :size="32" class="icon-content" /></div> -->
      <div class="company">
        <a :href="career.company.url"><div class="name">{{ career.company.name }}</div></a>
        <div class="role">{{ career.role }}</div>
        <div class="description">{{ career.description }}</div>
        <div class="technology">{{ career.technology }}</div>
      </div>
  </div>
</template>

<script>
// import OfficeBuilding from 'vue-material-design-icons/OfficeBuilding.vue';
export default {
  name: 'CareerContent',
  // components: {
  //   OfficeBuilding
  // },
  props: {
    career: Object
  }
}
</script>

<style lang="scss">
  .career-content {
    margin-top: 48px;
    display: flex;
    justify-content: center;
    
    .period {
      font-size: 0.8em;
      width: 12%;
      padding-right: 2%;
      // padding-top: 2%;
    }

    // .icon {
    //   position: relative;
    //   background-color: #2c3e50;
    //   opacity: 0.8;
    //   width: 5%;
    //   // height: 50%;
    //   border-radius: 50%;
    //   margin-top: 20px;
    //   margin-right: 24px;

    //   .icon-content {
    //     position: absolute;
    //     top: 14px;
    //     left: 14px;
    //     color: #fff;
    //   }
    // }
    
    .company {
      flex-wrap: wrap;
      width: 68%;

      a:link, a:visited, a:hover, a:active {
        color: inherit;
        text-decoration: none;
      }
      .name {
        font-size: 1.3em;
        font-weight: 800;
      }
      .role {
        padding-top: 8px;
        font-weight: 600;
      }
      .description {
        padding-top: 16px;
      }
      .technology {
        font-size: 0.8em;
        padding-top: 8px;
      }
    }
  }
</style>