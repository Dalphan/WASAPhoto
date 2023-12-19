<script>
export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			username: null,
			id: null
		}
	},
	methods: {
		// async refresh() {
		// 	this.loading = true;
		// 	this.errormsg = null;
		// 	try {
		// 		let response = await this.$axios.post("/session", {
		// 			name: 'Maria',
		// 		});
		// 		this.some_data = response.data;
		// 	} catch (e) {
		// 		this.errormsg = e.toString();
		// 	}
		// 	this.loading = false;
		// },
		async login() {
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.post("/session", {
					name: this.username,
				});
				this.id = response.data;
				localStorage.token = this.id
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		},
	},
	mounted() {
		// localStorage.token = 1
		var currentId = this.$getCurrentId()
		// Check if the user was already logged
		if (currentId) {
			this.$router.push("/home")
		}
	}
}
</script>

<template>
    <div id="app">
        <h2>Login</h2>
		<form>
			<div class="form-group">
			<label for="Username">Username</label>
			<input type="text" class="form-control" id="Username" placeholder="Enter username" v-model="username">
			<small id="usernameHelp" class="form-text font-italic text-danger">Username required</small>
		</div>
		<button type="submit" class="btn btn-outline-dark">Log in</button>
		</form>
	</div>
</template>

<style>
</style>
