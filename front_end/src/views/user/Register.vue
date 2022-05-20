<template>
  <div>
    <el-container>
      <el-main>
        <el-row>
          <el-col :span="6" :offset="12">
            <div>
              <el-form ref="form" :model="form" label-width="100px">
                <el-form-item label="用户名">
                  <el-input :model="form.username"></el-input>
                </el-form-item>
                <el-form-item label="密码">
                  <el-input :model="form.passward"></el-input>
                </el-form-item>
                <el-form-item label="确认密码">
                  <el-input :model="form.checkpswd"></el-input>
                </el-form-item>
                <el-form-item>
                  <el-button @click="register">注册</el-button>
                </el-form-item>
              </el-form>
            </div>
          </el-col>
        </el-row>
      </el-main>
    </el-container>
  </div>
</template>

<script>
export default {
  data() {
    return {
      form: {

      }
    }
  },
  methods:{
    register(){
      console.log("a")
    }
  },
  setup() {
    return {}
  }
}
</script>

<style lang="scss" scoped>
</style>