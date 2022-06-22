import VueRouter from "vue-router";
import Vue from "vue";
import storage from '../utils/storage'

const originalPush = VueRouter.prototype.push
VueRouter.prototype.push = function push (location, onResolve, onReject) {
  if (onResolve || onReject) {
    return originalPush.call(this, location, onResolve, onReject)
  }
  return originalPush.call(this, location).catch(err => err)
}

Vue.use(VueRouter);

const routes = [
  {
    path: '/',
    name: 'Home',
    meta: { title: '首页' },
    component: () => import('../components/Home.vue'),
    redirect: { name: 'Index' },
    children: [
      {
        path: '/index',
        name: 'Index',
        meta: { title: '首页' },
        component: () => import('../components/Index.vue')
      },
      {
        path: '/user',
        name: 'User',
        meta: { title: '用户' },
        component: () => import('../components/user/User.vue')
      },
      {
        path: '/task',
        name: 'Task',
        meta: { title: '用户' },
        component: () => import('../components/task/Task.vue')
      },
      {
        path: '/today',
        name: 'Today',
        meta: { title: '今日' },
        component: () => import('../components/today/Today.vue')
      },
      {
        path: '/calendar',
        name: 'Calendar',
        meta: { title: '日历' },
        component: () => import('../components/calendar/Calendar.vue')
      },
      {
        path: '/manage',
        name: 'Manage',
        meta: { title: '管理' },
        component: () => import('../components/manage/Manage.vue'),
        redirect: { name: 'Daily' },
        children: [
          {
            path: 'daily',
            name: 'Daily',
            meta: { title: '日计划' },
            component: () => import('../components/manage/Daily.vue'),
          },
          {
            path: 'todo',
            name: 'Todo',
            meta: { title: '待办事项' },
            component: () => import('../components/manage/Todo.vue'),
          },
          {
            path: 'goal',
            name: 'Goal',
            meta: { title: '待办事项' },
            component: () => import('../components/manage/Goal.vue'),
          },
          {
            path: 'quotes',
            name: 'Quotes',
            meta: { title: '待办事项' },
            component: () => import('../components/manage/Quotes.vue'),
          },
          {
            path: 'happy',
            name: 'Happy',
            meta: { title: '待办事项' },
            component: () => import('../components/manage/Happy.vue'),
          }
        ]
      },
    ],
    beforeEnter: isNotLogged
  },
  {
    path: '/login',
    name: 'Login',
    meta: { title: '登录' },
    component: () => import('../components/user/Login.vue'),
    beforeEnter: isLogged
  },
  {
    path: '/register',
    name: 'Register',
    meta: { title: '注册' },
    component: () => import('../components/user/Register.vue'),
    beforeEnter: isLogged
  },
  {
    path: '/forget',
    name: 'Forget',
    meta: { title: '忘记密码' },
    component: () => import('../components/user/Forget.vue'),
    beforeEnter: isLogged
  }
];

const router = new VueRouter({
  mode: 'history',
  routes
})

function isLogged (to, from, next) {
  let token = storage.getItem('token');
  if (token && token.length > 0) {
    next({ name: 'Home' });
  } else {
    next();
  }
}

function isNotLogged (to, from, next) {
  if (to.name === 'Index') {
    next();
  } else {
    let token = storage.getItem('token');
    if (token && token.length > 0) {
      next();
    } else {
      next({ name: 'Login' });
    }
    next();
  }
}



export default router;
