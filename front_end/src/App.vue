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

*::-webkit-scrollbar {
  background-color: transparent;
  height: 12px;
  width: 12px;
}

*::-webkit-scrollbar-corner {
  background-color: transparent;
}

*::-webkit-scrollbar-thumb {
  transition: background 0.2s ease-in-out;
  border: 3px solid transparent;
  -webkit-background-clip: content-box;
  background-clip: content-box;
  --tw-bg-opacity: 1;
  background-color: rgba(229, 231, 235, var(--tw-bg-opacity));
  border-radius: 9999px;
}
</style>
