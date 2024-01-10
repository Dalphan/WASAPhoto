<script setup>
</script>

<script>
export default {
    name: 'UsersListModal',
    props: ["list_mode", "username", "user_id", "bans"],
    data: function() {
        return {
            errormsg: null,
            loading: false,
            usersList: [],
            filteredList: [],
            searchedUser: null,
            title: "",
        }
    },
    watch:{
        list_mode: {
            immediate: true,
            handler(newVal, oldVal){
                if (newVal != oldVal) {
                    // Aggiorna lista di followers
                    if (newVal === 1) {
                        this.getFollowers();
                        this.title = "Followers";
                    } else if (newVal === 2) {
                        this.getFollowing();
                        this.title = "Following";
                    } else if (newVal === 3) {
                        // this.getBanned();
                        this.usersList = this.bans;
                        this.filteredList = this.usersList;
                        this.title = "Bans";
                    }
                }
            }
        }
    },	
    methods: {
        close() {
            this.$emit('close');
        },
        
        async getFollowers(){
            this.loading = true;
            this.errormsg = null;
            try {
                var path = `/users/${this.user_id}/followers`;
                let response = await this.$axios.get(path);

                if (response.status === 200) {
                    this.usersList = response.data;
                    this.filteredList = this.usersList;
                }
            } catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
        },

        async getFollowing(){
            this.loading = true;
            this.errormsg = null;
            try {
                var path = `/users/${this.user_id}/following`;
                let response = await this.$axios.get(path);

                if (response.status === 200) {
                    this.usersList = response.data;
                    this.filteredList = this.usersList;
                }
            } catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
        },

        filterUsersList(){
            if (this.searchedUser === "") {
                this.filteredList = this.usersList;
                return;
            }
            this.filteredList = this.usersList.filter(user => user.username.toLowerCase().startsWith(this.searchedUser.toLowerCase()));
        }
    },
}
</script>

<template>
    <div class="modal" @click="close">
        <div class="modal-inner" @click.stop="" style="max-height: 80vh;">
            <div class="container mt-2">
                <div class="text-center mb-4">
                    <h3>{{ username }}'s {{ title }}</h3>
                </div>

                <div class="input-group mb-3">
                    <input v-model="searchedUser" type="text" class="form-control" placeholder="Search...">
                    <div class="input-group-append">
                        <button class="input-group-text text-primary send-button" type="button" @click="filterUsersList">
                            <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#search"/></svg>
                        </button>
                    </div>
                </div>
                <div>
                    <ul class="list-group" style="max-height: 60vh; overflow-y: auto;"> <!-- Set max-height and enable vertical scrolling for the list -->
                        <li v-for="u in this.filteredList" class="list-group-item d-flex justify-content-between align-items-center" :key="u.id">
                            <RouterLink @click="close" :to="this.$pathToProfile(u.username)" class="user-link">
                                <strong>{{ u.username }}</strong><br>
                            </RouterLink>
                            <!--  
                            <button v-if="list_mode === 2 && this.$getCurrentId() == this.user_id" class="profile-buttons profile-buttons-danger" title="Remove following">
                                <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#x-circle"/></svg>
                            </button>
                            <button v-if="list_mode === 3 && this.$getCurrentId() == this.user_id" class="profile-buttons profile-buttons-success" title="Remove ban">
                                <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#user-plus"/></svg>
                            </button>-->

                            <!--<button v-else class="profile-buttons profile-buttons-primary">Follow</button>-->
                        </li>
                    </ul>
                </div>
            </div>
        </div>
    </div>
</template>