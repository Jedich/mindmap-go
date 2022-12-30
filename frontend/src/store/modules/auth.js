import axios from 'axios';

const state = () => ({
	loginApiStatus: "",
	userProfile: {
		id: 0
	},
	errors: null
});

const getters = {
	getLoginApiStatus(state) {
		return state.loginApiStatus;
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
				commit("setLoginApiStatus", "failed");
				commit("setErrors", err.response.data.errors);
			});
		if (response && response.data) {
			console.log(response.data)
			commit("setLoginApiStatus", "success");
			commit("setUserProfile", response.data.data);
			commit('maps/setMaps', response.data.data.maps, { root: true })
		}
	},
};

const mutations = {
	setLoginApiStatus(state, data) {
		state.loginApiStatus = data;
	},
	setUserProfile(state, data) {
		state.userProfile = data.user
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