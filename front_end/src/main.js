import { createApp } from 'vue'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import App from './App.vue'
import router from './router'
import request from './utils/request'
import storage from './utils/storage'
import './index.css'

const app = createApp(App)
app.use(ElementPlus)
app.use(router)
app.mount('#app')
app.config.globalProperties.$request = request
app.config.globalProperties.$storage = storage
