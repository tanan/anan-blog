<template>
  <div class="personal-content">
      <div class="icon"><img src="/profile.jpeg" alt=""></div>
      <div class="content">
        <div class="name">阿南肇史<span class="english">(Anan Toshifumi)</span></div>
        <div class="link github">github: <a href="https://github.com/tanan">https://github.com/tanan</a></div>
        <div class="link qiita">Qiita: <a href="https://qiita.com/tanan">https://qiita.com/tanan</a></div>
      </div>
  </div>
</template>

<style lang="scss">
  .personal-content {
    display: flex;
    margin-top: 24px;
    .icon {
      padding-left: 16%;

      img {
        width: 96px;
        height: 96px;
        border-radius: 50%;
      }
    }

    .content {
      padding-left: 48px;
      margin-top: auto;
      
      .name {
        font-size: 18px;
        font-weight: 600;

        .english {
          padding-left: 12px;
        }
      }

      .link {
        padding-top: 4px;
        font-weight: 500;
      }
    }
  }
  
</style>