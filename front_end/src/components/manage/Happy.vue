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
      <el-table :data="happyData" style="width: 100%">
        <el-table-column prop="content" label="开心的事">
        </el-table-column>
        <el-table-column prop="time" label="时间" sortable :formatter="formatterTime" width="180">
        </el-table-column>
      </el-table>
    </div>
  </div>
</template>

<script>
import { getHappy } from '../../api/happy'
export default {
  name: 'Happy',

  data () {
    return {
      date: "",
      keyword: "",
      happyData: []
    }
  },

  mounted () {
    this.initHappyData();
  },

  methods: {
    initHappyData () {
      getHappy({})
        .then((response) => {
          this.happyData = response.data;
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

  }
}
</script>

<style>
</style>