<template>
  <div class="container flex justify-center content-center">
    <el-form
      :model="loginForm"
      :rules="rules"
      ref="loginForm"
      class="bg-blue-100 pt-16 pb-10 px-16 mt-24 rounded-2xl border-2 border-blue-50 shadow-sm shadow-blue-200 hover:shadow-md duration-800"
    >
      <el-form-item prop="username">
        <el-input
        prefix-icon="el-icon-user"
          type="text"
          v-model="loginForm.username"
          clearable
        ></el-input>
      </el-form-item>
      <el-form-item prop="password">
        <el-input prefix-icon="el-icon-view" type="password" v-model="loginForm.password" clearable></el-input>
      </el-form-item>
      <el-form-item>
        <div class="flex justify-center">
          <el-button type="primary" @click="submitForm('loginForm')"
            >登录</el-button
          >
          <el-button @click="resetForm('loginForm')">重置</el-button>
        </div>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
export default {
  name: "Login",
  data() {
    return {
      loginForm: {
        username: "",
        password: "",
      },
      rules: {
        username: [
          { required: true, message: "请输入用户名", trigger: "blur" },
          {
            min: 6,
            max: 20,
            message: "长度在 6 到 20 个字符",
            trigger: "blur",
          },
        ],
        password: [
          { required: true, message: "请输入密码", trigger: "blur" },
          {
            min: 6,
            max: 20,
            message: "长度在 6 到 20 个字符",
            trigger: "blur",
          },
        ],
      },
    };
  },
  methods: {
    submitForm(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          alert("submit!");
        } else {
          console.log("error submit!!");
          return false;
        }
      });
    },
    resetForm(formName) {
      this.$refs[formName].resetFields();
    },
  },
};
</script>

<style>
</style>