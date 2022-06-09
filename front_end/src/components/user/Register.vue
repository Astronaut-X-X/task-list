<template>
  <div class="container flex justify-center content-center">
    <el-form :model="registerForm" :rules="rules" ref="registerForm" class="
        bg-blue-100
        pt-16
        pb-10
        px-16
        mt-24
        rounded-2xl
        border-2 border-blue-50
        shadow-sm shadow-blue-200
        hover:shadow-md
        duration-800
      ">
      <el-form-item prop="username">
        <el-input prefix-icon="el-icon-user" type="text" v-model="registerForm.username" clearable></el-input>
      </el-form-item>
      <el-form-item prop="password">
        <el-input prefix-icon="el-icon-view" type="password" v-model="registerForm.password" clearable></el-input>
      </el-form-item>
      <el-form-item prop="check_password">
        <el-input prefix-icon="el-icon-view" type="password" v-model="registerForm.check_password" clearable></el-input>
      </el-form-item>
      <el-form-item>
        <div class="flex justify-center">
          <el-button type="primary" @click="submitForm('registerForm')">注册</el-button>
          <el-button @click="resetForm('registerForm')">重置</el-button>
        </div>
      </el-form-item>
    </el-form>
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
  },
};
</script>

<style>
</style>