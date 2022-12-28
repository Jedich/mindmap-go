<template>
	<div>About:</div>
	<textarea :value="store.selectedNode.data.name" @input="updateName" :maxlength="nameLength" rows="3" />
	<span class="input-group-addon" v-text="(nameLength - store.selectedNode.data.name.length)"></span>
	<div>Description:</div>
	<textarea rows="5" />
	<div>Stroke color: <input type="color" :value="store.selectedNode.data.color" @input="throttledColor" /></div>
	
</template>

<script>
import { store } from '../store';
import * as d3 from "d3";
import Canvas from './Canvas.vue';
import throttle from 'lodash/throttle';

export default {
	components: {
		Canvas,
	},
	emits: ["updateSelection"],
	data() {
		return {
			store,
			nameLength: 45.
		}
	},
	watch: {
         color: function () {
             this.throttledColor()
         }
    },
	methods: {
		updateName(event) {
			store.selectedNode.data.name = event.target.value
			this.$emit("updateSelection")
		},
		throttledColor: throttle(function (event) {
            store.selectedNode.data.color = event.target.value
			store.selectedNode.s.select('rect').transition().duration(100).attr("stroke", d => d.data.color);
        }, 100)
	}
}
</script>