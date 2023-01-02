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

	error: null,
	status: "",
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
	getStatus(state) {
		return state.status;
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
				commit('setStatus', "failed")
				commit("setErrors", err.response.data.errors.data);
			});
		if (response && response.data) {
			console.log(response.data)
			commit('setStatus', "success")
			commit('setMaps', response.data.data)
		}
	},
	async selectMap({ commit, getters, dispatch, state }, map) {
		if (getters.getMaps === null) {
			await dispatch('loadMaps');
			console.log("Loaded maps from db", getters.getMaps);
		}
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
		await dispatch('getCardTree')
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
					console.log(err)
					commit('setStatus', "failed")
					commit("setErrors", err.response.data.errors.data);
				});
			if (response && response.data) {
				console.log(response.data)
				commit('setStatus', "success")
				commit('setCurrentTree', response.data.data)
			}
			map.tree = getters.getCurrentMap.tree
		} else {
			commit("setCurrentTree", map.tree);
		}
	},
};


const mutations = {
	setMaps(state, data) {
		state.mapsMap = data
		state.maps = Object.values(data)
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
	setErrors(state, data) {
		state.error = data
	},
	setStatus(state, data) {
		state.status = data;
	}
};

export default {
	namespaced: true,
	state,
	getters,
	actions,
	mutations
};