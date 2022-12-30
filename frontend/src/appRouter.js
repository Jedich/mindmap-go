import Login from "./components/Login.vue";
import Menu from "./components/Menu.vue";
import { createRouter, createWebHistory } from "vue-router";

const routes = [
	{ path: "/", component: Login },
	{ path: "/login", component: Login },
	{ path: "/app", component: Menu },
];

export const routeConfig = createRouter({
	history: createWebHistory(),
	routes: routes
});