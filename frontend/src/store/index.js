import { createStore } from "vuex";
import authModule from './modules/auth';
import mapsModule from './modules/maps';
import selectModule from './modules/select';

const store = createStore({
	modules: {
		auth: authModule,
		maps: mapsModule,
		select: selectModule
	}
});

export default store;