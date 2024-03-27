<script setup>
export default {
  name: 'Header',
}
</script>

<template>
  <div class="header">
    <div class="top-bar">
      <div class="logo">
        <img src="https://mlog.club/_nuxt/logo.Wwb-CAFS.png">
      </div>
    </div>
    <div class="show-login">
      <el-button size="mini" icon="el-icon-edit" type="primary">发表</el-button>
      <el-button size="mini" plain>登录</el-button>
    </div>
  </div>
</template>

<style scoped>
.header {
  height: 64px;
  background: #fff;
  padding-right: 120px;
  padding-left: 120px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}
.logo {
  img{
    height: 30px;
    //max-width: 100%;
  }
}
</style>
