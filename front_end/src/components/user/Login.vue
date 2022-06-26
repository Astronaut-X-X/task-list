<template>
  <div>
    <header class="py-8 px-10">
      <nav class="flex justify-between">
        <div class="flex content-end">
          <span class="text-3xl mr-8 divide font-extrabold" @click="home">TaskList</span>
        </div>
        <div class="flex content-end">
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
          <el-form-item prop="password">
            <el-input prefix-icon="el-icon-view" type="password" v-model="loginForm.password" clearable></el-input>
          </el-form-item>
          <el-form-item>
            <div>
              <el-button class="w-full" type="primary" @click="submitForm('loginForm')">登录</el-button>
            </div>
          </el-form-item>
          <el-form-item class="flex justify-end">
            <a class="forget" @click="forget">忘记密码</a>
          </el-form-item>
        </el-form>
      </div>
    </main>
  </div>
</template>

<script>
import { loginUser } from '../../api/user'
export default {
  name: "Login",
  data () {
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
    resetForm (formName) {
      this.$refs[formName].resetFields();
    },
    login () {
      loginUser({
        username: this.loginForm.username,
        password: this.loginForm.password
      })
        .then((response) => {
          if (response.token) {
            this.$storage.setItem("token", response.token);
          }
          this.$router.push('/today');
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
.forget {
  color: #409eff;
}
.forget:hover {
  color: #66b1ff;
}
</style>