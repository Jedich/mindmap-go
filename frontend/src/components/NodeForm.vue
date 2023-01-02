<template>
	<div>About:</div>
	<textarea :value="getCurrentNode.data.name" @input="updateName" :maxlength="nameLength" rows="3" />
	<span class="input-group-addon" v-text="(nameLength - getCurrentNode.data.name.length)"></span>
	<div>Description:</div>
	<textarea :value="getCurrentNode.data.text_data" @input="updateText" rows="5" />
	<div>Stroke color: <input type="color" :value="getCurrentNode.data.color" @input="throttledColor" /></div>
	<div v-if="getCurrentNode.data.created">
		<button v-on:click="deselect">Commit</button>
	</div>
	<div v-if="getCurrentNode.data.updated">
		<button v-on:click="deselect">Save</button>
	</div>

</template>

<script>
import * as d3 from "d3";
import Canvas from './Canvas.vue';
import throttle from 'lodash/throttle';
import { mapActions, mapGetters } from "vuex";
export default {
	computed: {
		...mapGetters("select", {
			getCurrentNode: "getCurrentNode",
		}),
	},
	components: {
		Canvas,
	},
	emits: ["updateSelection"],
	data() {
		return {
			nameLength: 45,
		}
	},
	watch: {
		color: function () {
			this.throttledColor()
		}
	},
	methods: {
		updateName(event) {
			this.getCurrentNode.data.name = event.target.value
			if(!this.getCurrentNode.data.created) {
				this.getCurrentNode.data.updated = true;
			}
			this.$emit("updateSelection")
		},
		updateText(event) {
			this.getCurrentNode.data.text_data = event.target.value
			if(!this.getCurrentNode.data.created) {
				this.getCurrentNode.data.updated = true;
			}
		},
		throttledColor: throttle(function (event) {
			this.getCurrentNode.data.color = event.target.value
			this.getCurrentNode.s.select('rect').transition().duration(100).attr("stroke", d => d.data.color);
			if(!this.getCurrentNode.data.created) {
				this.getCurrentNode.data.updated = true;
			}
		}, 100),
		...mapActions("select", {
			select: "select",
			deselect: "deselect",
			commitNode: "commitNode"
		}),
	}
}
</script>