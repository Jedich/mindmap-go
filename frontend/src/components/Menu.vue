<template>
	<div class="sidebar adiv">
		<p>adasdasd</p>
		<button v-on:click="updateTree">Greet</button>
		<div v-if="getCurrentNode">
			<NodeForm v-on:updateSelection="updateSelected" />
		</div>
	</div>

	<div class="content adiv">
		<ul class="nav nav-tabs">
			<li class="nav-item" v-for="map in maps">
				<p class="nav-link active" aria-current="page">{{ map.name }} <input class="btn btn-outline-danger x" type="button" value=""></p>
			</li>
			<li class="nav-item disabled">
				<a class="nav-link" href="#">+</a>
			</li>
		</ul>
		<div v-if="tree">
			<Canvas ref="canvas" :data="tree" />
		</div>
	</div>
</template>

<script>
import Canvas from './Canvas.vue';
import NodeForm from './NodeForm.vue';
import { mapActions, mapGetters } from "vuex";

export default {
	components: {
		Canvas,
		NodeForm
	},
	computed: {
		...mapGetters("maps", {
			getMaps: "getMaps",
			getCurrentTree: "getCurrentTree",
			getCurrentMap: "getCurrentMap",
			getStatus: "getStatus"
		}),
		...mapGetters("select", {
			getCurrentNode: "getCurrentNode"
		})
	},
	methods: {
		updateTree() {
			//this.tree.children.push()
			//this.$refs.canvas.a();
			this.$refs.canvas.updateFromSky({
				name: "",
				color: "#7F00FF",
				children: [],
				map_id: this.getCurrentMap.id,
				parent_id: this.getCurrentNode.data.id
			});
		},
		updateSelected() {
			this.$refs.canvas.updateSelected()
		},
		...mapActions("maps", {
			getCardTree: "getCardTree",
		}),
		async getTree() {
			await this.getCardTree();
			if (this.getStatus == "success") {
				this.tree = this.getCurrentTree
			}
		},
	},
	
	mounted() {
		this.getTree()
	},
	data() {
		return {
			maps: this.getMaps,
			tree: null,
			// 	name: "text text text text text text text text text",
			// 	color: "#ff0000",
			// 	children: [
			// 		{
			// 			name: "123490\n1123456711234 567890112345678901",
			// 			children: [{
			// 				name: "b",
			// 				color: "#32a852",
			// 				children: []
			// 			},
			// 			{
			// 				name: "c",
			// 				children: []
			// 			},
			// 			{
			// 				name: "c",
			// 				children: []
			// 			}]
			// 		},
			// 		{
			// 			name: "c",
			// 			children: [{
			// 				name: "b",
			// 				children: []
			// 			},
			// 			{
			// 				name: "c",
			// 				children: []
			// 			}]
			// 		}
			// 	]
			// }
		};
	},
}
</script>