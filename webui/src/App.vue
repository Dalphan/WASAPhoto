<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>
<script>
export default {
	data: function() {
		return {
			logged: null,
			userLink: null,
			currentUsername: null,
		}
	},
	watch:{
		$route (to, from){
			this.logged = this.$getCurrentId(); // Check if user is logged, if yes render the navbar
			this.currentUsername = this.$getCurrentUsername();
			this.userLink = "/user/" + this.currentUsername;
		}
	},	
	methods: {
		logout() {
			localStorage.removeItem("token");
			localStorage.removeItem("username");
			this.$setAuth();
			this.$router.push({name: 'Login'});
		}
	},
	mounted() {
		// localStorage.token = null
		this.$setAuth();

		this.userLink = "/user/" + this.$getCurrentUsername();

		this.$axios.interceptors.response.use(response => {
			return response;
		}, error => {
				// If the user is Unauthorized, redirect to login
			if (error.response.status === 401) {
				this.$router.push({ name: 'Login' })
				return;
			}
			else 
				return Promise.reject(error) // Leave other error handlers
		});
	}
}
</script>

<template>

	<header class="navbar navbar-dark sticky-top bg-dark flex-md-nowrap p-0 shadow">
		<a class="navbar-brand col-md-3 col-lg-2 me-0 px-3 fs-6" href="#/">WASAPhoto</a>
		<button class="navbar-toggler position-absolute d-md-none collapsed" type="button" data-bs-toggle="collapse" data-bs-target="#sidebarMenu" aria-controls="sidebarMenu" aria-expanded="false" aria-label="Toggle navigation">
			<span class="navbar-toggler-icon"></span>
		</button>
	</header>

	<!--<div class="container-fluid">
		<div class="row">
			<nav id="sidebarMenu" class="col-md-3 col-lg-2 d-md-block bg-light sidebar collapse">
				<div class="position-sticky pt-3 sidebar-sticky">
					<h6 class="sidebar-heading d-flex justify-content-between align-items-center px-3 mt-4 mb-1 text-muted text-uppercase">
						<span>General</span>
					</h6>
					<ul class="nav flex-column">
						<li class="nav-item">
							<RouterLink to="/home" class="nav-link">
								<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#home"/></svg>
								Home
							</RouterLink>
						</li>
						<li class="nav-item">
							<RouterLink to="/user" class="nav-link">
								<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#user"/></svg>
								Profile
							</RouterLink>
						</li>
						<li class="nav-item">
							<RouterLink to="/home" class="nav-link">
								<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#key"/></svg>
								Menu item 2
							</RouterLink>
						</li>
						<li class="nav-item" v-if="this.$getCurrentId()">
							<div class="nav-link" role="button" @click="logout">
								<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#log-out"/></svg>
								Log out
							</div>
						</li>
					</ul>

					<h6 class="sidebar-heading d-flex justify-content-between align-items-center px-3 mt-4 mb-1 text-muted text-uppercase">
						<span>Secondary menu</span>
					</h6>
					<ul class="nav flex-column">
						<li class="nav-item">
							<RouterLink :to="'/some/' + 'variable_here' + '/path'" class="nav-link">
								<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#file-text"/></svg>
								Item 1
							</RouterLink>
						</li>
					</ul>
				</div>
			</nav>

			<main class="col-md-9 ms-sm-auto col-lg-10 px-md-4">
				<RouterView />
			</main>
		</div>
	</div>-->
	<div class="container-fluid">
		<div class="row">
			<div v-if="logged"> <!-- render this if the user is logged-->
				<nav id="sidebarMenu" class="col-md-3 col-lg-2 d-md-block bg-light sidebar collapse">
					<div class="position-sticky pt-3 sidebar-sticky">
						<h6 class="sidebar-heading d-flex justify-content-between align-items-center px-3 mt-4 mb-1 text-muted text-uppercase">
							<span>General</span>
						</h6>
						<ul class="nav flex-column">
							<li class="nav-item">
								<RouterLink to="/home" class="nav-link">
									<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#home"/></svg>
									Home
								</RouterLink>
							</li>
							<li class="nav-item">
								<RouterLink :to="userLink" class="nav-link">
									<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#user"/></svg>
									Profile
								</RouterLink>
							</li>
							<li class="nav-item">
								<RouterLink to="/search" class="nav-link">
									<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#key"/></svg>
									Search
								</RouterLink>
							</li>
							<li class="nav-item">
								<div class="nav-link" role="button" @click="logout">
									<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#log-out"/></svg>
									Log out
								</div>
							</li>
						</ul>

						<h6 class="sidebar-heading d-flex justify-content-between align-items-center px-3 mt-4 mb-1 text-muted text-uppercase">
							<span>Logged as {{ currentUsername }}</span>
						</h6>

						<!-- <h6 class="sidebar-heading d-flex justify-content-between align-items-center px-3 mt-4 mb-1 text-muted text-uppercase">
							<span>Secondary menu</span>
						</h6>
						<ul class="nav flex-column">
							<li class="nav-item">
								<RouterLink :to="'/some/' + 'variable_here' + '/path'" class="nav-link">
									<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#file-text"/></svg>
									Item 1
								</RouterLink>
							</li>
						</ul> -->
					</div>
				</nav>
			

				<main class="col-md-9 ms-sm-auto col-lg-10 px-md-4">
					<RouterView />
				</main>
			</div>
			<div v-else> <!-- render this if the user is not logged -->
				<main class="col-md-12 ms-sm-auto col-lg-12 px-md-4">
					<RouterView />
				</main>
			</div> 
		</div>
	</div>
</template>

<style>
</style>
