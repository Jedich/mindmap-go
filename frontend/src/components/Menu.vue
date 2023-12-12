<template>
	<div class="outer">
		<div class="user-info-card">
			<h1>User information</h1>
			<p>Name: {{ getUserProfile.first_name }} {{ getUserProfile.last_name }}</p>
			<p>ID: {{ getUserProfile.id }}</p>
			<p>Username: {{ getUserProfile.account.username }}</p>
			<p>Email: {{ getUserProfile.account.email }}</p>
		</div>
	</div>
</template>

<script>
import { mapActions, mapGetters } from "vuex";
import * as d3 from "d3";

export default {
	components: {
	},
	computed: {
		...mapGetters("auth", {
			getUserProfile: "getUserProfile"
		}),
	},
	methods: {
	},

	mounted() {
	},
	data() {
		return {
		};
	},
}
</script>