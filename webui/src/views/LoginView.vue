<script>
export default {
	data: function() {
		return {
			errormsg: "",
			loading: false,
			username: null,
			id: null,
			error: false,
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
			this.errormsg = "";
			this.error = false;
			try {
				let response = await this.$axios.post("/session", {
					name: this.username,
				});

				// TODO: Check username validity?

				if (response.status == 200 || response.status == 201) {
					this.id = response.data;
					localStorage.token = this.id;
					localStorage.username = this.username;
					this.$setAuth();
					this.$router.push({name: 'Home'});
				}			
			} catch (e) {
				const {response} = e; // Get response from error body
				if (response.status == 400) { // The username is missing
					this.error = true;
					this.errormsg = "Username required"
				}
				else if (response.status == 406) { // The username is not valid
					this.error = true;
					this.errormsg = "Username should have from 3 to 16 characters, numbers, . and _"
				}
				// this.errormsg = e.toString();
			}
			this.loading = false;
		},
	},
	mounted() {
		var currentId = this.$getCurrentId();
		// Check if the user was already logged, if yes go to home page
		if (currentId) {
			this.$router.push({name: 'Home'});
		}
	}
}
</script>

<template>
    <!-- <div id="app">
        <h2>Login</h2>
			<div class="form-group">
			<label for="Username">Username</label>
			<input type="text" class="form-control" id="Username" placeholder="Enter username" v-model="username">
			<small id="usernameHelp" class="form-text font-italic text-danger" v-if="error" v-text="errormsg"></small>
		</div>
		<button @click="login" class="btn btn-outline-dark">Log in</button>
	</div> -->
	<div class="form-body">
        <div style="max-width: 100vh;" class="form-inner">
            <div class="mb-3">
                <h1>WASAPhoto Login</h1>
            </div>
			<div style="font-size: 1.3rem;">
				<div class="form-group mb-3">
					<label for="Username">Username</label>
					<input style="text-align: center;" v-model="username" id="Username" type="text" class="form-control" placeholder="Type username...">
					<!-- <div class="input-group-append">
						<button class="input-group-text text-primary send-button" type="button" @click="searchUsers">
							<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#search"/></svg>
						</button>
					</div> -->
				</div>
				<button style="font-size:1.2rem;" class="profile-buttons profile-buttons-primary" type="button" @click="login">
					<svg style="margin-left: 2px;" class="feather"><use href="/feather-sprite-v4.29.0.svg#log-in"/></svg>
					Login
				</button>
			</div>
			<div v-if="errormsg" >
				<hr>
				<ErrorMsg :msg="errormsg"></ErrorMsg>
			</div>
        </div>
    </div>
</template>

<style>
</style>
