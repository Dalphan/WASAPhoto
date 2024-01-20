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
			currentUsername: null,
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
		this.currentUsername = this.$getCurrentUsername();
		this.streamPage = 0;
		this.getUserStream();
		this.scroll();
	}
}
</script>

<template>
	<div class="stream-page">
		<div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
		<h2>{{ currentUsername }}'s Stream</h2>
	  </div>
  
	  <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
	  <div v-else>
		<div v-if="stream.length > 0">
		  <div class="post-container" v-for="post in stream" :key="post.id">
			<div class="user-info d-flex justify-content-between align-items-center">
				<div>
					<RouterLink :to="$pathToProfile(post.username)" class="user-link">
						<span class="post-name">{{ post.username }}</span>
					</RouterLink>
				</div>
				<div class="post-stats">
					{{ post.timestamp }}
					<text class="post-stats-heart"> {{ post.commentCount }} <svg class="feather" style="width: 23px; height: 23px;"><use href="/feather-sprite-v4.29.0.svg#message-circle"/></svg></text> 
					<text>{{ post.likeCount }} <svg class="feather" style="width: 23px; height: 23px;"><use href="/feather-sprite-v4.29.0.svg#heart"/></svg> </text>
				</div>
			</div>
			<div class="post-content">
			  <img role="button" :src="post.image" @click="toggleModal(post)" class="post-image"/>
			  <p v-if="post.caption" class="post-caption">{{ post.caption }}</p>
			</div>
		  </div>
		</div>
		<div v-else>
		  <p class="empty-message">There is nothing to display here. Begin by following someone!</p>
		</div>
	  </div>
	</div>
	<Modal v-if="selectedPost" @close="toggleModal(null)" :photo="selectedPost"></Modal>
</template>
  
<style scoped>
.stream-page {
	max-width: 70%;
	margin: 0 auto;
}

.post-container {
	display: flex;
	flex-direction: column;
	margin-bottom: 20px;
	border: 1px solid #ddd;
	border-radius: 8px;
	overflow: hidden;
}

.user-info {
	padding: 15px;
	background-color: #f8f8f8;
}

.post-name {
	font-size: 1.3rem;
	font-weight: bold;
}

.post-content {
	max-width: 100%;
	padding: 15px;
}

.post-image {
	width: 100%;
	cursor: pointer;
}

.post-caption {
	margin-top: 8px;
	font-size: 1rem;
	margin-bottom: 0px;
}

.post-stats {
	font-size: 0.85rem;
	color: #888;
	text-align:center;
}

.post-stats-heart{
	border-left: 1px solid #888; 
	padding-left: 5px; 
	padding-right: 5px;
}

.empty-message {
	font-size: 0.9rem;
	color: #888;
	text-align: center;
}
</style>
  