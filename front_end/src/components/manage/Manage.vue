<template>
  <div class="flex">
    <div class="px-6 py-3 basis-1/6 border-r">
      <div class="flex flex-col items-end">
        <a class="mx-4 p-2 tracking-wide hover:text-blue-500 transition-all" @click="daily">日计划</a>
        <a class="mx-4 p-2 tracking-wide hover:text-blue-500 transition-all" @click="task">任务</a>
        <a class="mx-4 p-2 tracking-wide hover:text-blue-500 transition-all" @click="todo">代办事项</a>
        <a class="mx-4 p-2 tracking-wide hover:text-blue-500 transition-all" @click="quote">励志名言</a>
        <a class="mx-4 p-2 tracking-wide hover:text-blue-500 transition-all" @click="happy">开心的事</a>
      </div>
    </div>
    <div class="px-6 basis-5/6">
      <router-view></router-view>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Manage',
  data () {
    return {
    };
  },
  mounted: function () {

  },
  methods: {
    daily () {
      this.$router.push({ path: '/manage/daily' });
    },
    task () {
      this.$router.push({ path: '/manage/task' });
    },
    todo () {
      this.$router.push({ path: '/manage/todo' });
    },
    quote () {
      this.$router.push({ path: '/manage/quote' });
    },
    happy () {
      this.$router.push({ path: '/manage/happy' });
    }
  }
}
</script>

<style scoped>
.border-r {
  border-right: solid 2px #dedede;
}
</style>