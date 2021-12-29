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

<style scoped lang="scss">
  .personal-content {
    display: flex;
    justify-content: center;
    margin-top: 24px;
    .icon {
      width: 8%;

      img {
        border-radius: 50%;
        width:100%;
        max-width: 100%;
        height: auto;
      }
    }

    .content {
      padding-left: 2%;
      padding-bottom: 1%;
      padding-right: 2%;
      margin-top: auto;
      font-size: 1.2em;
      
      .name {
        font-weight: 600;

        .english {
          padding-left: 12px;
        }
      }

      .link {
        padding-top: 4px;
        padding-left: 2px;
        font-size: 0.7em;
        font-weight: 500;
      }
    }
  }
  
</style>