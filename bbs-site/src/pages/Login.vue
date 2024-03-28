<script>

import { imageCaptchaAPI } from "@/api/sign-in/login";
export default {
  name: 'PwdLogin',
  data() {
    return {
      captchaUri: "",
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
  created() {
    this.getCaptchaImg();
  },
  methods: {
    getCaptchaImg() {
      imageCaptchaAPI()
        .then((res) => {
          this.form.captcha_id = res.captcha_id;
          this.captchaUri = res.pic_path;
        })
        .catch((e) => {
          console.log(e);
        });
    },
    toRegister() {
      this.$router.push('/sign-up');
    },
    submitForm(form) {
      this.$refs[form].validate(async valid => {
        if (!valid) return this.$message.error('非法输入数据，请重新输入')
        const { data: res } = await this.$http.post('sj/login', this.form)
        if (!res.success) return this.$message.error(res.msg)
        window.localStorage.setItem('token', res.data.token)
        // 将权限数据存到store中
        this.$store.commit('setRightList', res.data.menu_list)
        this.$store.commit('setUsername', res.data.username)
        this.$store.commit('setPhoto', res.data.cover_image_link)
        // 将用户所具备的权限动态添加到路由规则
        initDynamicRouter()
        await this.$router.push('admin/index')
      })
    },
    resetForm(form) {
      this.$refs[form].resetFields()
    }
  }
}
</script>

<template>
  <div class="main_container">
    <div class="loginBox">
      <div class="title">
        <h1>欢迎光临！</h1>
      </div>
      <el-form ref="form" :rules="rules" :model="form" class="loginForm" size="middle">
        <el-form-item prop="username">
          <el-input v-model="form.username" prefix-icon="el-icon-mobile-phone" placeholder="请输入用户名/手机号"></el-input>
        </el-form-item>
        <el-form-item prop="password">
          <el-input v-model="form.password" type="password" prefix-icon="el-icon-lock" placeholder="请输入密码"></el-input>
        </el-form-item>
        <el-form-item prop="captcha_code">
          <div class="captcha">
            <el-input v-model="form.captcha_code" placeholder="图片验证码" prefix-icon="el-icon-document"></el-input>
            <img @click="getCaptchaImg" class="captcha__img" :src="captchaUri" alt="" />
          </div>
        </el-form-item>
        <el-form-item>
          <el-button class="loginBtn" type="primary" @click="submitForm('form')">登录</el-button>
        </el-form-item>
        <el-form-item>
          <el-button class="regBtn" type="primary" @click="toRegister('form')">注册</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<style scoped>
.container {
  height: 100%;
  /*background: url("../assets/992390.jpg") no-repeat center;;*/
  /* background-color: #282c34; */
  /* 假设用flex布局的话*/
  /*display: flex;*/
  /*justify-content: center;*/
  /*align-items: center;*/
}

.loginBox {
  width: 450px;
  height: 450px;
  /*background-color:  #f9f9f9;*/
  position: absolute;
  top: 45%;
  left: 50%;
  transform: translate(-50%, -60%);
  /*border-radius: 20px;*/
  background-color: #ffffff; /* 容器背景颜色 */
  border: 1px solid #dddddd; /* 边框 */
  border-radius: 10px; /* 圆角 */
  box-shadow: 0px 2px 5px rgba(0, 0, 0, 0.1); /* 阴影 */
  padding: 20px; /* 内边距 */
  margin: 20px; /* 外边距 */
}
.title {
  display: flex;
  justify-content: center;
}

.loginForm {
  background-color: rgba($color: #fff, $alpha: 0.1);
  position: absolute;
  bottom: 10%;
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
  background: linear-gradient(to right, #193441, #3e606f);
  background-blend-mode: normal, normal;
}
.regBtn {
  height: 44px;
  width: 100%;
  font-size: 20px;
  font-weight: normal;
  font-stretch: normal;
  letter-spacing: 2px;
  color: #ffffff;
  background: linear-gradient(to right, #d1dbbd, #ecf0f1);
  background-blend-mode: normal, normal;
}

.captcha {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
}

.captcha__img {
  margin-left: 20px;
  width: 150px;
  height: 40px;
  border-bottom: 1px solid #dbdbdb;
}
</style>
