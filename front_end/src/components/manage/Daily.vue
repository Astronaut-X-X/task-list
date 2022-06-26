<template>
  <div class="flex">
    <div class="px-6 py-3 basis-1/6 border-r">
      <div class="flex flex-col items-end">
        <i class="el-icon-circle-plus-outline text-xl hover:opacity-60 transition-opacity" @click="insert"></i>
      </div>
      <div class="flex flex-col items-end ">
        <div v-for="(item, index)  in dailyPlanData" v-bind:item="item" v-bind:index="index" v-bind:key="item.ID"
          class="py-2 w-full flex flex-row items-center justify-center border-b">
          <div class="flex-1 flex justify-center" @dblclick="edit(index)">
            <a v-if="index !== editing" class="tracking-wide hover:text-blue-500 transition-all"
              @click="clickPlan(item.ID)">{{item.name}}</a>
            <el-input v-else v-model="item.name" @keyup.enter.native="finishEdit(item)" @blur="finishEdit(item)"
              size="small" />
          </div>
          <i class="el-icon-delete text-xl opacity-20 ml-6  hover:opacity-100 transition-opacity"
            @click="deletePlan(item)"></i>
        </div>
      </div>
    </div>
    <div class="px-6 py-3 basis-3/6">
      <div class="flex flex-col items-end">
        <i class="el-icon-circle-plus-outline text-xl hover:opacity-60 transition-opacity"
          v-show="currentDailyPlanID !== -1" @click="insertDetail"></i>
      </div>
      <div class="flex flex-col items-end ">
        <div v-for="(item, index)  in dailyDetailData" v-bind:item="item" v-bind:index="index" v-bind:key="item.ID"
          class="py-1 w-full flex flex-row items-center justify-center border-b" v-show="currentDailyPlanID !== -1">
          <div class="flex-1 flex justify-start items-center" @dblclick="editDetail(index)">
            <div class="flex justify-evenly items-center mr-4">
              <el-time-select class="time-select" placeholder="起始时间" v-model="item.start_time"
                @change="finishEditDetail(item)" :picker-options="pickerOptions" size="small">
              </el-time-select>
              <div class="mx-2">-</div>
              <el-time-select class="time-select" placeholder="结束时间" v-model="item.end_time"
                @change="finishEditDetail(item)" :picker-options="pickerOptions" size="small">
              </el-time-select>
            </div>
            <a v-if="index !== editingDetail"
              class="tracking-wide hover:text-blue-500 transition-all">{{item.detail}}</a>
            <el-input v-else v-model="item.detail" @keyup.enter.native="finishEditDetail(item)"
              @blur="finishEditDetail(item)" size="small" />
          </div>
          <i class="el-icon-delete text-xl opacity-20 ml-6  hover:opacity-100 transition-opacity"
            @click="deleteDetail(item)"></i>
        </div>
      </div>
    </div>
    <div class="px-6 py-3 basis-2/6">
      <div class="flex flex-col items-start" v-show="currentDailyPlanID !== -1">
        <div class="m-2" v-for="(item, index)  in weekPlanData" v-bind:item="item" v-bind:index="index"
          v-bind:key="item.ID">
          <el-checkbox v-model="item.check" size="medium" @change="check(item)"
            :disabled="item.daily_plan_id !== -1 && item.daily_plan_id !== currentDailyPlanID" border>{{item.name}}
          </el-checkbox>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { getDailyPlan, addDailyPlan, updateDailyPlan, deleteDailyPlan } from '../../api/dailyPlan'
