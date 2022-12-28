<template>
	<div class="header">
		<h3 style="padding-left: 20px" href="#">mindmap-go</h3>
	</div>
	<div id="app" class="page">

		<div class="sidebar adiv">
			<p>adasdasd</p>
			<button v-on:click="updateTree">Greet</button>
			<div v-if="store.selectedNode">
				<NodeForm v-on:updateSelection="updateSelected"/>
			</div>
		</div>

		<div class="content adiv">
			<ul class="nav nav-tabs">
				<li class="nav-item">
					<p class="nav-link active" aria-current="page">Active <input class="btn btn-outline-danger"
							style="--bs-btn-padding-y: .05rem; --bs-btn-padding-x: .4rem; --bs-btn-font-size: .5rem; --bs-btn-border-radius: 1rem;"
							type="button" value=""></p>
				</li>
				<li class="nav-item">
					<a class="nav-link" href="#">Link</a>
				</li>
				<li class="nav-item">
					<a class="nav-link" href="#">Link</a>
				</li>
				<li class="nav-item">
					<a class="nav-link disabled">Disabled</a>
				</li>
			</ul>
			<div>
				<Canvas ref="canvas" :data="tree" />
			</div>
		</div>

	</div>
</template>

<script>
import Canvas from './components/Canvas.vue';
import axios from 'axios';
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
			tree: {
				name: "text text text text text text text text text",
				color: "#ff0000",
				children: [
					{
						name: "123490\n1123456711234 567890112345678901",
						children: [{
							name: "b",
							color: "#32a852",
							children: []
						},
						{
							name: "c",
							children: []
						},
						{
							name: "c",
							children: []
						}]
					},
					{
						name: "c",
						children: [{
							name: "b",
							children: []
						},
						{
							name: "c",
							children: []
						}]
					}
				]
			}
		};
	},
	methods: {
		updateTree() {
			//this.tree.children.push()
			//this.$refs.canvas.a();
			this.$refs.canvas.updateFromSky({
				name: "toddler",
				color: "violet",
				children: []
			});
		},
		updateSelected() {
			this.$refs.canvas.updateSelected()
		},
		getMaps() {
			axios.get("/api/users").then((response) => (this.users = response.data.data));
		},
	},
});
</script>