import { createStore } from "vuex";
import authModule from './modules/auth';
import mapsModule from './modules/maps';
import selectModule from './modules/select';
import VuexPersistence from 'vuex-persist';

const vuexLocal = new VuexPersistence({
	storage: window.localStorage,
	reducer: (state) => ({
		maps: {
			mapsMap: state.maps.mapsMap,
			tabs: state.maps.tabbedMaps ? Object.keys(state.maps.tabbedMaps) : null,
			currentMapID: state.maps.currentMap ? state.maps.currentMap.id : null,
		}
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