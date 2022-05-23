<template>
  <div class="container flex justify-center">
    <el-form ref="form" :model="form" class="mt-48 w-96 bg-blue-300 pt-16 pb-14 px-14 shadow-md hover:shadow-lg">
      <el-form-item >
        <el-input v-model="form.username"></el-input>
      </el-form-item>
      <el-form-item class="">
        <el-input v-model="form.password"></el-input>
        <a href="#">忘记密码</a>
      </el-form-item>
      <el-form-item>
        <el-button>登录</el-button>
        <el-button @click="register">注册</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import router from "../../router";

export default {
  data() {
    return {
      form: {},
    };
  },
  methods: {
    register() {
      router.push("/register");
    },
  },
};
</script>

<style lang="scss" scoped>
</style>