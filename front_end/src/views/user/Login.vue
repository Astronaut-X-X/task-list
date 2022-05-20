<template>
  <el-container>
    <el-main>
      <el-row>
        <el-col :span="6" :offset="12">
          <div>
            <el-form ref="form" :model="form" label-width="180px">
              <el-form-item label="用户名">
                <el-input v-model="form.username"></el-input>
              </el-form-item>
              <el-form-item label="密码">
                <el-input v-model="form.password"></el-input>
                <a href="#">忘记密码</a>
              </el-form-item>
              <el-form-item>
                <el-button>登录</el-button>
                <el-button>注册</el-button>
              </el-form-item>
            </el-form>
          </div>
        </el-col>
      </el-row>
    </el-main>
  </el-container>
</template>

<script>
export default {
  data() {
    return {
      form: {

      }
    }
  }
}
</script>

<style lang="scss" scoped>
</style>