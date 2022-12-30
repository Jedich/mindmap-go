import { createStore } from "vuex";
import authModule from './modules/auth';
import mapsModule from './modules/maps';

const store = createStore({
	modules: {
		auth: authModule,
		maps: mapsModule
	}
});

export default store;