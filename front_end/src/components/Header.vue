<template>
  <header class="py-8 px-10">
    <nav class="flex justify-between">
      <div class="flex content-end">
        <span class="text-3xl mr-8 divide">TaskList</span>
        <a class="mx-4 p-2 font-light tracking-wide hover:text-blue-500 transition-all" v-if="isLogin"
          @click="today">今日</a>
        <a class="mx-4 p-2 font-light tracking-wide hover:text-blue-500 transition-all" v-if="isLogin"
          @click="calendar">日历</a>
        <a class="mx-4 p-2 font-light tracking-wide hover:text-blue-500 transition-all" v-if="isLogin"
          @click="manage">管理</a>
      </div>
      <div class="flex content-end" v-if="!isLogin">
        <a class="mx-4 p-2 font-light tracking-wide hover:text-blue-500 transition-all" @click="login">登录</a>
        <a class="mx-4 p-2 font-light tracking-wide hover:text-blue-500 transition-all" @click="register">注册</a>
      </div>
      <div class="flex content-end" v-else>
        <!-- <a class="mx-4 p-2 font-light tracking-wide hover:text-blue-500 transition-all" @click="task">计划</a> -->
        <a class="mx-4 p-2 font-light tracking-wide hover:text-blue-500 transition-all" @click="user">用户</a>
        <a class="mx-4 p-2 font-light tracking-wide hover:text-blue-500 transition-all" @click="logout">退出</a>
        <!-- <img src="" alt=""> -->
      </div>
    </nav>
  </header>
</template>

<script>
import { getUserByIDInContext } from '../api/user';
export default {
  name: 'Header',
  data () {
    return {
      isLogin: false
    };
  },
  mounted: function () {
    this.checkIsLogin();
  },
  methods: {
    // 坚持用户登录情况
    checkIsLogin () {
      let token = this.$storage.getItem('token');
      if (token && token.length > 0) {
        getUserByIDInContext({}).then((response) => {
          this.$storage.setItem('user', response);
        }).catch((error) => {
          console.log(error);
        })
        this.isLogin = true;
      } else {
        this.isLogin = false;
      }
    },
    manage () {
      this.$router.push('/manage');
    },
    calendar () {
      this.$router.push('/calendar');
    },
    today () {
      this.$router.push('/today');
    },
    logout () {
      this.$storage.clearItem('token');
      this.$router.go(0);
    },
    task () {
      this.$router.push('/task');
    },
    user () {
      this.$router.push('/user');
    },
    login () {
      this.$router.push('/login');
    },
    register () {
      this.$router.push('/register');
    }
  }
}
</script>

<style>
</style>