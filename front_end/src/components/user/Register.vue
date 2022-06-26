<template>
  <div>
    <header class="py-8 px-10">
      <nav class="flex justify-between">
        <div class="flex content-end">
          <span class="text-3xl mr-8 divide font-extrabold" @click="home">TaskList</span>
        </div>
        <div class="flex content-end">
          <a class="mx-4 p-2 tracking-wide hover:text-blue-500 transition-all" @click="login">登录</a>
        </div>
      </nav>
    </header>
    <main class="flex justify-center">
      <div class="h-full w-72 mt-20">
        <el-form :model="registerForm" :rules="rules" ref="registerForm">
          <el-form-item prop="username">
            <el-input prefix-icon="el-icon-user" type="text" v-model="registerForm.username" clearable></el-input>
          </el-form-item>
          <el-form-item prop="password">
            <el-input prefix-icon="el-icon-view" type="password" v-model="registerForm.password" clearable></el-input>
          </el-form-item>
          <el-form-item prop="check_password">
            <el-input prefix-icon="el-icon-view" type="password" v-model="registerForm.check_password" clearable>
            </el-input>
          </el-form-item>
          <el-form-item>
            <div class="flex justify-center">
              <el-button class="w-full" type="primary" @click="submitForm('registerForm')">注册</el-button>
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
      registerForm: {
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
        check_password: [
          { required: true, message: "请输入确认密码", trigger: "blur" },
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
    submitForm (formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          if (this.registerForm.password === this.registerForm.check_password)
            this.register();
          else
            this.$message({
              message: "密码不一致",
              center: true,
            });
        } else {
          this.$message({
            message: "error input",
            center: true,
          });
          return false;
        }
      });
    },
    resetForm (formName) {
      this.$refs[formName].resetFields();
    },
    register () {
      this.$request({
        url: "/api/auth/register",
        method: "post",
        data: {
          username: this.registerForm.username,
          password: this.registerForm.password,
        },
      })
        .then((response) => {
          this.$router.push("/");
        })
        .catch((error) => {
          this.$message({
            message: error,
            center: true,
          });
        });
    },
    home () {
      this.$router.push("/");
    },
    login () {
      this.$router.push("/login");
    }
  },
};
</script>

<style>
</style>