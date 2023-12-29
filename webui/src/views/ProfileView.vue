<script setup>

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
		}
	},
	watch:{
		async '$route.params.username'(newUsername) {
			if (!newUsername) return;  // If there is no username, don't do anything 
			this.getProfile();
		},
	},	
	methods: {
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

				// Load user posts
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
	},
	mounted() {
		this.getProfile();
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
          <p v-if="user.name || user.surname" class="card-text"> {{ user.name }} {{ user.surname }}</p>
		<small v-else class="fst-italic">No name and surname</small>
		<div v-if="this.$getCurrentId() == user.id" class="mt-2">
			<button class="profile-buttons profile-buttons-primary" type="button" data-bs-toggle="collapse" data-bs-target="#divEdit" aria-expanded="false" aria-controls="divEdit">Modifica</button>
		</div>
		<div class="collapse" id="divEdit">
			<div class="mt-2 mb-3 row">
				<label for="inputUsername" class="col-sm-4 col-form-label">Username</label>
				<div class="col-sm-8">
					<input type="text" class="input-profile form-control" id="inputUsername" v-model="username">
				</div>
			</div>
			<div class="mt-2 mb-3 row">
				<label for="inputName" class="col-sm-4 col-form-label">Nome</label>
				<div class="col-sm-8">
					<input type="text" class="input-profile form-control" id="inputName" v-model="name">
				</div>
			</div>
			<div class="mb-3 row">
				<label for="inputSurname" class="col-sm-4 col-form-label">Cognome</label>
				<div class="col-sm-8">
					<input type="text" class="input-profile form-control" id="inputSurname" v-model="surname">
				</div>
			</div>
			<div class="d-flex justify-content-between">
				<button class="profile-buttons profile-buttons-danger" type="button" data-bs-toggle="collapse" data-bs-target="#divEdit" aria-expanded="false" aria-controls="divEdit">Cancella</button>
				<button class="profile-buttons profile-buttons-primary" type="button" @click="editProfile">Salva</button>
			</div>
		</div>
	</div>
	<ul class="list-group list-group-flush">
		<li class="list-group-item">Followers: 100 </li>
		<li class="list-group-item">Following:  102 </li>
		<li class="list-group-item">Images: 40 </li>
	</ul>
</div>
    </div>
    <div class="col-md-9">
      <!-- User Images -->
      <div class="card-columns">
        <!-- <div class="card" v-for="(image, index) in images" :key="index"> -->
		<div class="card">
			<div class="card-body">
				<h5 class="card-title"> {{ user.username }}</h5>
				<img src="https://mdbcdn.b-cdn.net/img/Photos/Lightbox/Original/img%20(107).webp" class="card-img-top img-fluid" alt="">			
				<p class="card-text">Likes: 10, Comments: 12</p>
			</div>
        </div>
      </div>
    </div>
  </div>
</div>
</template>

<style>
.input-profile {
	font-size: 0.9rem;
}

.profile-buttons-primary {
  background-color: #3b82f6;
  color: #f3f7fe;
}

.profile-buttons-danger {
  background-color: #fc4646;
  color: #f3f7fe;
}

.profile-buttons {
  font-size: 0.9rem;
  border: none;
  border-radius: 20px;
  width: auto;
  height: auto;
  padding: 0.3rem 0.6rem;
  transition: .3s;
}

.profile-buttons-danger:hover {
  background-color: #f3f7fe;
  box-shadow: 0 0 0 5px #f63b3b5f;
  color: #fc4646;
}

.profile-buttons-primary:hover {
  background-color: #f3f7fe;
  box-shadow: 0 0 0 5px #3b83f65f;
  color: #3b82f6;
}

</style>