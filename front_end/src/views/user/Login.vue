<template>
  <el-container>
    <el-main>
      <el-row>
        <el-col :span="8" :offset="12">
          <el-form ref="form" :model="form">
            <el-form-item label="用户名">
              <el-input v-model="form.usernamw"></el-input>
            </el-form-item>
            <el-form-item label="密码">
              <el-input v-model="form.password"></el-input>
            </el-form-item>
          </el-form>
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