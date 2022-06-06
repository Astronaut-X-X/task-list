import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  {
    path: '/',
    name: 'Home',
    meta: { title: '首页' },
    component: () => import(/* webpackChunkName: "home" */ '../components/Home.vue')
  },
  {
    path: '/login',
    name: 'Login',
    meta: { title: '登录' },
    component: () => import(/* webpackChunkName: "login" */ '../components/user/Login.vue')
  },
  {
    path: '/register',
    name: 'Register',
    meta: { title: '注册' },
    component: () => import(/* webpackChunkName: "register" */ '../components/user/Register.vue')
  }
  // {
  // path: '/about',
  // name: 'About',
  // route level code-splitting
  // this generates a separate chunk (about.[hash].js) for this route
  // which is lazy-loaded when the route is visited.
  // component: () => import(/* webpackChunkName: "about" */ '../views/About.vue')
  // }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
