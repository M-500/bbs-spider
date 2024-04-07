import Vue from 'vue'
import VueRouter from 'vue-router'
import Layout from '@/views/Layout'
import home from "../pages/home.vue";
import Login from "../pages/Login.vue";
import Article from "../pages/edit/Article.vue";
import Register from "../pages/Register.vue"
import Detail from "../pages/Detail.vue"
import curUser from '@/utils/cur-user'

Vue.use(VueRouter)



const routes = [
  {
    path: '/',
    name: 'Layout',
    component: Layout,
    children: [
      {
        path: '/home',
        name: 'home',
        component: home,
      },
      {
        path: '/article/edit',
        name: 'Article',
        component: Article,
      },

      {
        path: '/detail',
        name: 'Detail',
        component: Detail
      }
    ]
  },
  {
    path: '/login',
    name: 'Login',
    component: Login
  },
  {
    path: '/sign-up',
    name: 'Register',
    component: Register
  },

]

const router = new VueRouter({
  mode: 'history',
  // base: process.env.BASE_URL,
  routes
})

const whiteListRouter = [
  "/login",
  "/sign-up"
]
/**
 * 路由守卫
 * */
router.beforeEach((to, from, next) => {
  /** 判断token是否存在*/
  const token = curUser.getToken()
  console.log("是不是没有登录吗", token)
  /** 页面是否需要登陆才能操作？ */
  if (!token && !whiteListRouter.includes(to.path)) {
    next('/login')
  } else {
    next()
  }
})

export default router