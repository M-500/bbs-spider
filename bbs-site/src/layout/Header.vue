<script>
export default {
  name: 'Header',
  data() {
    return {

    }
  },
  methods: {
    toLogin() {
      this.$router.push('/login')
    },
    toHome() {
      this.$router.push('/home')
    },
    toEdit() {
      this.$router.push('/article/edit')
    }
  },

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
      <el-button size="mini" icon="el-icon-edit" type="primary" @click="toEdit">发表</el-button>
      <el-button size="mini" plain @click="toLogin">登录</el-button>
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
</style>
