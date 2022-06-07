<template>
    <!-- <div class="bg-blue-300 h-screen"> -->
    <div class="h-screen">
        <header class="pt-8 pr-10">
            <nav class="flex justify-end">
                <div class="" v-if="true">
                    <a class="mx-4 p-2 w-28 font-light tracking-wide hover:bg-blue-200 transition-all"
                        @click="login">Login</a>
                    <a class="mx-4 p-2 w-28 font-light tracking-wide hover:bg-blue-200 transition-all"
                        @click="register">Register</a>
                </div>
                <div v-else>
                    <a>Task</a>
                    <img src="" alt="">
                </div>
            </nav>
        </header>
        <main>
            <transition mode="out-in">
                <router-view></router-view>
            </transition>
        </main>
    </div>
</template>

<script>
export default {
    name: 'Home',
    methods: {
        login() {
            this.$router.push('/login');
        },
        register() {
            this.$router.push('/register');
        }
    }
}
</script>

<style>
</style>