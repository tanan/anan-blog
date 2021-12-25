<template>
  <div class="content-list">
    <a href="#">ホーム</a>
    <a href="#">about</a>
    <a href="#">profile</a>
  </div>    
</template>

<style lang="scss">
  .content-list {
    height: 48px;
    padding-left: 24px;
    display: flex;
    align-items: center;
    
    a {
      width: 76px;
      padding: 8px 16px;
      display: block;
      font-size: 16px;
      text-align: center;
      text-decoration: none;
      font-weight: 600;
      color: white;
      opacity: 0.3;

      &:hover {
        opacity: 1;
      }
    }
  }
</style>