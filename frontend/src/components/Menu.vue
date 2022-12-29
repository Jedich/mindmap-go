<template>
	<div>
		<Canvas ref="canvas" :data="tree" />
	</div>
	<ul class="nav nav-tabs">
		<li class="nav-item">
			<p class="nav-link active" aria-current="page">Active <input class="btn btn-outline-danger"
					style="--bs-btn-padding-y: .05rem; --bs-btn-padding-x: .4rem; --bs-btn-font-size: .5rem; --bs-btn-border-radius: 1rem;"
					type="button" value=""></p>
		</li>
		<li class="nav-item">
			<a class="nav-link" href="#">Link</a>
		</li>
		<li class="nav-item">
			<a class="nav-link" href="#">Link</a>
		</li>
		<li class="nav-item">
			<a class="nav-link disabled">Disabled</a>
		</li>
	</ul>
</template>

<script>
import Canvas from './Canvas.vue';

export default {
	components: {
		Canvas,
	},
	methods: {
		updateTree() {
			//this.tree.children.push()
			//this.$refs.canvas.a();
			this.$refs.canvas.updateFromSky({
				name: "toddler",
				color: "violet",
				children: []
			});
		},
		updateSelected() {
			this.$refs.canvas.updateSelected()
		},
	},
	data() {
		return {
			store,
			tree: {
				name: "text text text text text text text text text",
				color: "#ff0000",
				children: [
					{
						name: "123490\n1123456711234 567890112345678901",
						children: [{
							name: "b",
							color: "#32a852",
							children: []
						},
						{
							name: "c",
							children: []
						},
						{
							name: "c",
							children: []
						}]
					},
					{
						name: "c",
						children: [{
							name: "b",
							children: []
						},
						{
							name: "c",
							children: []
						}]
					}
				]
			}
		};
	},
}
</script>