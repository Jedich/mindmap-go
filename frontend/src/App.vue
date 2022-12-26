<template>
	<div class="header">
		<h3 style="padding-left: 20px" href="#">mindmap-go</h3>
	</div>
	<!-- <nav class="navbar bg-body-tertiary">
		<div class="container-fluid">
			<a class="navbar-brand" href="#">Navbar</a>
		</div>
	</nav> -->
	<div id="app" class="page">
		<div class="sidebar adiv">
			<p>adasdasd</p>
		</div>

		<div class="content adiv">
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

			<div>
				<Canvas />
			</div>
		</div>

	</div>
</template>

<script>
import Canvas from './components/Canvas.vue';
import axios from 'axios';
import { defineComponent } from 'vue'
export default defineComponent({
	name: "App",
	components: {
		Canvas
	},
	data() {
		return {
			users: null,
		};
	},
	methods: {
		getMaps() {
			axios.get("/api/users").then((response) => (this.users = response.data.data));
		},
	},
});
</script>