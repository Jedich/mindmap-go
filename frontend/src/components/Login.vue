<template>
	<div class="container">
		<div class="row">
			<div class="col">
			</div>
			<div class="col-8 card" style="margin:20px; padding: 10px">
				<h4>Login Form IIT</h4>
				<div class="mb-3">
					<label class="form-label">Email</label>
					<input type="email" class="form-control" v-model="loginData.email" />
					<div class="alert alert-danger" role="alert"
						v-if="isValidationError('Login') && 'email' in getErrors.data">
						{{ getErrors.data.email }}
					</div>
				</div>
				<div class="mb-3">
					<label class="form-label">Password</label>
					<input type="password" class="form-control" v-model="loginData.password" />
					<div class="alert alert-danger" role="alert"
						v-if="isValidationError('Login') && 'password' in getErrors.data">
						{{ getErrors.data.password }}
					</div>
				</div>
				<div class="alert alert-danger" role="alert" v-if="checkError('Login')">
					{{ getErrors.data }}
				</div>
				<button type="button" class="btn btn-outline-primary" @click="login()">
					Submit
				</button>
			</div>
			<div class="col">
			</div>
		</div>

		<div class="row">
			<div class="col">
			</div>
			<div class="col-8 card" style="margin:20px; padding: 10px">
				<h4>Register Form</h4>
				<div class="mb-3">
					<label class="form-label">Username</label>
					<input type="text" class="form-control" v-model="regData.username" />
					<div class="alert alert-danger" role="alert"
						v-if="isValidationError('Reg') && 'username' in getErrors.data">
						{{ getErrors.data.username }}
					</div>
				</div>
				<div class="mb-3">
					<label class="form-label">Email</label>
					<input type="email" class="form-control" v-model="regData.email" />
					<div class="alert alert-danger" role="alert"
						v-if="isValidationError('Reg') && 'email' in getErrors.data">
						{{ getErrors.data.email }}
					</div>
				</div>
				<div class="mb-3">
					<label class="form-label">Password</label>
					<input type="password" class="form-control" v-model="regData.password" />
					<div class="alert alert-danger" role="alert"
						v-if="isValidationError('Reg') && 'password' in getErrors.data">
						{{ getErrors.data.password }}
					</div>
				</div>
				<div class="alert alert-danger" role="alert" v-if="checkError('Reg')">
					{{ getErrors.data }}
				</div>
				<button type="button" class="btn btn-outline-primary" @click="register()">
					Submit
				</button>
			</div>
			<div class="col">
			</div>
		</div>
	</div>
</template>

<script>
import { mapActions, mapGetters } from "vuex";
export default {
	data() {
		return {
			loginData: {
				email: "",
				password: "",
			},
			regData: {
				username: "",
				email: "",
				password: "",
			}
		};
	},
	computed: {
		...mapGetters("auth", {
			getApiStatus: "getApiStatus",
			getUserProfile: "getUserProfile",
			getErrors: "getErrors"
		}),
	},
	methods: {
		checkError: function (type) {
			return this.getApiStatus === ('failed' + type) && this.getErrors.type === ""
		},
		isValidationError: function (type) {
			return this.getApiStatus === ('failed' + type) && this.getErrors.type === 'validation'
		},
		...mapActions("auth", {
			actionLoginApi: "loginApi",
			actionRegisterApi: "registerApi"
		}),
		async login() {
			console.log(this.email, this.password);
			const payload = this.loginData;
			await this.actionLoginApi(payload);
			if (this.getApiStatus == "successLogin") {
				this.$router.push("/app");
			}
		},
		async register() {
			console.log(this.email, this.password);
			const payload = this.regData;
			await this.actionRegisterApi(payload);
			if (this.getApiStatus == "successReg") {
				this.$router.push("/app");
			}
		},
	},
};
</script>