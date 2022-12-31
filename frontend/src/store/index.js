import { createStore } from "vuex";
import authModule from './modules/auth';
import mapsModule from './modules/maps';
import selectModule from './modules/select';
import VuexPersistence from 'vuex-persist';

const vuexLocal = new VuexPersistence({
	storage: window.localStorage,
	reducer: (state) => ({
		maps: state.maps
	})
});

const store = createStore({
	modules: {
		auth: authModule,
		maps: mapsModule,
		select: selectModule
	},
	plugins: [vuexLocal.plugin]
});

export default store;