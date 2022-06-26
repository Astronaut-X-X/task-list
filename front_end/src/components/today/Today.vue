<template>
  <div class="flex">
    <div class="px-6 py-3 basis-4/12">
      <p class="mb-6 text-2xl">日计划</p>
      <div v-for="(item, index)  in todaDailyDetailData" v-bind:item="item" v-bind:index="index" v-bind:key="item.ID">
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
    </div>
    <div class="px-6 py-3  basis-4/12">
      <div class="flex flex-col">
        <div class="basis-4/12 min-h-40">
          <p class="mb-6 text-2xl">励志名言</p>
          <p class="text-right">
            <i class="el-icon-refresh text-xl opacity-20 ml-6  hover:opacity-100 transition-opacity"
              @click="refreshQuote(item)"></i>
          </p>
          <p class="indent-xl">希望是附丽于存在的，有存在，便有希望，有希望，便是光明。</p>
          <p class="text-right pre-content">鲁迅</p>
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
export default {
  name: 'Today',
  data () {
    return {
      todaDailyDetailData: [],
      quote: {
        
      }
    };
  },
  mounted: function () {
    this.initTodaDailyDetailData();
  },
  methods: {
    initTodaDailyDetailData () {
      getTodayDailyDetail({})
        .then((response) => {
          this.todaDailyDetailData = response.data;
        })
        .catch((error) => {
          this.$message({
            message: error,
            center: true,
          });
        });
    },
    refreshQuote () {

    }
  }
}
</script>

<style>
.pre-content::before {
  content: "——";
}
</style>