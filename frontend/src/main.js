import { createApp } from 'vue'
import App from './App.vue'
import axios from 'axios'
import * as appRouter from './appRouter'
import store from './store/index'
import VueCookies from 'vue-cookies'
import "bootstrap/dist/css/bootstrap.css"
import "bootstrap/dist/js/bootstrap.js"

const client = axios.create({
	baseURL: "/api",
});

import './assets/main.css'

const app = createApp(App);

app.use(appRouter.routeConfig);
app.use(store);
app.use(VueCookies);
app.mount('#app');
