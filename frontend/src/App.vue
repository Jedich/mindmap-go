<template>
	<div class="header" style="display:flex; justify-content:space-between; align-items:center; padding:10px;">
		<h3 style="padding-left: 20px" href="#">mindmap-go</h3>
		<div style="margin-left:auto;" v-if="$route.path === '/app'">
        	<button class="btn btn-outline-danger" v-on:click="logout">Logout</button>
      	</div>
	</div>
	<div class="page">
		<router-view></router-view>
	</div>
</template>

<script>
import axios from 'axios';
import Canvas from './components/Canvas.vue';
import { defineComponent } from 'vue';
import { mapActions, mapGetters } from "vuex";

export default defineComponent({
	inject: ["$cookies"],
	name: "App",
	components: {
		Canvas
	},
	data() {
		return {
			users: null,
		}
	},
	methods: {
		getMaps() {
			axios.get("/api/users").then((response) => (this.users = response.data.data));
		},
		...mapActions("maps", {
			initState: "initState",
			clearAllTabs: "clearAllTabs"
		}),
		logout() {
			this.$cookies.remove('token');
			this.clearAllTabs();
			localStorage.clear();
			window.location.reload();
			this.$router.push("/login");
		}
	},
	created() {
		this.initState();
	},
	mounted() {
		if (!this.$cookies.isKey("token")) {
			localStorage.clear();
			this.$router.push("/login");
		} else {
			this.$router.push("/app");
		}
	}
});
</script>