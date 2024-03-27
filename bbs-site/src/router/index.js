import Vue from 'vue'
import Router from 'vue-router'
import Layout from '@/views/Layout'
import home from "../pages/home.vue";
import Login from "../pages/Login.vue";
import Article from "../layout/edit/Article.vue";

Vue.use(Router)

export default new Router({
  routes: [
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
          path: '/login',
          name: 'Login',
          component: Login
        }
      ]
    },

  ]
})
