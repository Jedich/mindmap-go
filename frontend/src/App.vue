<template>
	<div class="header">
		<h3 style="padding-left: 20px" href="#">mindmap-go</h3>
	</div>
	<div class="page">
		<router-view></router-view>
	</div>
</template>

<script>
import axios from 'axios';
import Canvas from './components/Canvas.vue';
import { defineComponent } from 'vue'
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
	},
	mounted() {
		if (!this.$cookies.isKey("token")) {
			this.$router.push("/login");
		} else {
			this.$router.push("/app");
		}
	}
});
</script>