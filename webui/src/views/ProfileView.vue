<script setup>

</script>
<script>
export default {
	data: function() {
		return {
			
		}
	},
	watch:{
        async '$route.params.username'(newUsername) {
            if (!newUsername) return;  // If there is no new username, don't do anything 
			this.getProfile();
		},
	},	
	methods: {
        async getProfile(){
            alert(this.$route.params.username);
        }
	},
	mounted() {
		this.getProfile();
	}
}
</script>

<template>

</template>