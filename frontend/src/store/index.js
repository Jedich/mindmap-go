import { createStore } from "vuex";
import authModule from './modules/auth';
import VuexPersistence from 'vuex-persist';

const vuexLocal = new VuexPersistence({
	storage: window.localStorage,
	reducer: (state) => ({
		auth: {
			...state.auth
		},
	})
});

const store = createStore({
	modules: {
		auth: authModule
	},
	plugins: [vuexLocal.plugin]
});

export default store;