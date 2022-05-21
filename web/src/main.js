import {createApp} from 'vue'
import App from './App.vue'
import router from "./router";
import elementPlus from 'element-plus'

console.log(process)
console.log(1111)

const app = createApp(App)
app.use(router)
app.use(elementPlus)
app.mount('#app')
