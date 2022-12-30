import * as d3 from "d3";

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
            .attr("stroke", d => { if (!d.data.color) { d.data.color = "#FFA500" } return d.data.color; })
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
    deselect({ commit, getters }) {
        var selectedNode = getters.getCurrentNode
        console.log(selectedNode)
        selectedNode.s
            .select('rect')
            .style("fill", d => d._children ? "#fff" : "#eee")
            .attr("stroke", d => { if (!d.data.color) { d.data.color = "#FFA500" } return d.data.color; })
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