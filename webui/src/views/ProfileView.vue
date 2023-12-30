<script setup>
import Modal from '../components/PhotoModal.vue'
import UsersListModal from '../components/UsersListModal.vue'
</script>
<script>
export default {
	data: function() {
		return {
			loading: false,
			errormsg: null,
			user: null,
			username: "",
			name: "",
			surname: "",
			photos: [],
			moreContent: false,
			streamPage: 0,
			selectedPost: null, 
			followPath: null,
			bans: [],
		}
	},
	components: {
		Modal,
		UsersListModal,
	},
	watch:{
		async '$route.params.username'(newUsername) {
			if (!newUsername) return;  // If there is no username, don't do anything 
			this.photos = [];
			this.getProfile();
		},
	},	
	methods: {
		scroll () {
			// Check when user reach the end of the page
			window.onscroll = () => {
				let bottomOfWindow = window.scrollY + window.innerHeight >= document.documentElement.scrollHeight;
				if (bottomOfWindow) {
					if (this.moreContent) {
						this.streamPage++;
						this.getPhotos();
					}
				}
			}
		},

		async getProfile(){
			// Load user profile
			this.loading = true;
			this.errormsg = null;
			var path = `/users?username=${this.$route.params.username}`;
			try {
				let response = await this.$axios.get(path);
				if (response.status == 200) {
					this.user = response.data;
					this.username = this.user.username;
					this.name = this.user.name;
					this.surname = this.user.surname;	
				}
			} catch (e) {
				this.errormsg = e.toString();
				return;				
			}
			this.loading = false;

			this.getPhotos();

			if (this.$getCurrentId() == this.user.id) this.getBanned();
		},

		async getPhotos(){
			// Load user posts
			this.loading = true;
			this.errormsg = null;
			var path = `/users/${this.user.id}/photos?page=${this.streamPage}`;
			try {
				let response = await this.$axios.get(path);
				if (response.status == 200) {
					response.data.forEach( e => {
						// stream[index].image = 'data:image/*;base64,' + stream[index].image;
						e.image = `data:image/*;base64,${e.image}`;
						e.timestamp = this.$timestampToDate(e.timestamp);
					});
					if (response.data !== null) {
						this.moreContent = response.data.length === 10;
						this.photos = this.photos.length > 0 ? this.photos.concat(response.data) : response.data;
					}
					else 
						this.moreContent = false;	
				}
			} catch (e) {
				this.errormsg = e.toString();
				return;				
			}
			this.loading = false;
		},

		async getBanned() {
			this.loading = true;
            this.errormsg = null;
            try {
                var path = `/users/${this.user.id}/banned`;
                let response = await this.$axios.get(path);

                if (response.status === 200) {
                    this.bans = response.data;
					console.log(this.bans);
                }
            } catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
		},

		async editProfile(){
			this.loading = true;
			this.errormsg = null;
			var path = `/users/${this.$getCurrentId()}`;
			try {
				let response = await this.$axios.put(path, {
					username: this.username,
					name: this.name,
					surname: this.surname,
				});
				if (response.status == 200) {
					if (this.user.username !== this.username) {
						this.$router.replace(`/user/${this.username}`);
						localStorage.username = this.username;	
					}
					this.user = response.data;	
				}
			} catch (e) {
				this.errormsg = e.toString();	
				this.username = this.user.username;
				this.name = this.user.name;
				this.surname = this.user.surname;	
			}
			this.loading = false;
		},

		toggleModal(post) {
			this.selectedPost = post;
		},

		// 1 for Followers, 2 for Following, 3 for Banned
		toggleFollowModal(index) {
			this.followPath = index
		}
	},
	mounted() {
		this.getProfile();
		this.streamPage = 0;
		this.scroll();
	}
}
</script>

