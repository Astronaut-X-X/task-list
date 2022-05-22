<template>
  <transition>
    <router-view />
  </transition>
</template>

<script>
export default {
  name: "app",
};
</script>

<style>
#app {
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}

html,body{
  height: 100%;
  background-image: linear-gradient(120deg, #a1c4fd 0%, #c2e9fb 100%);
  overflow:hidden;
}
</style>
