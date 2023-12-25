<script setup>
import Modal from '../components/PhotoModal.vue'
</script>
<script>
export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			stream: null,
			selectedPost: null,
		}
	},
	components: {
		Modal,
	},
	methods: {
		async getUserStream() {
			this.loading = true;
			this.errormsg = null;
			var path = `/users/${this.$getCurrentId()}/stream`;
			try {
				let response = await this.$axios.get(path);
				if (response.status == 200) {
					this.stream = response.data;

					this.stream.forEach( e => {
						// stream[index].image = 'data:image/*;base64,' + stream[index].image;
						e.image = `data:image/*;base64,${e.image}`;
						e.timestamp = this.$timestampToDate(e.timestamp);
					});
				}
			} catch (e) {
				this.errormsg = e.toString();				
			}
			this.loading = false;
		},

		toggleModal(post) {
			this.selectedPost = post;
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
				<img :src="post.image" @click="toggleModal(post)">
				<p v-text="post.caption"></p>
				<p>Commenti: {{ post.commentCount }} Likes: {{ post.likeCount }} {{ post.timestamp }}</p>
			</div>
		</div>
		<div v-else>
			<p> Sorry nothing to show here </p>	
		</div>	
	</div>
	<Modal v-if="selectedPost" @close="toggleModal(null)" :photo="selectedPost"> </Modal>

</template>

<style>
</style>