<template>
	<div class="col-md-12 ms-sm-auto col-lg-12 px-md-4 mt-5">
		<div class="row" v-if="user">
			<div class="col-md-3">
			<!-- User Information -->
				<div class="card">
					<div class="card-body">
						<h5 class="card-title"> {{ user.username }} </h5>
						<p v-if="user.name || user.surname" class="card-text mb-1"> {{ user.name }} {{ user.surname }}</p>
						<small v-else class="fst-italic">No name and surname</small>
						<div v-if="this.$getCurrentId() == user.id" class="mt-2">
							<button class="profile-buttons profile-buttons-primary" type="button" data-bs-toggle="collapse" data-bs-target="#divEdit" aria-expanded="false" aria-controls="divEdit">
								Edit
							</button>
						</div>
						<div class="collapse" id="divEdit">
							<div class="mt-2 mb-3 row">
								<label for="inputUsername" class="col-sm-4 col-form-label">Username</label>
								<div class="col-sm-8">
									<input type="text" class="input-profile form-control" id="inputUsername" v-model="username">
								</div>
							</div>
							<div class="mt-2 mb-3 row">
								<label for="inputName" class="col-sm-4 col-form-label">Name</label>
								<div class="col-sm-8">
									<input type="text" class="input-profile form-control" id="inputName" v-model="name">
								</div>
							</div>
							<div class="mb-3 row">
								<label for="inputSurname" class="col-sm-4 col-form-label">Surname</label>
								<div class="col-sm-8">
									<input type="text" class="input-profile form-control" id="inputSurname" v-model="surname">
								</div>
							</div>
							<div class="d-flex justify-content-between">
								<button class="profile-buttons profile-buttons-danger" type="button" data-bs-toggle="collapse" data-bs-target="#divEdit" aria-expanded="false" aria-controls="divEdit">
									Cancel
								</button>
								<button class="profile-buttons profile-buttons-primary" type="button" @click="editProfile">Save</button>
							</div>
						</div>
					</div>
					<ul class="list-group list-group-flush">
						<li class="list-group-item d-flex align-items-center" @click="toggleFollowModal(1)">
							<label role="button">Followers:&nbsp;</label> 
							<strong role="button" class="fs-6">{{ user.followersCount }}</strong> 
						</li>
						<li class="list-group-item d-flex align-items-center" @click="toggleFollowModal(2)">
							<label role="button">Following:&nbsp;</label> 
							<strong role="button" class="fs-6"> {{ user.followingCount }}</strong> 
						</li>
						<li v-if="this.$getCurrentId() == user.id" class="list-group-item d-flex align-items-center" @click="toggleFollowModal(3)">
							<label role="button">Bans:&nbsp;</label> 
							<strong role="button" class="fs-6"> {{ bans.length }}</strong> 
						</li>
						<li class="list-group-item d-flex align-items-center">
							<label>Posts:&nbsp;</label> 
							<strong class="fs-6"> {{ user.photoCount }}</strong> 
						</li>
					</ul>
				</div>
			</div>
			<div class="col-md-9">
			<!-- User Images -->
				<div class="card-columns">
					<!-- <div class="card" v-for="(image, index) in images" :key="index"> -->
					<div v-if="photos.length > 0">
						<div class="card" v-for="(photo, index) in photos" :key="index">
							<div class="card-body">
								<h5 class="card-title"> {{ user.username }}</h5>
								<img :src="photo.image" class="card-img-top img-fluid" alt="" @click="toggleModal(photo)">	
								<p class="card-text" v-if="photo.caption" v-text="photo.caption"></p>		
								<p class="card-text">Likes: {{ photo.likeCount }}, Comments: {{ photo.commentCount }}</p>
							</div>
						</div>
					</div>
					<div v-else>
						<div class="card align-items-center">
							<div class="card-body">
								<h5 class="card-title mb-0"> Nothing to show here </h5>
							</div>
						</div>
					</div>
				</div>
			</div>
		</div>
	</div>
	<Modal v-if="selectedPost" @close="toggleModal(null)" :photo="selectedPost"> </Modal>
	<UsersListModal v-if="followPath" @close="toggleFollowModal(null)" :path="followPath" :name="user.username" :id="user.id" :bans="bans"></UsersListModal>
</template>

<style>
.input-profile {
	font-size: 0.9rem;
}

</style>