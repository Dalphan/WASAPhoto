<script>
export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			stream: null,
		}
	},
	methods: {
		async refresh() {
			// this.loading = true;
			// this.errormsg = null;
			// try {
			// 	let response = await this.$axios.post("/session", {
			// 		name: 'Maria',
			// 	});
			// 	this.some_data = response.data;
			// } catch (e) {
			// 	this.errormsg = e.toString();
			// }
			// this.loading = false;
		},
		async getUserStream() {
			this.loading = true;
			this.errormsg = null;
			var path = "/users/" + this.$getCurrentId() + "/stream";
			try {
				let response = await this.$axios.get(path);
				console.log(response);
				if (response.status == 200) {
					this.stream = response.data
					// for (var post in this.stream) {
					// 	post.image = 'data:image/png;base64,' + btoa(post.image);
					// }
					// this.stream.forEach(function(e, index, stream){
					// 	stream[index].image = 'data:image/jpeg;base64,' + stream[index].image
					// });
				}
			} catch (e) {
				this.errormsg = e.toString();				
			}
			this.loading = false;
		}
	},
	mounted() {
		this.getUserStream()
	}
}
</script>

<template>
	<div>
		<div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">Home page</h1>
			<div class="btn-toolbar mb-2 mb-md-0">
				<div class="btn-group me-2">
					<button type="button" class="btn btn-sm btn-outline-secondary" @click="refresh">
						Refresh
					</button>
					<button type="button" class="btn btn-sm btn-outline-secondary" @click="exportList">
						Export
					</button>
				</div>
				<div class="btn-group me-2">
					<button type="button" class="btn btn-sm btn-outline-primary" @click="newItem">
						New
					</button>
				</div>
			</div>		
		</div>

		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
		<div>
			<div class="row" v-for="post in stream">
				<p v-text="post.user"></p>
				
						<img :src="post.image">
				
			</div>
		</div>
	</div>

</template>

<style>
</style>
