<template>
	<div class="header">
		<h3 style="padding-left: 20px" href="#">mindmap-go</h3>
	</div>
	<div id="app" class="page">

		<div class="sidebar adiv">
			<p>adasdasd</p>
			<button v-on:click="updateTree">Greet</button>
			<div v-if="store.selectedNode">
				<NodeForm v-on:updateSelection="updateSelected" />
			</div>
		</div>

		<div class="content adiv">
			<router-link to="/login" class="nav-link">Login</router-link>
			<router-view></router-view>
			<div>
				<Canvas ref="canvas" :data="tree" style="visibility: hidden;" />
			</div>
		</div>

	</div>
</template>

<script>
import axios from 'axios';
import Canvas from './components/Canvas.vue';
import { defineComponent } from 'vue'
import { store } from './store';
import NodeForm from './components/NodeForm.vue';
export default defineComponent({
	name: "App",
	components: {
		Canvas,
		NodeForm
	},
	data() {
		return {
			store,
			selection: store.selectedNode,
			users: null,
		}
	},
	methods: {
		getMaps() {
			axios.get("/api/users").then((response) => (this.users = response.data.data));
		},
	},
});
</script>