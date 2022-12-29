<template>
	<div>
		<h4>Login Form</h4>
		<form>
			<div class="mb-3">
				<label for="txtEmail" class="form-label">Email</label>
				<input type="text" class="form-control" id="txtEmail" aria-describedby="emailHelp" v-model="email" />
				<div v-if="getLoginApiStatus === 'failed'  && getErrors.type === 'validation'">
					{{ getErrors.data.email }}
				</div>
			</div>
			<div class="mb-3">
				<label for="txtPassword" class="form-label">Password</label>
				<input type="password" class="form-control" id="txtPassword" v-model="password" />
				<div v-if="getLoginApiStatus === 'failed' && getErrors.type === 'validation'">
					{{ getErrors.data.password }}
				</div>
			</div>
			<div v-if="checkError()">
				{{ getErrors }}
			</div>
			<button type="button" class="btn btn-primary" @click="login()">
				Submit
			</button>
		</form>
		<ul class="navbar-nav me-auto mb-2 mb-lg-0" v-if="getUserProfile.id !== 0">
			<li>
				<h5>
					<span class="badge bg-primary">{{ getUserProfile.username }}</span>
				</h5>
			</li>
			<li class="nav-item">
				<router-link to="/app" class="nav-link">To map</router-link>
			</li>
			<li>
				<!-- <span @click="">Logout</span> -->
			</li>
		</ul>
	</div>
</template>

<script>
import { mapActions, mapGetters } from "vuex";
export default {
	data() {
		return {
			email: "",
			password: "",
		};
	},
	computed: {
		...mapGetters("auth", {
			getLoginApiStatus: "getLoginApiStatus",
			getUserProfile: "getUserProfile",
			getErrors: "getErrors"
		}),
	},
	methods: {
		checkError: function () {
			console.log(Array.isArray(this.getErrors))
			console.log((this.getErrors))
			return this.getLoginApiStatus === 'failed' && this.getErrors.type === ""
		},
		...mapActions("auth", {
			actionLoginApi: "loginApi",
		}),
		async login() {
			console.log(this.email, this.password);
			const payload = {
				email: this.email,
				password: this.password,
			};
			await this.actionLoginApi(payload);
			if (this.getLoginApiStatus == "success") {
				this.$router.push("/app");
			} else {
				//alert("failed")
			}
		},
	},
};
</script>