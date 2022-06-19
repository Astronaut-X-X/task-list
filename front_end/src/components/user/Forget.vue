<template>
  <div>
    <header class="py-8 px-10">
      <nav class="flex justify-between">
        <div class="flex content-end">
          <span class="text-3xl mr-8 divide" @click="home">TaskList</span>
        </div>
        <div class="flex content-end">
          <a class="mx-4 p-2 tracking-wide hover:text-blue-500 transition-all" @click="login">登录</a>
          <a class="mx-4 p-2 tracking-wide hover:text-blue-500 transition-all" @click="register">注册</a>
        </div>
      </nav>
    </header>
    <main class="flex justify-center">
      <div class="h-full w-72 mt-20">
        <el-form :model="loginForm" :rules="rules" ref="loginForm">
          <el-form-item prop="username">
            <el-input prefix-icon="el-icon-user" type="text" v-model="loginForm.username" clearable></el-input>
          </el-form-item>
          <el-form-item>
            <div>
              <el-button class="w-full" type="primary" @click="submitForm('loginForm')">修改</el-button>
            </div>
          </el-form-item>
        </el-form>
      </div>
    </main>
  </div>
</template>

<script>
export default {
  name: "Login",
  data () {
    return {
      loginForm: {
        username: "",
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
        ]
      },
    };
  },
  methods: {
    submitForm (formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          this.login();
        } else {
          this.$message({
            message: "error submit!!",
            center: true,
          });
          return false;
        }
      });
    },
    home () {
      this.$router.push("/");
    },
    login () {
      this.$router.push("/login");
    },
    register () {
      this.$router.push('/register');
    },
    forget () {
      this.$router.push('/forget');
    }
  },
};
</script>

<style>
</style>