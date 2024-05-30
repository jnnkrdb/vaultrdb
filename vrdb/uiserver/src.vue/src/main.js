import { createApp } from 'vue'
import PrimeVue from 'primevue/config'
import App from './App.vue'
import axios from 'axios'
import Tooltip from 'primevue/tooltip';


// config axios
axios.defaults.baseURL = 'http://localhost:30080/'

const app = createApp(App);
app.config.globalProperties.$axios = axios

// create the user auth for the app
app.config.globalProperties.$username = ""
app.config.globalProperties.$password = ""


app.directive('tooltip', Tooltip);
app.use(PrimeVue, {})

app.mount('#app')