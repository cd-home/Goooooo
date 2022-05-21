import {createApp} from 'vue'
import App from '@/App.vue'
import router from "@/router"
import elementPlus from 'element-plus'
import api from '@/api/module'


const app = createApp(App)
app.use(router)

// Global APIS
app.config.globalProperties.$api = api

app.use(elementPlus)
app.mount('#app')
