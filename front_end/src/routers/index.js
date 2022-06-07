import VueRouter from "vue-router";
import Vue from "vue";
Vue.use(VueRouter);

const routes = [
  {
    path: '/',
    name: 'Home',
    meta: { title: '首页' },
    component: () => import('../components/Home.vue'),
    children: [
      {
        path: '/login',
        name: 'Login',
        meta: { title: '登录' },
        component: () => import('../components/user/Login.vue')
      },
      {
        path: '/register',
        name: 'Register',
        meta: { title: '注册' },
        component: () => import('../components/user/Register.vue')
      }
    ]
  }
];


const router = new VueRouter({
  mode: 'history',
  routes
})

export default router;
