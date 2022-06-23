<template>
  <div class="flex">
    <div class="px-6 py-3 basis-1/6 border-r">
      <div class="flex flex-col items-end">
        <i class="el-icon-circle-plus-outline text-xl hover:opacity-60 transition-opacity" @click="insert"></i>
      </div>
      <div class="flex flex-col items-end ">
        <div v-for="(item, index)  in dailyPlanData" v-bind:item="item" v-bind:index="index" v-bind:key="item.id"
          class="py-3 w-full flex flex-row items-center justify-center border-b">
          <div class="flex-1 flex justify-center">
            <a class="tracking-wide hover:text-blue-500 transition-all" @click="detail(item.id)">{{item.name}}</a>
          </div>
          <i class="el-icon-delete text-xl opacity-20 ml-6  hover:opacity-100 transition-opacity"
            @click="deleteItem(index)"></i>
        </div>
      </div>
    </div>
    <div class="px-6 py-3 basis-5/6">
    </div>
  </div>
</template>

<script>
import { getDailyPlan } from '../../api/dailyPlan'
export default {
  name: 'Daily',
  data () {
    return {
      dailyPlanData: [
        {
          id: 1,
          name: '模板'
        }
      ]
    };
  },
  mounted: function () {

  },
  methods: {
    init () {
      getDailyPlan.then((response) => {

      }).cathc((error) => {

      });
    },
    insert () {
      this.dailyPlanData.push({
        id: 0,
        name: '模板'
      });
    },
    deleteItem (index) {
      this.dailyPlanData = this.dailyPlanData.filter((elem, i, arr) => { return i !== index });
    },
    detail (id) {
      console.log(id);
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
</style>