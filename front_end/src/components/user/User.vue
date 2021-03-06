<template>
  <div class="container flex justify-center content-center">
    <div class="h-full w-2/5 mt-6">
      <el-row>
        <el-col :span="8">
          <p class="text-3xl">用户信息</p>
        </el-col>
        <el-col :span="16">
          <div class="flex justify-end content-center">
            <el-button type="primary" @click="edit">编辑</el-button>
          </div>
        </el-col>
      </el-row>
      <el-row class="flex items-end">
        <el-col :span="8">
          <span>用户头像</span>
        </el-col>
        <el-col :span="8">
          <el-avatar :size="64" :src="user.image || '/header_image.jpg'"></el-avatar>
        </el-col>
        <el-col :span="8">
          <div class="flex justify-end content-center">
            <el-upload :action="imageUploadUrl" :headers="headers" :show-file-list="false"
              :on-success="handleAvatarSuccess" :before-upload="beforeAvatarUpload">
              <el-button type="info">上传</el-button>
            </el-upload>
          </div>
        </el-col>
      </el-row>
      <el-form :model="user" :rules="rules" ref="userForm" class="user-form">
        <el-row class="mt-6">
          <el-col :span="8">
            <span>用户名</span>
          </el-col>
          <el-col :span="16">
            <span v-if="!isEditing">{{user.username}}</span>
            <el-form-item v-else prop="username">
              <el-input type="text" v-model="user.username" placeholder="用户名" clearable></el-input>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row class="mt-6">
          <el-col :span="8">
            <span>用户邮箱</span>
          </el-col>
          <el-col :span="16">
            <span v-if="!isEditing">{{user.email}}</span>
            <el-form-item v-else prop="email">
              <el-input v-model="user.email" placeholder="邮箱" clearable></el-input>
            </el-form-item>
          </el-col>
        </el-row>
      </el-form>
      <el-row class="mt-6">
        <el-col :span="8">
          <span>用户创建时间</span>
        </el-col>
        <el-col :span="16">
          <span>{{user.CreatedAt}}</span>
        </el-col>
      </el-row>
      <el-row class="mt-6">
        <el-col :span="8">
          <span>用户修改时间</span>
        </el-col>
        <el-col :span="16">
          <span>{{user.UpdatedAt}}</span>
        </el-col>
      </el-row>
      <transition name="el-fade-in-linear">
        <el-row type="flex" v-if="isEditing" justify="center" class="mt-6">
          <el-col :span="8">
            <div class="flex justify-center content-center">
              <el-button class="w-1/2" type="primary" @click="submit('userForm')">保存</el-button>
            </div>
          </el-col>
          <el-col :span="8">
            <div class="flex justify-center content-center">
              <el-button class="w-1/2" type="info" @click="cancel">取消</el-button>
            </div>
          </el-col>
        </el-row>
      </transition>
    </div>
  </div>
</template>

<script>
import { updateUser } from '../../api/user';
export default {
  name: "Login",
  data () {
    let validateUsername = (rule, value, callback) => {
      if (value.length < 6 || value.length > 20) {
        callback(new Error('长度在 6 到 20 个字符'));
      } else {
        callback();
      }
    };
    let validateEmail = (rule, value, callback) => {
      let regex = /^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$/;
      if (!regex.test(value)) {
        callback(new Error('邮箱不符合规范'));
      } else {
        callback();
      }
    };
    return {
      isEditing: false,
      dialogVisible: false,
      imageUploadUrl: "/api/v1/user/image",
      fileList: [],
      headers: {
        'Authorization': this.$storage.getItem('token')
      },
      user: {
        username: '',
        email: ''
      },
      rules: {
        username: [{ validator: validateUsername, trigger: 'blur' }],
        email: [{ validator: validateEmail, trigger: 'blur' }]
      }
    };
  },
  mounted: function () {
    this.initUserInfo();
  },
  methods: {
    initUserInfo () {
      this.user = this.$storage.getItem('user');
    },
    edit () {
      this.initUserInfo();
      this.isEditing = !this.isEditing;
    },
    submit (formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          updateUser({
            username: this.user.username,
            email: this.user.email
          }).then((response) => {
            if (response.user) {
              this.$storage.setItem('user', response.user);
              this.user = response.user;
            }
            this.isEditing = false;
            this.$message.success({ message: "修改成功" })
          }).catch((error) => {
            this.$message.error({ message: error });
          });
        } else {
          this.$message.error({ message: "错误输入" });
          return false;
        }
      });
    },
    cancel () {
      this.initUserInfo();
      this.isEditing = false;
    },
    handleAvatarSuccess (res, file) {
      let user = this.$storage.getItem('user');
      user.image = res.data.filename;
      this.$storage.setItem('user', user);
      // this.imageUrl = URL.createObjectURL(file.raw);
    },
    beforeAvatarUpload (file) {
      const isJPG = file.type === 'image/jpeg';
      const isLt2M = file.size / 1024 / 1024 < 2;

      if (!isJPG) {
        this.$message.error('上传头像图片只能是 JPG 格式!');
      }
      if (!isLt2M) {
        this.$message.error('上传头像图片大小不能超过 2MB!');
      }
      return isJPG && isLt2M;
    }
  }
};
</script>

<style scoped>
.user-form .el-form-item {
  margin-bottom: 0px;
}
</style>