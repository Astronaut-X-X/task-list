<template>
  <div>
    <div class="flex flex-rol justify-between">
      <div class="w-3/12">
        <el-input placeholder="请输入内容" suffix-icon="el-icon-search" v-model="keyword">
        </el-input>
      </div>
      <div>
        <el-date-picker v-model="date" type="daterange" start-placeholder="开始日期" end-placeholder="结束日期"
          :default-time="['00:00:00', '23:59:59']">
        </el-date-picker>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Task',
  data () {
    return {
      date: "",
      keyword: ""
    }
  },
  methods: {
    
  }
}
</script>

<style>
</style>