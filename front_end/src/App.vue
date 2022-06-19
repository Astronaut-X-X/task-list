<template>
  <div class="box m-32px bg-light-50 rounded-16px font-light">
    <router-view></router-view>
  </div>
</template>

<script>
export default {
  name: "app",
};
</script>

<style>
body {
  background-color: #e9ebef;
}

#app {
  padding: 0;
  margin: 0;
  min-height: 100vh;
}

.box {
  height: calc(100vh - 64px);
}
</style>
