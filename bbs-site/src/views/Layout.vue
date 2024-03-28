<template>
  <el-container class="container">
    <el-header class="container-header">
      <Header></Header>
    </el-header>
    <el-main>
      <router-view></router-view>
    </el-main>
  </el-container>
</template>

<script>

import Header from "../layout/Header.vue";
import Footer from "../layout/Footer.vue";
import Container from "../layout/Container.vue";
export default {
  name: 'Layout',
  components: { Footer, Header, Container },
}
</script>
<style scoped>
.container {
  height: 100%;

  //background-color: #ecf5ff;
}
.el-main {
  padding: 10px;
  background: linear-gradient(to right, #91aa9d, #d1dbbd);
  /* background-color: #91aa9d; */
}
.el-header {
  padding: 0;
  align-items: center;
}
</style>
