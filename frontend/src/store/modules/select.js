import * as d3 from "d3";
import axios from 'axios';
import Cookies from 'js-cookie';

const state = () => ({
	selectedNode: null,
});

const getters = {
	getCurrentNode(state) {
		return state.selectedNode;
	},
};

const actions = {
	select({ commit, getters }, node) {
		commit('setCurrentNode', node)
		node.s
			.select('rect')
			.attr("stroke", d => { if (!d.data.color) { d.data.color = "#ffa500" } return d.data.color; })
			.attr("stroke-dasharray", "15,5");
		//blink()
		node.s
			.select('circle.create')
			.style("visibility", "visible")
			.transition()
			.attr("cy", -15)

		node.s
			.select('circle.hide')
			.transition()
			.attr("cy", 15)
	},
	deselect({ commit, getters, dispatch }) {
		var selectedNode = getters.getCurrentNode
		if (selectedNode.data.created) {
			dispatch('commitNode', selectedNode.data);
		}
		if (selectedNode.data.updated) {
			dispatch('updateNode', selectedNode.data);
		}
		selectedNode.s
			.select('rect')
			.style("fill", d => d._children ? "#fff" : "#eee")
			.attr("stroke", d => { if (!d.data.color) { d.data.color = "#ffa500" } return d.data.color; })
			.attr("stroke-dasharray", null);

		selectedNode.s
			.select('circle.create')
			.transition()
			.attr("cy", 0)
			.on('end', function () {
				d3.select(this).style("visibility", "hidden");
			});

		selectedNode.s
			.select('circle.hide')
			.transition()
			.attr("cy", 0)

		commit('setCurrentNode', null)
	},
	async commitNode({ commit }, payload) {
		if (!payload.created) {
			return
		}
		payload.created = null
		const response = await axios
			.post("/api/cards/", payload,
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
			payload.id = response.data.data.id;
		}
	},
	async uploadFile({ commit, getters }, payload) {
		var formData = new FormData();
		formData.append("uploadFile", payload.file);
		for (var key of formData.entries()) {
			console.log(key[0] + ', ' + key[1]);
		}
		const response = await axios
			.post("/api/cards/file/" + getters.getCurrentNode.data.id, formData,
				{
					headers: {
						'Content-Type': 'multipart/form-data',
						'Authorization': `Bearer ${Cookies.get("token")}`
					}
				})
			.catch((err) => {
				console.log(err)
			});
		if (response && response.data) {
			console.log(response.data)
			getters.getCurrentNode.data.file = payload.raw
		}
	},
	async updateNode({ commit }, payload) {
		if (!payload.updated) {
			return
		}
		payload.updated = null
		const response = await axios
			.patch("/api/cards/", payload,
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
		}
	},
	async deleteNode({ commit }, node) {
		const response = await axios
			.delete("/api/cards/" + node.id,
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
		}
	}
};


const mutations = {
	setCurrentNode(state, data) {
		state.selectedNode = data
	},
};

export default {
	namespaced: true,
	state,
	getters,
	actions,
	mutations,
};