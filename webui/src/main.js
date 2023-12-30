import {createApp, reactive} from 'vue'
import App from './App.vue'
import router from './router'
import { axios, pathToProfile, setAuth } from './services/axios.js';
import { getCurrentId, getCurrentUsername } from './services/authentication.js';
import ErrorMsg from './components/ErrorMsg.vue'
import LoadingSpinner from './components/LoadingSpinner.vue'

import './assets/dashboard.css'
import './assets/main.css'
import './assets/styles.css'


const timestampToDate = (timestamp) => {
    const difference = new Date() - new Date(timestamp);	

    var seconds = Math.floor(difference / 1000);
    var minutes = Math.floor(seconds / 60);
    var hours = Math.floor(minutes / 60);
    var days = Math.floor(hours / 24);
    
    if (days > 0) {
        return `${days} day${days > 1 ? 's' : ''} ago`;
    } else if (hours > 0) {
        return `${hours} hour${hours > 1 ? 's' : ''} ago`;
    } else if (minutes > 0) {
        return `${minutes} minute${minutes > 1 ? 's' : ''} ago`;
    } else {
        return `${seconds} seconds ago`;
    }
}

const app = createApp(App)
app.config.globalProperties.$axios = axios;
app.config.globalProperties.$setAuth = setAuth;
app.config.globalProperties.$getCurrentId = getCurrentId;
app.config.globalProperties.$getCurrentUsername = getCurrentUsername;
app.config.globalProperties.$timestampToDate = timestampToDate;
app.config.globalProperties.$pathToProfile = pathToProfile;

app.component("ErrorMsg", ErrorMsg);
app.component("LoadingSpinner", LoadingSpinner);
app.use(router)
app.mount('#app')
