<template>
    <input :value="store.selectedNode.data()[0].data.name" @change="updateMyValue" />
</template>

<script>
import { store } from '../store';
import * as d3 from "d3";
import Canvas from './Canvas.vue';

export default {
    components: {
		Canvas,
    },
    data() {
        return {
            store,
        }
    },
    methods: {
        updateMyValue (event) {
            store.selectedNode.data()[0].data.name = event.target.value
            console.log(store.selectedNode.data()[0].data.name)
        },

    }
}
</script>