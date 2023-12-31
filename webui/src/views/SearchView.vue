<script>
export default {
	data: function() {
		return {
			errormsg: "",
			loading: false,
		}
	},
	methods: {
		// async refresh() {
		// 	this.loading = true;
		// 	this.errormsg = null;
		// 	try {
		// 		let response = await this.$axios.post("/session", {
		// 			name: 'Maria',
		// 		});
		// 		this.some_data = response.data;
		// 	} catch (e) {
		// 		this.errormsg = e.toString();
		// 	}
		// 	this.loading = false;
		// },

	},
	mounted() {

	}
}
</script>

<template>
    <div class="form-body">
        <div class="form-inner">
            <div class="mb-3">
                <h3>Search users</h3>
            </div>
            <div class="input-group mb-3">
                <input v-model="searchedUser" type="text" class="form-control" placeholder="Type username...">
                <div class="input-group-append">
                    <button class="input-group-text text-primary send-button" type="button" @click="filterUsersList">
                        <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#search"/></svg>
                    </button>
                </div>
            </div>
            <hr>
            
        </div>
    </div>
</template>

<style>

</style>
