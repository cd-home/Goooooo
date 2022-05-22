import {createApp} from 'vue'
import App from '@/App.vue'
import router from "@/router"
import store from '@/store'
import elementPlus from 'element-plus'
import apis from '@/api/module'
import notify from '@/utils/notify'

// need to import
import 'element-plus/dist/index.css'


const app = createApp(App)
app.use(router)
app.use(store)

app.use(elementPlus)

// Global APIS
app.config.globalProperties.$api = apis
app.config.globalProperties.$notify = notify

app.mount('#app')
