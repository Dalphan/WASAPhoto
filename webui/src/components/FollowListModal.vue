<script>
export default {
    name: 'FollowModal',
    props: {
        path: 0,
    },
    data: function() {
        return {
            errormsg: null,
            loading: false,
            follows: [],
            title: "",
        }
    },
    watch:{
        path: {
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
                    }
                }
            }
        }
    },	
    methods: {
        close() {
            this.$emit('close');
        },

        dontclose() {

        },
        
        async getFollowers(){

        },

        async getFollowing(){

        },
    },
}
</script>

<template>
    <div class="modal" @click="close">
        <div class="modal-inner" @click.stop="">        
            <div class="container mt-2">
                <div class="text-center mb-4">
                    <h1> {{ title }} </h1>
                </div>

                <div class="input-group mb-3">
                    <input type="text" class="form-control" placeholder="Search..." aria-label="Search" aria-describedby="searchButton">
                    <div class="input-group-append">
                    <button class="btn btn-outline-secondary" type="button" id="searchButton">
                        <i class="fas fa-search"></i>
                    </button>
                    </div>
                </div>

                <ul class="list-group">
                    <li class="list-group-item d-flex justify-content-between align-items-center">
                    John Doe
                    <button class="btn btn-primary">Action</button>
                    </li>
                    <li class="list-group-item d-flex justify-content-between align-items-center">
                    Jane Smith
                    <button class="btn btn-primary">Action</button>
                    </li>
                    <!-- Add more list items as needed -->
                </ul>
            </div>
        </div>
    </div>
</template>