<template>
  <div class="flex">
    <div class="px-6 py-3 basis-4/12">
      <p class="mb-6 text-2xl">日计划</p>
      <div v-for="(item, index)  in todayDailyDetailData" v-bind:item="item" v-bind:index="index" v-bind:key="item.ID">
        <div class="my-2 flex flex-rol justify-between text-xl">
          <div class="w-34 flex flex-rol justify-start">
            <div>{{item.start_time}}</div>
            <div class="mx-2">-</div>
            <div>{{item.end_time}}</div>
          </div>
          <div class="flex-1 flex items-start">{{item.detail}}</div>
        </div>
      </div>
    </div>
    <div class="px-6 py-3 basis-4/12">
      <p class="mb-6 text-2xl">任务</p>
      <p class="text-right">
        <i class="el-icon-circle-plus-outline text-xl opacity-20 hover:opacity-100 transition-opacity"
          @click="insertTask"></i>
      </p>
      <div class="flex flex-col items-start">
        <div class="m-2 flex items-center" v-for="(item, index)  in todayTaskData" v-bind:item="item"
          v-bind:index="index" v-bind:key="item.ID">
          <el-input v-model="item.content" size="small" @keyup.enter.native="finishEditTask(item)"
            @blur="finishEditTask(item)" v-show="index === editingTask">
          </el-input>
          <el-checkbox v-model="item.done" v-show="index !== editingTask" @change="finishEditTask(item)" border>
            {{item.content}}
          </el-checkbox>
          <i class="el-icon-edit text-xl opacity-20 ml-6  hover:opacity-100 transition-opacity"
            @click="editTask(item,index)"></i>
          <i class="el-icon-delete text-xl opacity-20 ml-6  hover:opacity-100 transition-opacity"
            @click="deleteTask(item)"></i>
        </div>
      </div>
    </div>
    <div class="px-6 py-3 basis-4/12">
      <div class="flex flex-col">
        <div class="basis-4/12 min-h-40">
          <p class="mb-6 text-2xl">励志名言</p>
          <p class="text-right">
            <i class="el-icon-refresh text-xl opacity-20 ml-6  hover:opacity-100 transition-opacity"
              @click="refreshQuote(item)"></i>
          </p>
          <p class="indent-xl">希望是附丽于存在的，有存在，便有希望，有希望，便是光明。</p>
          <p class="text-right pre-quote-author">鲁迅</p>
        </div>
        <div class="basis-4/12 min-h-40 max-h-60 overflow-scroll overflow-x-hidden">
          <p class="mb-6 text-2xl">代办事项</p>

        </div>
        <div class="basis-4/12 min-h-40 max-h-60 overflow-scroll">
          <p class="mb-6 text-2xl">开心的事</p>

        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { getTodayDailyDetail } from '../../api/dailyDetail'
import { getTodayTask, addTask, updateTask, deleteTask } from '../../api/Task'
export default {
  name: 'Today',
  data () {
    return {
      todayDailyDetailData: [],
      todayTaskData: [],
      editingTask: -1,
      quote: {

      }
    };
  },
  mounted: function () {
    this.initTodaDailyDetailData();
    this.initTodaTaskData();
  },
  methods: {
    initTodaDailyDetailData () {
      getTodayDailyDetail({})
        .then((response) => {
          this.todayDailyDetailData = response.data;
        })
        .catch((error) => {
          this.$message({
            message: error,
            center: true,
          });
        });
    },

    initTodaTaskData () {
      getTodayTask({})
        .then((response) => {
          this.todayTaskData = response.data;
        })
        .catch((error) => {
          this.$message({
            message: error,
            center: true,
          });
        });
    },

    refreshQuote () {

    },

    insertTask () {
      addTask({
        done: false,
        content: '今日任务'
      })
        .then((response) => {
          this.initTodaTaskData();
        })
        .catch((error) => {
          this.$message({
            message: error,
            center: true,
          });
        });
    },

    editTask (item, index) {
      if (this.editingTask !== -1) {
        this.finishEditTask(item);
        this.editingTask = -1;
      } else {
        this.editingTask = index;
      }
    },

    finishEditTask (item) {
      this.editingTask = -1;
      updateTask({
        id: item.ID,
        content: item.content,
        done: item.done
      })
        .then((response) => {
          this.initTodaTaskData();
        })
        .catch((error) => {
          this.$message({
            message: error,
            center: true,
          });
        });
    },

    deleteTask (item) {
      deleteTask({
        id: item.ID
      })
        .then((response) => {
          this.initTodaTaskData();
        })
        .catch((error) => {
          this.$message({
            message: error,
            center: true,
          });
        });
    }
  }
}
</script>

<style>
.pre-quote-author::before {
  content: "—— ";
}
</style>