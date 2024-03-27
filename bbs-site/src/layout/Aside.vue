<script setup>
export default {
  name: 'Aside',
}
</script>

<template>
  <div class="nav">

  </div>
</template>

<style scoped>
.nav {
  //height: 100%;
  //background-color: rebeccapurple;
  display: flex;
  justify-content: flex-end;
}
.tag-list{
  padding: 10px 20px 10px;
  width: 60%;
  background-color: #f9f9f9;
}
.tag-box{
  display: flex;
  justify-content: flex-start;
  img{
    width: 20px;
    background-color: white;
  }
  .title{
    margin-left: 8px;
  }
}
</style>
