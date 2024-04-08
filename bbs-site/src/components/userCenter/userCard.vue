

<template>
  <div class="profile" :style="{ backgroundImage: 'url(' + backgroundImage + ')' }">
    <div v-if="isOwner" class="file is-light is-small change-bg">
      <label class="file-label">
        <input class="file-input" type="file" @change="uploadBackground" />
        <span class="file-cta">
          <span class="file-icon">
            <i class="iconfont icon-upload" />
          </span>
          <span class="file-label">设置背景</span>
        </span>
      </label>
    </div>
<!--    <avatar :user="localUser" :round="true" size="100" class="profile-avatar" />-->
    <img  class="profile-avatar" :src="avatar" alt="">
    <div class="profile-info">
      <div class="metas">
        <h1 class="nickname">
          吴琳琳
        </h1>
        <div  class="description">
          <p>淡泊名利</p>
        </div>
      </div>
      <div class="action-btns">
        <followBtn />
      </div>
    </div>
  </div>
</template>
<script>
import followBtn from "../FollowBtn.vue";
export default{
  name:'userCard',
  components:{
    followBtn
  },
  data (){
    return {
      isOwner:false,
      backgroundImage:'https://cdn.jsdelivr.net/gh/M-500/picx-images-hosting@master/图片.pf4aqou0t.webp',
      avatar: "https://cdn.jsdelivr.net/gh/M-500/picx-images-hosting@master/425380192_25678980908368088_6576232889963075542_n.8ad1dchumr.webp",
    }
  }
}
</script>
<style scoped>
.profile {
  display: flex;
  margin-bottom: 10px;
  border-top-left-radius: 6px;
  border-top-right-radius: 6px;
  background-size: cover;
  background-position: 50%;
  height: 220px;

// filter: blur(2px) contrast(0.8);

  position: relative;

  .profile-avatar {
    position: absolute;
    top: 90px;
    left: 10px;
    height: 100px;
    border-radius: 50%;
  }

  .change-bg {
    position: absolute;
    top: 10px;
    right: 10px;
    opacity: 0.5;
    &:hover {
      opacity: 1;
    }
  }

  .profile-info {
    display: flex;
    justify-content: space-between;
    align-items: center;
    width: 100%;
    margin-top: 150px;
    padding: 10px 10px 10px 120px;
    background-image: linear-gradient(
      90deg,
      #ffffffff,
      rgba(255, 255, 255, 0.5),
      #dce9f200
    );

    .metas {
      display: flex;
      align-items: flex-start;
      flex-direction: column;
      width: 100%;
      .nickname {
        font-size: 18px;
        font-weight: 700;
        a {
          color: #ffffff;
          &:hover {
            color: #ffffff;
            text-decoration: underline;
          }
        }
      }

      .description {
        font-size: 14px;
        color: #1f2d3d;
      }

      .homepage {
        font-size: 14px;
        a {
          color: var(--color2);
          &:hover {
            color: var(--text-link-color);
            text-decoration: underline;
          }
        }
      }
    }
    .action-btns {
      margin-left: 10px;
    }
  }
}

.dark-mode {
  .profile-info {
    background-image: linear-gradient(
      90deg,
      #00000088,
      rgba(255, 255, 255, 0.5),
      #dce9f200
    );
  }
}
</style>
