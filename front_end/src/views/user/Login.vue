<template>
  <el-container>
    <el-main>
      <el-row type="flex" justify="center">
        <el-col :span="7">
          <div class="login-box">
            <el-form ref="form" :model="form" class="form">
              <el-form-item class="">
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
        </el-col>
      </el-row>
    </el-main>
  </el-container>
</template>

<script>
import router from '../../router'

export default {
  data() {
    return {
      form: {

      }
    }
  },
  methods:{
    register(){
      router.push('/register')
    } 
  }
}
</script>

<style lang="scss" scoped>

</style>