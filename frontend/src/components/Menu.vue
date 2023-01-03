<template>
	<div class="sidebar adiv">
		<p></p>
		<div style="margin:10px;" v-if="getCurrentMap">Map name: <input style="width:100%;" class="text-field form-control" :value="getCurrentMap.name" @input="updateMapName" :maxlength="20" />
			<button style="margin-top:5px; width:100%;" class="btn btn-outline-success" v-if="getCurrentMap.updated" v-on:click="updateThisMap">Save</button></div>
		<hr />
		<div v-if="getCurrentNode">
			<NodeForm v-on:updateSelection="updateSelected" />
		</div>
	</div>

	<div class="content adiv">
		<ul class="nav nav-tabs" style="margin-left:5px; padding-top: 5px;">
			<li class="nav-item" v-for="map in getTabs.sort((a, b) => a.order - b.order)">
				<p class="nav-link" :class="{ 'active': map.selected }" @click="select(map)" aria-current="page">
					{{ map.name }}
					<input class="btn btn-outline-danger x" type="button" @click.self.stop="closeTabHere(map)" value="">
				</p>
			</li>
			<li class="nav-item dropdown">
				<a class="nav-link" data-bs-toggle="dropdown" @click="loadMaps" role="button"
					aria-expanded="false">+</a>
				<ul class="dropdown-menu mt-0">
					<li v-for="map in getUnselectedMaps()"><a class="dropdown-item"
							@click="select(map)">{{ map.name }}</a>
					</li>
					<li v-if="getUnselectedMaps().length > 0">
						<hr class="dropdown-divider">
					</li>
					<li><a class="dropdown-item" @click="createMap()">New Map</a></li>
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
			getCurrentMap: "getCurrentMap",
			getTabs: "getTabs",
			getTabMap: "getTabMap",
			isError: "isError"
		}),
		...mapGetters("select", {
			getCurrentNode: "getCurrentNode"
		})
	},
	methods: {
		...mapActions("maps", {
			getCardTree: "getCardTree",
			selectMap: "selectMap",
			newMap: "newMap",
			closeTab: "closeTab",
			updateMap: "updateMap", 
		}),
		updateMapName(event) {
			this.getCurrentMap.name = event.target.value;
			this.getCurrentMap.updated = true;
		},
		async updateThisMap() {
			var map = this.getCurrentMap
			var payload = {
				id: map.id,
				name: map.name,
				desc: map.desc
			};
			console.log(payload)
			await this.updateMap(payload)
		},
		getUnselectedMaps() {
			if (!this.getCurrentMap) {
				return this.getMaps ? this.getMaps : []
			}
			if (this.getMaps) {
				return this.getMaps.filter(prop => this.getTabMap[prop.id] === undefined);
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
		async loadMaps() {
			await this.loadMaps
		},
		async select(map) {
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
			if (this.getCurrentMap) {
				this.tree = this.getCurrentMap.tree
			}
		},
		async createMap() {
			await this.newMap();
			if (!this.isError) {
				this.select(this.getCurrentMap);
			}
		}
	},

	mounted() {
		if (this.getCurrentMap) {
			this.tree = this.getCurrentMap.tree
		}
	},
	data() {
		return {
			maps: this.getMaps,
			tree: null
		};
	},
}
</script>