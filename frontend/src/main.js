import { createApp } from 'vue'
import App from './App.vue'
import axios from 'axios'
import fabric from 'fabric';
import "bootstrap/dist/css/bootstrap.css"
import "bootstrap/dist/js/bootstrap.js"

const client = axios.create({
	baseURL: "/api",
});

import './assets/main.css'

const app = createApp(App)
app.mount('#app')
