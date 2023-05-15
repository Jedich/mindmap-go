import axios from 'axios';

const state = () => ({
	apiStatus: "",
	userProfile: null,
	errors: null
});

const getters = {
	getApiStatus(state) {
		return state.apiStatus;
	},
	getUserProfile(state) {
		return state.userProfile;
	},
	getErrors(state) {
		return state.errors;
	}
};

const actions = {
	async loginApi({ commit }, payload) {
		const response = await axios
			.post("/api/auth/login",
				payload, { withCredentials: true, credentials: 'include' })
			.catch((err) => {
				console.log(err)
				commit("setApiStatus", "failedLogin");
				commit("setErrors", err.response.data.errors);
			});
		if (response && response.data) {
			console.log(response.data)
			commit("setApiStatus", "successLogin");
			commit("setUserProfile", response.data.data);
			commit('maps/setMaps', response.data.data.maps, { root: true })
		}
	},
	async registerApi({ commit }, payload) {
		const response = await axios
			.post("/api/auth/register",
				payload, { withCredentials: true, credentials: 'include' })
			.catch((err) => {
				console.log(err)
				commit("setApiStatus", "failedReg");
				commit("setErrors", err.response.data.errors);
			});
		if (response && response.data) {
			console.log(response.data)
			commit("setApiStatus", "successReg");
			commit("setUserProfile", response.data.data.user.account);
			commit('maps/setMaps', response.data.data.maps, { root: true })
		}
	},
};

const mutations = {
	setApiStatus(state, data) {
		state.apiStatus = data;
	},
	setUserProfile(state, data) {
		state.userProfile = data
	},
	setErrors(state, data) {
		state.errors = data
	}
};

export default {
	namespaced: true,
	state,
	getters,
	actions,
	mutations,
};