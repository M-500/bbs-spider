<template>
  <el-container class="container">
    <el-header>
      <Header></Header>
    </el-header>
    <el-main>
      <router-view></router-view>
    </el-main>
  </el-container>
</template>

<script>

import Header from "../layout/Header.vue";
import Aside from "../layout/Aside.vue";
import Footer from "../layout/Footer.vue";
import Container from "../layout/Container.vue";
export default {
  name: 'Layout',
  components: {Footer, Aside, Header,Container},
}
</script>
<style scoped>
.container {
  height: 100%;

  //background-color: #ecf5ff;
}
.el-main{
  padding: 0;
}
.el-header{
  padding: 0;
}

</style>
