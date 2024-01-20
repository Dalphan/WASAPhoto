<script setup>
import Modal from '../components/PhotoModal.vue'
import UsersListModal from '../components/UsersListModal.vue'
import ErrorMsg from '../components/ErrorMsg.vue'
import NewPostModal from '../components/NewPostModal.vue'
</script>
<script>
export default {
	data: function() {
		return {
			loading: false,
			errormsg: null,
			currentId: null,
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
			createPost: false,
		}
	},
	components: {
		Modal,
		UsersListModal,
		ErrorMsg,
		NewPostModal,
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
				this.user = null;
				this.errormsg = e.response.data !== null ? e.response.data : e.toString();
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
					if (response.data !== null) {
						response.data.forEach( e => {
							e.image = `data:image/*;base64,${e.image}`;
							e.timestamp = this.$timestampToDate(e.timestamp);
						});

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

		async deletePost(index){
			// Load user posts
			this.loading = true;
			this.errormsg = null;
			try {
				var path = `/photos/${this.photos[index].id}`;
				let response = await this.$axios.delete(path);

				if (response.status == 200) {
					this.photos.splice(index, 1);
					this.user.photoCount--;
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

                if (response.status == 200) {
					if (response.data !== null) this.bans = response.data;
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
						this.user.username = this.username;
					}
					// this.user = response.data;	
					this.user.name = this.name;
					this.user.surname = this.surname;
				}
			} catch (e) {
				this.errormsg = e.toString();	
				this.username = this.user.username;
				this.name = this.user.name;
				this.surname = this.user.surname;	
			}
			this.loading = false;
		},

		async FollowUser() {
			this.loading = true;
            this.errormsg = null;
			let id = this.$getCurrentId();
            try {
                var path = `/users/${id}/following/${this.user.id}`;
                let response = await this.$axios.put(path);

                if (response.status == 200 || response.status == 201) {
					if (id === this.user.id) {
						// The user is in their profile page
						// NOT IMPLEMENTED
					} else {
						// The user is not in their profile page
						this.user.followed = true;
						this.user.followersCount++;
					}
                }
            } catch (e) {
				this.errormsg = e.toString();
            }
            this.loading = false;
		},

		async UnfollowUser() {
			this.loading = true;
            this.errormsg = null;
			let id = this.$getCurrentId();
            try {
                var path = `/users/${id}/following/${this.user.id}`;
                let response = await this.$axios.delete(path);

                if (response.status == 200 || response.status == 201) {
					if (id === this.user.id) {
						// The user is in their profile page
						// NOT IMPLEMENTED
					} else {
						// The user is not in their profile page
						this.user.followed = false;
						this.user.followersCount--;
					}
                }
            } catch (e) {
				this.errormsg = e.toString();
            }
            this.loading = false;
		},

		toggleFollow() {
			if (this.user.followed) this.UnfollowUser();
			else this.FollowUser();
		},

		async BanUser() {
			this.loading = true;
            this.errormsg = null;
			let id = this.$getCurrentId();
            try {
                var path = `/users/${id}/banned/${this.user.id}`;
                let response = await this.$axios.put(path);

                if (response.status == 200 || response.status == 201) {
					if (id === this.user.id) {
						// The user is in their profile page
						// NOT IMPLEMENTED
					} else {
						// The user is not in their profile page
						this.user.banned = true;

						// The ban removes the follow too
						this.user.followed = false;
						this.user.followersCount--;
						this.user.followingCount = response.data.followingCount;
					}
                }
            } catch (e) {
				this.errormsg = e.toString();
            }
            this.loading = false;
		},

		async UnbanUser() {
			this.loading = true;
            this.errormsg = null;
			let id = this.$getCurrentId();
            try {
                var path = `/users/${id}/banned/${this.user.id}`;
                let response = await this.$axios.delete(path);

                if (response.status == 200 || response.status == 201) {
					if (id === this.user.id) {
						// The user is in their profile page
						// NOT IMPLEMENTED
					} else {
						// The user is not in their profile page
						this.user.banned = false;
					}
                }
            } catch (e) {
				this.errormsg = e.toString();
            }
            this.loading = false;
		},

		toggleBan() {
			if (this.user.banned) this.UnbanUser();
			else this.BanUser();
		},

		toggleModal(post) {
			this.selectedPost = post;
		},

		// 1 for Followers, 2 for Following, 3 for Banned
		toggleFollowModal(index) {
			this.followPath = index
		},

		toggleCreateModal(reload){
			this.createPost = !this.createPost;
			if (reload) {
				this.streamPage = 0;
				this.photos = [];
				this.user.photoCount++;
				this.getPhotos();
			}
		}
	},
	mounted() {
		this.currentId = this.$getCurrentId();
		this.getProfile();
		this.streamPage = 0;
		this.scroll();
	}
}
</script>

<template>
	<div class="col-md-12 ms-sm-auto col-lg-12 px-md-4 mt-5">
		<div class="row" v-if="errormsg">
			<ErrorMsg :msg="errormsg"></ErrorMsg>
		</div>
		<div class="row" v-if="user">
			<div class="col-md-3">
			<!-- User Information -->
				<div class="card">
					<div class="card-body">
						<h4 class="card-title"> {{ user.username }} </h4>
						<p v-if="user.name || user.surname" class="card-text mb-1"> {{ user.name }} {{ user.surname }}</p>
						<small v-else class="fst-italic">No name and surname</small>
						<div v-if="currentId == user.id" class="mt-2">
							<button class="profile-buttons profile-buttons-primary" type="button" data-bs-toggle="collapse" data-bs-target="#divEdit" aria-expanded="false" aria-controls="divEdit">
								Edit
							</button>
						</div>
						<div v-else class="mt-2">
							<div class="d-flex justify-content-between">
								<div @click="toggleFollow">
									<button v-if="user.followed" class="profile-buttons profile-buttons-primary" type="button">
										Unfollow
									</button>
									<button v-else class="profile-buttons profile-buttons-secondary" type="button">
										Follow
									</button>
								</div>
								<div @click="toggleBan">
									<button v-if="user.banned" class="profile-buttons profile-buttons-success" type="button">
										Unban
									</button>
									<button v-else class="profile-buttons profile-buttons-danger" type="button">
										Ban
									</button>
								</div>
							</div>
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
						<li v-if="currentId == user.id" class="list-group-item d-flex align-items-center" @click="toggleFollowModal(3)">
							<label role="button">Bans:&nbsp;</label> 
							<strong role="button" class="fs-6"> {{ bans.length }}</strong> 
						</li>
						<li class="list-group-item d-flex align-items-center justify-content-between">
							<div>
								<label>Posts:&nbsp;</label> 
								<strong class="fs-6">{{ user.photoCount }}</strong> 
							</div>
							<button v-if="currentId == user.id" class="profile-buttons profile-buttons-success" @click="toggleCreateModal(false)">
								<svg style="margin-left: 2px;" class="feather"><use href="/feather-sprite-v4.29.0.svg#plus"/></svg>
								New post
							</button>
						</li>
					</ul>
				</div>
			</div>
			<div class="col-md-9">
			<!-- User Images -->
				<div class="card-columns">
					<!-- <div class="card" v-for="(image, index) in images" :key="index"> -->
					<div v-if="photos.length > 0">
						<div class="card post" v-for="(photo, index) in photos" :key="index">
							<div class="user-info d-flex justify-content-between align-items-center">
								<span class="post-name"> {{ user.username }}</span>
								<div class="d-flex">
									<div class="post-stats">
										{{ photo.timestamp }}
										<text class="post-stats-heart"> {{ photo.commentCount }} <svg class="feather" style="width: 23px; height: 23px;"><use href="/feather-sprite-v4.29.0.svg#message-circle"/></svg></text> 
										<text>{{ photo.likeCount }} <svg class="feather" style="width: 23px; height: 23px;"><use href="/feather-sprite-v4.29.0.svg#heart"/></svg> </text>
									</div>
									<div v-if="currentId == user.id" title="Delete Post" @click="deletePost(index)">
										<svg role="button" class="feather text-danger" style="width: 24px; height: 24px;">
											<use href="/feather-sprite-v4.29.0.svg#trash-2"/>
										</svg>
									</div>
								</div>
							</div>
							<div class="card-body">
								<img :src="photo.image" class="card-img-top img-fluid" alt="" @click="toggleModal(photo)">	
								<p class="post-caption" v-if="photo.caption" v-text="photo.caption"></p>
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
	<UsersListModal v-if="followPath" @close="toggleFollowModal(null)" :list_mode="followPath" :username="user.username" :user_id="user.id" :bans="bans"></UsersListModal>
	<NewPostModal v-if="createPost" @close="toggleCreateModal(false)" @closeReload="toggleCreateModal(true)" :username="user.username"></NewPostModal>
</template>

<style>
.input-profile {
	font-size: 0.9rem;
}

.user-info {
	padding: 15px;
	background-color: #f8f8f8;
}

.post-name {
	font-size: 1.3rem;
	font-weight: bold;
}

.post {
	margin-bottom: 20px;
}

.post-caption {
	margin-top: 8px;
	font-size: 1rem;
	margin-bottom: 0px;
}
  
.post-stats {
	font-size: 0.85rem;
	color: #888;
	margin-right: 15px;

}

.post-stats-heart{
	border-left: 1px solid #888; 
	padding-left: 5px; 
	padding-right: 5px;
}

</style>