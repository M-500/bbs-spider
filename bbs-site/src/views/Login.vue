<template>
  <div class="container">

    <div class="loginBox">
      <h1 class="title">欢迎光临数鲸数字！</h1>
      <el-form ref="form" :rules="rules" :model="form" class="loginForm" size="middle">
        <el-form-item prop="user_name">
          <el-input v-model="form.username" prefix-icon="el-icon-mobile-phone" placeholder="请输入用户名/手机号"></el-input>
        </el-form-item>
        <el-form-item prop="password">
          <el-input v-model="form.password" type="password" prefix-icon="el-icon-lock" placeholder="请输入密码"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button class="loginBtn" type="primary" @click="submitForm('form')">登录</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
import { initDynamicRouter } from '../router'

export default {
  name: 'Login',
  data() {
    return {
      form: {
        username: '',
        password: ''
      },
      rules: {
        username: [
          {
            required: true,
            message: '请输入用户名',
            trigger: 'blur'
          },
          {
            min: 3,
            max: 15,
            message: '用户名长度不合法',
            trigger: 'blur'
          }
        ],
        password: [
          {
            required: true,
            message: '请输入密码',
            trigger: 'blur'
          },
          {
            min: 6,
            max: 15,
            message: '密码长度不合法',
            trigger: 'blur'
          }
        ]
      }
    }
  },
  methods: {
    submitForm(form) {
      this.$refs[form].validate(async valid => {
        if (!valid) return this.$message.error('非法输入数据，请重新输入')
        const { data: res } = await this.$http.post('sj/login', this.form)
        if (!res.success) return this.$message.error(res.msg)
        // window.localStorage.setItem('token', res.data.token)
        // // 将权限数据存到store中
        // this.$store.commit('setRightList', res.data.menu_list)
        // this.$store.commit('setUsername', res.data.username)
        // this.$store.commit('setPhoto', res.data.cover_image_link)
        // // 将用户所具备的权限动态添加到路由规则
        // initDynamicRouter()
        await this.$router.push('admin/index')
      })
    },
    resetForm(form) {
      this.$refs[form].resetFields()
    }
  }
}
</script>

<style scoped>
.container {
  height: 100%;
  /*background: url("../assets/992390.jpg") no-repeat center;;*/
  background: #01020b url('../assets/login/bj.jpg') no-repeat center/cover;
  background-size: cover;
  /*background-color: #282c34;*/
  /* 假设用flex布局的话*/
  /*display: flex;*/
  /*justify-content: center;*/
  /*align-items: center;*/
}

.title {
  color: #f9f9f9;
  font-size: 46px;
  font-weight: bold;
  margin: 0 auto;
  text-align: center;
}

.loginBox {
  width: 450px;
  height: 500px;
  /*background-color:  #f9f9f9;*/
  position: absolute;
  top: 55%;
  left: 50%;
  transform: translate(-50%, -60%);
  /*border-radius: 20px;*/
}

.loginForm {
  background-color: rgba($color: #fff, $alpha: 0.1);
  width: 100%;
  position: absolute;
  bottom: 15%;
  padding: 0 25px;
  box-sizing: border-box;
}

.loginBtn {
  height: 44px;
  width: 100%;
  font-size: 20px;
  font-weight: normal;
  font-stretch: normal;
  letter-spacing: 2px;
  color: #ffffff;
  background-image: linear-gradient(0deg, #0176e4 0%, #00b8ff 100%),
    linear-gradient(#00b8ff, #00b8ff);
  background-blend-mode: normal, normal;
}
</style>
