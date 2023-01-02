<template>
	<div class="sidebar adiv">
		<p>adasdasd</p>
		<div v-if="getCurrentNode">
			<NodeForm v-on:updateSelection="updateSelected" />
		</div>
	</div>

	<div class="content adiv">
		<ul class="nav nav-tabs">
			<li class="nav-item" v-for="map in getTabs.sort(fn)">
				<p class="nav-link" :class="{ 'active': map.selected }" @click="select(map)" aria-current="page">{{map.name}}
					<input class="btn btn-outline-danger x" type="button" @click.self.stop="closeTabHere(map)" value="">
				</p>
			</li>
			<li class="nav-item dropdown">
				<a class="nav-link" data-bs-toggle="dropdown" @click="loadMaps" role="button"
					aria-expanded="false">+</a>
				<ul class="dropdown-menu mt-0">
					<li v-for="map in getUnselectedMaps()"><a class="dropdown-item" @click="select(map)" href="#">{{map.name}}</a>
					</li>
					<li v-if="getUnselectedMaps().length > 0">
						<hr class="dropdown-divider">
					</li>
					<li><a class="dropdown-item" href="#">New Map</a></li>
				</ul>
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
import * as d3 from "d3";

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
			getStatus: "getStatus",
			getTabs: "getTabs",
			getTabMap: "getTabMap"
		}),
		...mapGetters("select", {
			getCurrentNode: "getCurrentNode"
		})
	},
	methods: {
		fn: (a, b) => {
			return a.order - b.order
		},
		getUnselectedMaps() {
			if (!this.getCurrentMap) {
				return this.getMaps
			}
			if (this.getMaps) {
				return this.getMaps.filter(prop => this.getTabMap[prop.id] === undefined);// this.getTabMap.has(prop.id)); // prop.id !== this.getCurrentMap.id)
			}
			return []
		},
		updateTree() {
			//this.tree.children.push()
			//this.$refs.canvas.a();
			this.$refs.canvas.updateFromSky({
				name: "",
				color: "#7f00ff",
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
			selectMap: "selectMap",
			closeTab: "closeTab"
		}),
		async loadMaps() {
			await this.loadMaps
		},
		async select(map) {
			console.log("me too")
			if (map.selected) {
				return
			}
			await this.selectMap(map);
			d3.selectAll("g").remove();
			this.tree = map.tree;
		},
		closeTabHere(map) {
			if (map.id === this.getCurrentMap.id) {
				this.tree = null
			}
			this.closeTab(map)
			if(this.getCurrentMap) {
				this.tree = this.getCurrentMap.tree
			}
		}
	},

	mounted() {
		console.log(this.getUnselectedMaps())
	},
	data() {
		return {
			maps: this.getMaps,
			tree: null
		};
	},
}
</script>