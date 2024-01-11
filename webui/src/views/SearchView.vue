<script setup>
import ErrorMsg from '../components/ErrorMsg.vue'
</script>
<script>
export default {
	data: function() {
		return {
			errormsg: "",
			loading: false,
			searchUsername: null,
			searched: null,
		}
	},
	components: {
		ErrorMsg,
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
		async searchUsers(){
			this.loading = true;
			this.errormsg = null;
			try {
				var path = `/users?username=${this.searchUsername}`;
				let response = await this.$axios.get(path);

				if (response.status == 200) {
					this.searched = response.data;
				}
			} catch (e) {
				this.errormsg = e.response.data === null ? e.toString() : e.response.data;
			}
			this.loading = false;
		}

	},
	mounted() {

	}
}
</script>

<template>
    <div class="form-body">
        <div class="form-inner">
            <div class="mb-3">
                <h3>Search users</h3>
            </div>
            <div class="input-group mb-3">
                <input v-model="searchUsername" type="text" class="form-control" placeholder="Type username...">
                <div class="input-group-append">
                    <button class="input-group-text text-primary send-button" type="button" @click="searchUsers">
                        <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#search"/></svg>
                    </button>
                </div>
            </div>
            <hr>
			<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
            <div v-if="searched" class="card">
				<div class="card-title mt-2 mb-0">
					<RouterLink :to="$pathToProfile(searched.username)" class="user-link">
						<h4>{{ searched.username }}</h4>
					</RouterLink>	
				</div>
				<hr class="mt-0 mb-0" style="margin: 30px;">
				<div class="card-body text-center">
					<div class="d-flex justify-content-between">
						<div>
							<p class="mb-0">{{ searched.followersCount }}</p>
							<p class="mb-0 text-muted">Followers</p>
						</div>
						<div>
							<p class="mb-0">{{ searched.followingCount }}</p>
							<p class="mb-0 text-muted">Following</p>
						</div>
						<div>
							<p class="mb-0">{{ searched.photoCount }}</p>
							<p class="mb-0 text-muted">Posts</p>
						</div>
					</div>
    			</div>
			</div>
        </div>
    </div>
</template>

<style>

</style>
