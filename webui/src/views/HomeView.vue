<script setup>
import { RouterLink } from 'vue-router';
import Modal from '../components/PhotoModal.vue'
</script>
<script>
export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			stream: [],
			selectedPost: null,
			streamPage: 0,
			moreContent: false,
		}
	},
	components: {
		Modal,
	},
	methods: {
		scroll () {
			// Check when user reach the end of the page
			window.onscroll = () => {
				let bottomOfWindow = window.scrollY + window.innerHeight >= document.documentElement.scrollHeight;
				if (bottomOfWindow) {
					if (this.moreContent) {
						this.streamPage++;
						this.getUserStream();
					}
				}
			}
		},

		async getUserStream() {
			this.loading = true;
			this.errormsg = null;
			var path = `/users/${this.$getCurrentId()}/stream?page=${this.streamPage}`;
			try {
				let response = await this.$axios.get(path);
				if (response.status == 200) {
					if (response.data !== null) {
						// adjust image and timestamp field			
						response.data.forEach( e => {
							e.image = `data:image/*;base64,${e.image}`;
							e.timestamp = this.$timestampToDate(e.timestamp);
						});
						this.moreContent = response.data.length === 10;
						this.stream = this.stream.length > 0 ? this.stream.concat(response.data) : response.data;
					}
					else this.moreContent = false;			
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
		this.streamPage = 0;
		this.getUserStream();
		this.scroll();
	}
}
</script>

<template>
	<div>
		<div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2"> {{ this.$getCurrentUsername() }}'s stream </h1>
			<!-- <div class="btn-toolbar mb-2 mb-md-0">
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
			</div>		 -->
		</div>

		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
		<div v-if="stream.length > 0">
			<div class="row" v-for="post in stream" :key="post.id">
				<RouterLink :to="this.$pathToProfile(post.username)" class="user-link">
					<p v-text="post.username"></p>
				</RouterLink>
				<img role="button" :src="post.image" @click="toggleModal(post)">
				<p v-if="post.caption" v-text="post.caption"></p>
				<p>{{ post.commentCount }} comments {{ post.likeCount }} likes {{ post.timestamp }}</p>
			</div>
		</div>
		<div v-else>
			<p> There is nothing to display here. Begin by following someone! </p>	
		</div>	
	</div>
	<Modal v-if="selectedPost" @close="toggleModal(null)" :photo="selectedPost"> </Modal>

</template>

<style>
</style>
