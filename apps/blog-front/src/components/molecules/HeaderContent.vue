<template>
  <div class="content-list">
    <a class="home" v-bind:class="{ active: isHome }" href="#">ホーム</a>
    <a class="about" v-bind:class="{ active: isAbout }" href="#">about</a>
    <a class="profile" v-bind:class="{ active: isProfile }" href="#">profile</a>
  </div>    
</template>

<script>
export default {
  name: 'HeaderContent',
  props: {
    currentPage: String
  },
  data () {
    return {
      isHome: true,
      isAbout: false,
      isProfile: false
    }
  }
}
</script>

<style lang="scss">
  .content-list {
    height: 48px;
    padding-left: 24px;
    display: flex;
    align-items: stretch;
    
    a {
      display: flex;
      align-items: center;
      justify-content: center;
      padding: 0 20px;
      font-size: 16px;
      text-decoration: none;
      font-weight: 600;
      color: white;
      opacity: 0.3;

      &:hover {
        opacity: 1;
      }
    }
    .active {
      opacity: 1;
      border-bottom: solid 3px #408100;
    }
  }
</style>