import axios from 'axios';
import Cookies from 'js-cookie';

const state = () => ({
	maps: null,
	tabbedMaps: null,
	currentTree: null,
	currentMap: null,
	error: null,
	status: ""
});

const getters = {
	getMaps(state) {
		return state.maps;
	},
	getCurrentTree(state) {
		return state.currentTree;
	},
	getCurrentMap(state) {
		return state.currentMap;
	},
	getStatus(state) {
		return state.status;
	}
};

const actions = {
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
	async getCardTree({ commit, getters, dispatch }) {
		if(getters.getMaps === null) {
			await dispatch('loadMaps');
			console.log("Loaded maps from db", getters.getMaps)
		}
		if (!getters.getCurrentMap) {
			var map = getters.getMaps[0]
			map.selected = true
			commit('setCurrentMap', map)
		}
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
			map.tree = getters.getCurrentTree
		} else {
			commit("setCurrentTree", map.tree);
			//getters.getCurrentTree = getters.getCurrentMap.tree
		}
	},
};


const mutations = {
	setMaps(state, data) {
		state.maps = data
	},
	setCurrentTree(state, data) {
		state.currentTree = data
	},
	setCurrentMap(state, data) {
		state.currentMap = data
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