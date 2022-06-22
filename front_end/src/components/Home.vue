<template>
  <div>
    <header-vue></header-vue>
    <main class="px-16 py-6">
      <transition name="el-fade-in-linear" mode="out-in">
        <router-view></router-view>
      </transition>
    </main>
  </div>
</template>

<script>
import { getUserByIDInContext } from '../api/user';
import HeaderVue from './Header.vue';
export default {
  name: 'Home',
  data () {
    return {
      isLogin: false
    };
  },
  components: {
    HeaderVue
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

    // 

    logout () {
      this.$storage.clearItem('token');
      this.$router.push('/');
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