import Vue from 'vue'
import App from './App.vue'
import store from 'vuex'
import request from './utils/request'
import storage from './utils/storage'
import router from './routers'
import 'virtual:windi.css'
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';

Vue.config.productionTip = false
Vue.prototype.$request = request;
Vue.prototype.$storage = storage;
Vue.use(ElementUI);

new Vue({
    router,
    store,
    render: h => h(App)
}).$mount("#app")
