<script>
import curUser from "@/utils/cur-user";
export default {
  name: 'Header',
  data() {
    return {
      username: curUser.getUserName(),
      avatar: curUser.getUserAvatar(),
    }
  },
  methods: {
    toHome() {
      this.$router.push('/home')
       },
    toEdit() {
       this.$router.push( '/article/edit')
    },
    handleCommand(command) {
      if (command === "exitSys") {
        this.logout();
      }
      if (command === "userCenter") {
        this.$router.push({
          name: "userCenter",
        });
      }
    },
    logout(){
      window.sessionStorage.clear()
      window.localStorage.clear();
      // 如果当前界面刷新 那么就会清空vuex的数据
      this.$router.push("/login");
      window.location.reload();
    },
  }
}
</script>

<template>
  <div class="header">
    <div class="top-bar">
      <div class="logo">
        <img :src="require('@/assets/logo.svg')" @click="toHome">
      </div>
    </div>
    <div class="show-login">
      <el-button class="editBtn" size="mini" icon="el-icon-edit" type="primary" @click="toEdit">写文章</el-button>
      <el-dropdown class="userDrop" @command="handleCommand">
        <div class="avatar-warp">
          <!-- <img class="avatar" :src="cover_image_link" alt="" /> -->
          <img class="avatar" :src="avatar" alt="" />
          <span>{{ username }}</span>
          <i class="el-icon-arrow-down el-icon--right"></i>
        </div>
        <el-dropdown-menu>
           <el-dropdown-item command="userCenter">个人中心</el-dropdown-item>
          <el-dropdown-item command="exitSys">用户退出</el-dropdown-item>
        </el-dropdown-menu>
      </el-dropdown>
    </div>

  </div>
</template>

<style scoped>
.header {
  /* height: 44px; */
  height: 100%;
  background: #fff;
  padding-right: 120px;
  padding-left: 150px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}
.logo {
  align-items: center;
  img {
    height: 30px;
    //max-width: 100%;
  }
}
.editBtn {
  background-color: #00bbc9;
  border: #00bbc9 1px solid;
}
.avatar-warp{
  display: flex;
  align-items: center;
  img{
    height: 30px;
    border-radius: 50%;
  }
}
.show-login{
  display: flex;
}
.userDrop{
  margin-left: 20px;
}
.avatar-warp{
  span{
    margin-left: 5px;
  }
}
</style>
