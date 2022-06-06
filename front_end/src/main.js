import Vue from 'vue'
import App from './App.vue'
import store from 'vuex'
import request from './utils/request'
import storage from './utils/storage'
import router from './router'

Vue.config.productionTip = false
Vue.prototype.$request = request;
Vue.prototype.$storage = storage;

new Vue({
    router,
    store,
    render: h => h(App)
}).$mount("#app")