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
    <div class="flex flex-col justify-start">
      <el-table :data="taskData" style="width: 100%">
        <el-table-column prop="done" label="情况" :formatter="formatterDone" width="100">
        </el-table-column>
        <el-table-column prop="content" label="任务">
        </el-table-column>
        <el-table-column prop="time" label="时间" sortable :formatter="formatterTime" width="180">
        </el-table-column>
      </el-table>
    </div>
  </div>
</template>

<script>
import { getTask } from '../../api/task'
export default {
  name: 'Task',

  data () {
    return {
      date: "",
      keyword: "",
      taskData: []
    }
  },

  mounted () {
    this.initTaskData();
  },

  methods: {
    initTaskData () {
      getTask({})
        .then((response) => {
          this.taskData = response.data;
        })
        .catch((error) => {
          this.$message({
            message: error,
            center: true,
          });
        });
    },

    formatterTime (row, column) {
      return row.time.slice(0, 10);
    },

    formatterDone (row, column) {
      if (row.done) {
        return '完成';
      }
      return '未完成';
    }
  }
}
</script>

<style>
</style>