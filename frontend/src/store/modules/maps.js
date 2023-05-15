import axios from 'axios';
import Cookies from 'js-cookie';
import * as d3 from "d3";



const state = () => ({
	mapsMap: Object.create(null),
	currentMapID: null,
	tabs: null,

	maps: null,
	tabbedMaps: Object.create(null),
	currentMap: null,

	isError: false,
	order: 0,
});

const getters = {
	getMaps(state) {
		return state.maps;
	},
	getTabs(state) {
		return Object.values(state.tabbedMaps);//Array.from(state.tabbedMaps.values());
	},
	getTabMap(state) {
		return state.tabbedMaps;
	},
	getCurrentMap(state) {
		return state.currentMap;
	},
	isError(state) {
		return state.isError
	}
};

const actions = {
	initState({ state }) {
		state.maps = Object.keys(state.mapsMap).length !== 0 ? Object.values(state.mapsMap) : null;
		state.tabbedMaps = Object.create(null);
		if (state.tabs) {
			state.tabs.forEach(element => {
				state.tabbedMaps[element] = state.mapsMap[element];
			});
		}
		state.currentMap = Object.keys(state.mapsMap).length !== 0 ? state.mapsMap[state.currentMapID] : null;
	},
	async loadMaps({ commit, getters }) {
		const response = await axios
			.get("/api/maps/",
				{
					headers: {
						'Authorization': `Bearer ${Cookies.get("token")}`
					}
				})
			.catch((err) => {
				console.log(err)
			});
		if (response && response.data) {
			console.log(response.data)
			commit('setMaps', response.data.data)
		}
	},
	async selectMap({ commit, getters, dispatch, state }, map) {
		if (getters.getMaps === null) {
			await dispatch('loadMaps');
			console.log("Loaded maps from db", getters.getMaps);
		}
		console.log(getters.getCurrentMap)
		if (getters.getCurrentMap) {
			getters.getCurrentMap.selected = false;
		}
		map.selected = true;
		if (!(map.id in getters.getTabMap)) {
			map.order = state.order;
			state.order++;
			if (state.order > 50) {
				state.order = 0
			}
		}
		commit('setTab', map);
		commit('setCurrentMap', map);
		await dispatch('getCardTree');
	},
	closeTab({ commit, getters, dispatch }, map) {
		map.selected = false;
		commit('deleteTab', map);
		if (getters.getCurrentMap == map) {
			d3.selectAll("g").remove();
			if (getters.getTabMap && Object.keys(getters.getTabMap).length !== 0) {
				dispatch('selectMap', getters.getTabs[getters.getTabs.length - 1])
			} else {
				commit('setCurrentMap', null);
				commit('setCurrentTree', null);
			}
		}
	},
	async newMap({ commit, dispatch, getters }) {
		const response = await axios
			.post("/api/maps/", null,
				{
					headers: {
						'Authorization': `Bearer ${Cookies.get("token")}`
					}
				})
			.catch((err) => {
				console.log(err);
				commit('error', true);
			});
		if (response && response.data) {
			commit('error', false);
			commit('setTab', response.data.data);
			console.log(getters.getCurrentMap)
			if (getters.getCurrentMap) {
				getters.getCurrentMap.selected = false;
			}
			commit('setCurrentMap', response.data.data);

			commit('addMap', response.data.data);
		}
	},
	async updateMap({ commit }, payload) {
		const response = await axios
			.patch("/api/maps/", payload,
				{
					headers: {
						'Authorization': `Bearer ${Cookies.get("token")}`
					}
				})
			.catch((err) => {
				console.log(err);
			});
		if (response && response.data) {
			console.log(response.data);
			commit('updateMap', payload);
		}
	},
	async getCardTree({ commit, getters, dispatch }) {
		var map = getters.getCurrentMap
		if (!map.tree) {
			const response = await axios
				.get("/api/cards/" + map.id,
					{
						headers: {
							'Authorization': `Bearer ${Cookies.get("token")}`
						}
					})
				.catch((err) => {
					console.log(err);
				});
			if (response && response.data) {
				console.log(response.data)
				commit('setCurrentTree', response.data.data)
			}
			map.tree = getters.getCurrentMap.tree
		} else {
			commit("setCurrentTree", map.tree);
		}
	},
	saveMapToImage() {
		var svg = document.getElementsByTagName("svg")[0];

		var serializer = new XMLSerializer();
		var source = serializer.serializeToString(svg);

		if(!source.match(/^<svg[^>]+xmlns="http\:\/\/www\.w3\.org\/2000\/svg"/)){
			source = source.replace(/^<svg/, '<svg xmlns="http://www.w3.org/2000/svg"');
		}
		if(!source.match(/^<svg[^>]+"http\:\/\/www\.w3\.org\/1999\/xlink"/)){
			source = source.replace(/^<svg/, '<svg xmlns:xlink="http://www.w3.org/1999/xlink"');
		}

		source = '<?xml version="1.0" standalone="no"?>\r\n' + source;

		var url = "data:image/svg+xml;charset=utf-8,"+encodeURIComponent(source);

		var iframe = `<a href="${url}" id="link" download>Save map</a>
		<iframe src='${url}' frameborder="0" style="border:0; top:0px; left:0px; bottom:0px; right:0px; width:100%; height:98%;" allowfullscreen>
		</iframe>`
		var x = window.open();
		x.document.open();
		x.document.write(iframe);
		x.document.close();
	}
};

const mutations = {
	setMaps(state, data) {
		state.mapsMap = data
		state.maps = Object.values(data)
	},
	addMap(state, data) {
		state.mapsMap[data.id] = data
		state.maps.push(data);
	},
	updateMap(state, data) {
		var map = state.mapsMap[data.id];
		map.updated = false;
		map.name = data.name;
		map.desc = data.desc;
		state.maps = Object.values(state.mapsMap)
	},
	setTab(state, data) {
		state.tabbedMaps[data.id] = data;
	},
	deleteTab(state, data) {
		delete state.tabbedMaps[data.id];
	},
	setCurrentTree(state, data) {
		if (state.currentMap != null) {
			state.currentMap.tree = data;
		}
	},
	setCurrentMap(state, data) {
		state.currentMap = data
		if (state.currentMap != null) {
			state.currentMapID = data.id
		}
	},
	error(state, data) {
		state.isError = data;
	}
};

export default {
	namespaced: true,
	state,
	getters,
	actions,
	mutations
};