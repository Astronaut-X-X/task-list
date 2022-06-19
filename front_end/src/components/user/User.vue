<template>
  <div
    class="container flex flex-col justify-center content-center px-16 pt-16 pb-8 mt-16 w-1/2 bg-light-200 rounded-2xl border-2 border-blue-50 shadow-sm shadow-blue-200 hover:shadow-md duration-800 ">
    <el-row>
      <el-col :span="8">
        <p>用户信息</p>
      </el-col>
      <el-col type="flex" :span="16" justify="end">
        <div class="flex justify-end content-center">
          <el-button @click="edit">编辑</el-button>
        </div>
      </el-col>
    </el-row>
    <el-row class="mt-16">
      <el-col :span="8">
        <span>用户头像</span>
      </el-col>
      <el-col :span="8">
        <el-avatar :size="size" :src="'/header_image.jpg'"></el-avatar>
      </el-col>
      <el-col :span="8">
        <el-button @click="upload">上传头像</el-button>
      </el-col>
    </el-row>
    <el-row class="mt-10">
      <el-col :span="8">
        <span>用户名</span>
      </el-col>
      <el-col :span="16">
        <span v-if="!isEditing">{{user.username}}</span>
        <el-input v-else v-model="user.username" placeholder="用户名"></el-input>
      </el-col>
    </el-row>
    <el-row class="mt-10">
      <el-col :span="8">
        <span>用户邮箱</span>
      </el-col>
      <el-col :span="16">
        <span v-if="!isEditing">{{user.email}}</span>
        <el-input v-else v-model="user.email" placeholder="邮箱"></el-input>
      </el-col>
    </el-row>
    <el-row class="mt-10">
      <el-col :span="8">
        <span>用户创建时间</span>
      </el-col>
      <el-col :span="16">
        <span>{{user.CreatedAt}}</span>
      </el-col>
    </el-row>
    <el-row class="mt-10">
      <el-col :span="8">
        <span>用户修改时间</span>
      </el-col>
      <el-col :span="16">
        <span>{{user.UpdatedAt}}</span>
      </el-col>
    </el-row>
    <transition name="el-fade-in-linear">
      <el-row type="flex" v-if="isEditing" justify="center" class="mt-8">
        <el-col :span="8">
          <div class="flex justify-center content-center">
            <el-button @click="submit">保存</el-button>
          </div>
        </el-col>
        <el-col :span="8">
          <div class="flex justify-center content-center">
            <el-button @click="cancel">取消</el-button>
          </div>
        </el-col>
      </el-row>
    </transition>
  </div>
</template>

<script>
export default {
  name: "Login",
  data () {
    return {
      isEditing: false
    };
  },
  computed: {
    user: function () {
      return this.$storage.getItem('user');
    }
  },
  methods: {
    edit () {
      this.isEditing = !this.isEditing;
    },
    submit () {

    },
    cancel () {
      this.isEditing = false;
    }
  }
};
</script>

<style>
</style>