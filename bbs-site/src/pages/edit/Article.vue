<script>
import mavonEditor from 'mavon-editor'
import 'mavon-editor/dist/css/index.css'
import { editArticleAPI } from '@/api/article/edit'
export default {
  name: 'ArticleEdit',
  data() {
    return {
      form: {
        id: 0,
        content_type: 'article',
        title: '',
        summary: '',
        cover: '',
        content: ''
      }
    }
  },
  methods: {
    handleAvatarSuccess(res, file) {
      this.form.cover = URL.createObjectURL(file.raw);
    },
    beforeAvatarUpload(file) {
      const isJPG = file.type === 'image/jpeg';
      const isLt2M = file.size / 1024 / 1024 < 2;

      if (!isJPG) {
        this.$message.error('上传头像图片只能是 JPG 格式!');
      }
      if (!isLt2M) {
        this.$message.error('上传头像图片大小不能超过 2MB!');
      }
      return isJPG && isLt2M;
    },
    onSubmit() {
      editArticleAPI(this.form).then((res) => {

      }).catch((e) => {
        this.$message({
          message: e.msg,
          type: "error",
        });
      })
      console.log(this.form)
    },
  }
}
</script>

<template>
  <div class="main-container">
    <div class="edit-container">
      <div class="title">发文章</div>
      <el-form ref="form" :model="form" label-width="80px">
        <el-form-item label="文章标题">
          <el-input v-model="form.title" :autosize="{ minRows: 1, maxRows: 2}"></el-input>
        </el-form-item>
        <el-form-item label="文章摘要">
          <el-input type="textarea" v-model="form.summary" :autosize="{ minRows: 2, maxRows: 4}"></el-input>
        </el-form-item>
        <el-form-item label="文章内容">
          <mavon-editor v-model="form.content" />
          <!-- <el-input type="textarea" v-model="form.desc" :autosize="{ minRows: 4, maxRows: 10}"></el-input> -->
        </el-form-item>
        <el-form-item label="文章图片">
          <el-upload class="avatar-uploader" action="https://jsonplaceholder.typicode.com/posts/" :show-file-list="false" :on-success="handleAvatarSuccess" :before-upload="beforeAvatarUpload">
            <img v-if="form.cover" :src="form.cover" class="avatar">
            <i v-else class="el-icon-plus avatar-uploader-icon"></i>
          </el-upload>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="onSubmit">发布</el-button>
          <el-button>取消</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<style scoped>
.main-container {
  width: 100%;
  display: flex;
  justify-content: center;
}
.edit-container {
  width: 80%;
  margin: 10px 30px 10px;
  padding: 10px 30px 10px;
  border-radius: 5px;
  background-color: white;
}
.title {
  margin: 10px 20px 30px;
  font-weight: 700;
  font-size: 27px;
}

.avatar-uploader .el-upload {
  border: 1px dashed #d9d9d9;
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
}
.avatar-uploader .el-upload:hover {
  border-color: #409eff;
}
.avatar-uploader-icon {
  font-size: 28px;
  color: #8c939d;
  width: 178px;
  height: 178px;
  line-height: 178px;
  text-align: center;
}
.avatar {
  width: 178px;
  height: 178px;
  display: block;
}
.el-icon-plus {
  border: #d9d9d9 1px solid;
}
</style>
