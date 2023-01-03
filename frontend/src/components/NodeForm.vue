<template>
	<form style="margin:10px;">
		<div style="display: block;">Name: <span class="input-group-addon limit" v-text="(nameLength - getCurrentNode.data.name.length)"></span>
			<textarea class="form-control" :value="getCurrentNode.data.name" @input="updateName" :maxlength="nameLength" rows="2" />
		</div>
		<div>Description:
			<textarea class="form-control" :value="getCurrentNode.data.text_data" @input="updateText" rows="5" />
		</div>
		<div>Stroke color: 
			<input type="color" style="width:100%" :value="getCurrentNode.data.color" @input="throttledColor" /></div>
		<div v-if="getCurrentNode.data.created">
			<button style="margin-top:5px; width:100%;" class="btn btn-outline-success" v-on:click="deselect">Create</button>
		</div>
		<div v-if="getCurrentNode.data.updated">
			<button style="margin-top:5px; width:100%;" class="btn btn-outline-success" v-on:click="deselect">Save</button>
		</div>
	</form>
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
			if (!this.getCurrentNode.data.created) {
				this.getCurrentNode.data.updated = true;
			}
			this.$emit("updateSelection")
		},
		updateText(event) {
			this.getCurrentNode.data.text_data = event.target.value
			if (!this.getCurrentNode.data.created) {
				this.getCurrentNode.data.updated = true;
			}
		},
		throttledColor: throttle(function (event) {
			this.getCurrentNode.data.color = event.target.value
			this.getCurrentNode.s.select('rect').transition().duration(100).attr("stroke", d => d.data.color);
			if (!this.getCurrentNode.data.created) {
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