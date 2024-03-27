<script setup>
export default {
  name: 'ArticleEdit',
}
</script>

<template>
  <div class="main-container">
    <div class="edit-container">
      <div class="title">发文章</div>
      <div class="edit-form">

      </div>

      <div class="edit-btn">
        <el-button type="success">发布</el-button>
      </div>
    </div>
  </div>
</template>

<style scoped>
.main-container{
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: center;
  background-color: #282c34;
}
.edit-container{
  width: 80%;
  margin: 10px 30px 10px;
  padding: 10px 30px 10px;
  background-color: white;

}
</style>
