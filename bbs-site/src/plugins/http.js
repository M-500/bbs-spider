/*
 * @Date: 2021-09-13 15:04:44
 * @LastEditors: 王俊
 * @LastEditTime: 2021-12-23 11:47:56
 * @Description: 修改为本地接口对接
 */
import Vue from 'vue'
import axios from 'axios'
import { Message } from 'element-ui'

const Url = process.env.VUE_APP_BASE_SERVE_URL
axios.defaults.baseURL = Url

axios.interceptors.request.use(config => {
  config.headers.Authorization = `Bearer ${window.localStorage.getItem('token')}`
  return config
})
axios.interceptors.response.use(response => Promise.resolve(response),
  error => {
    if (error && error.response && error.response.status) {
      switch (error.response.status) {
        case 500:
          Message.error(error.response.data.msg)
          break
        case 404:
          Message.error(error.response.data.msg)
          break
        case 403:
          Message.error(error.response.data.msg)
          // do something...
          break
        default:
          // do something...
          break
      }
      switch (error.response.data.code) {
        case '9003':
          // console.log('下水道', error.response.data.code)
          window.localStorage.clear()
          this.$router.refresh()
          break
        case '9002':
          // console.log('下水道', error.response.data.code)
          window.localStorage.clear()
          this.$router.refresh()
          break
      }
      throw error.response.data
    }
  })
Vue.prototype.$http = axios

export { Url }
