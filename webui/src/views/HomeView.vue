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
		async getUserStream() {
			this.loading = true;
			this.errormsg = null;
			var path = `/users/${this.$getCurrentId()}/stream`;
			try {
				let response = await this.$axios.get(path);
				console.log(response);
				if (response.status == 200) {
					this.stream = response.data
					// this.stream.forEach(function(e, index, stream){
					// 	stream[index].image = 'data:image/*;base64,' + stream[index].image;
					// });
				}
			} catch (e) {
				this.errormsg = e.toString();				
			}
			this.loading = false;
		},

		timestampToDate(timestamp){
			const today = new Date(Date.now())

			console.log(new Date(timestamp));
			console.log(today)

			var difference = new Date() - new Date(timestamp);	
			console.log(difference);
			var seconds = Math.floor(difference / 1000);
			var minutes = Math.floor(seconds / 60);
			var hours = Math.floor(minutes / 60);
			var days = Math.floor(hours / 24);
			
			if (days > 0) {
				return `${days} day${days > 1 ? 's' : ''} ago`;
			} else if (hours > 0) {
				return `${hours} hour${hours > 1 ? 's' : ''} ago`;
			} else if (minutes > 0) {
				return `${minutes} minute${minutes > 1 ? 's' : ''} ago`;
			} else {
				return `${seconds} seconds ago`;
			}
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
		<div v-if="stream">
			<div class="row" v-for="(post, index) in stream" :key="post.id">
				<p v-text="post.username"></p>
				<img :src="`data:image/*;base64,${post.image}`">
				<p v-text="post.caption"></p>
				<p>Commenti: {{ post.commentCount }} Likes: {{ post.likeCount }} {{ timestampToDate(post.timestamp) }}</p>
			</div>
		</div>
		<div v-else>
			<p> Sorry nothing to show here </p>	
		</div>	
	</div>

</template>

<style>
</style>