import { getDailyDetail, addDailyDetail, updateDailyDetail, deleteDailyDetail } from '../../api/dailyDetail'
import { getWeekPlan, addWeekPlan, updateWeekPlan, deleteWeekPlan } from '../../api/weekPlan'
export default {
  name: 'Daily',
  data () {
    return {
      editing: -1,
      currentDailyPlanID: -1,
      editingDetail: -1,
      dailyPlanData: [],
      dailyDetailData: [
        {
          ID: 0,
          name: "XXX",
        }
      ],
      pickerOptions: {
        start: '00:00',
        step: '00:10',
        end: '24:00'
      },
      weekPlanData: []
    };
  },

  mounted: function () {
    this.init();
  },

  methods: {
    initWeek () {
      this.weekPlanData.length = 0;
      for (let i = 0; i <= 6; i++) {
        this.weekPlanData.push(
          {
            daily_plan_id: -1,
            week: i,
            check: false,
            name: this.formatWeek(i),
          }
        )
      }
      let tmpData = [];
      getWeekPlan({})
        .then((response) => {
          tmpData = response.data;
          tmpData.forEach((value, index, array) => {
            this.weekPlanData[value.week].ID = value.ID;
            this.weekPlanData[value.week].daily_plan_id = value.daily_plan_id;
            this.weekPlanData[value.week].check = true;
          })
        })
        .catch((error) => {
          this.$message({
            message: error,
            center: true,
          });
        });
    },

    formatWeek (week) {
      switch (week) {
        case 0:
          return '星期天';
        case 1:
          return '星期一';
        case 2:
          return '星期二';
        case 3:
          return '星期三';
        case 4:
          return '星期四';
        case 5:
          return '星期五';
        case 6:
          return '星期六';
      }
    },

    init () {
      getDailyPlan({})
        .then((response) => {
          this.dailyPlanData = response.data;
        })
        .catch((error) => {
          this.$message({
            message: error,
            center: true,
          });
        });
    },

    insert () {
      addDailyPlan({
        name: '模板'
      })
        .then((response) => {
          this.init();
        })
        .catch((error) => {
          this.$message({
            message: error,
            center: true,
          });
        });
    },

    deletePlan (item) {
      deleteDailyPlan({
        id: item.ID
      })
        .then((response) => {
          this.init();
        })
        .catch((error) => {
          this.$message({
            message: error,
            center: true,
          });
        });
      if (this.currentDailyPlanID === item.ID) {
        this.currentDailyPlanID = -1;
      }
    },

    edit (index) {
      this.editing = index;
    },

    finishEdit (item) {
      updateDailyPlan({
        id: item.ID,
        name: item.name
      })
        .then((response) => {
          this.init();
        })
        .catch((error) => {
          this.$message({
            message: error,
            center: true,
          });
        });
      this.editing = -1
    },

    editDetail (index) {
      this.editingDetail = index;
    },

    insertDetail () {
      addDailyDetail({
        daily_plan_id: this.currentDailyPlanID,
        detail: '活动',
        start_time: '00:00',
        end_time: '00:00'
      })
        .then((response) => {
          this.getDetailData(this.currentDailyPlanID);
        })
        .catch((error) => {
          this.$message({
            message: error,
            center: true,
          });
        });
    },

    deleteDetail (item) {
      deleteDailyDetail({
        id: item.ID
      })
        .then((response) => {
          this.getDetailData(this.currentDailyPlanID);
        })
        .catch((error) => {
          this.$message({
            message: error,
            center: true,
          });
        });
      if (this.currentDailyPlanID === item.ID) {
        this.currentDailyPlanID = -1;
      }
    },

    finishEditDetail (item) {
      updateDailyDetail({
        id: item.ID,
        daily_plan_id: item.daily_plan_id,
        detail: item.detail,
        start_time: item.start_time,
        end_time: item.end_time
      })
        .then((response) => {
          this.getDetailData(item.daily_plan_id);
        })
        .catch((error) => {
          this.$message({
            message: error,
            center: true,
          });
        });
      this.editingDetail = -1;
    },

    getDetailData (id) {
      this.currentDailyPlanID = id;
      getDailyDetail({
        id: id
      })
        .then((response) => {
          this.dailyDetailData = response.data;
        })
        .catch((error) => {
          this.$message({
            message: error,
            center: true,
          });
        });
    },

    clickPlan (id) {
      this.getDetailData(id);
      this.initWeek();
    },

    check (item) {
      if (item.check) {
        addWeekPlan({
          daily_plan_id: this.currentDailyPlanID,
          week: item.week
        })
          .then((response) => {
            this.initWeek();
          })
          .catch((error) => {
            this.$message({
              message: error,
              center: true,
            });
          });
      } else {
        deleteWeekPlan({
          id: item.ID
        })
          .then((response) => {
            this.initWeek();
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
}
</script>

<style scoped>
.border-r {
  border-right: solid 2px #dedede;
}
.border-b {
  border-bottom: solid 2px #dedede;
}

.time-select {
  width: 114px;
}
</style>