<script setup>
</script>
<script>
export default {
    name: 'NewPostModal',
    props: {
        username: "",
    },
	data: function() {
		return {
			errormsg: "",
			loading: false,
            selectedImage: null,
            caption: "",
            file: null,
		}
	},
	methods: {
        close() {
            this.$emit('close');
        },
        handleImageChange(event) {
            this.file = event.target.files[0];
            if (this.file) {
                const reader = new FileReader();
                reader.onload = (e) => {
                    this.selectedImage = e.target.result;
                };
                reader.readAsDataURL(this.file);
            }
        },

        async postPhoto() {
			this.loading = true;
			this.errormsg = null;
			try {
                const formData = new FormData();
                formData.append('image', this.file);
                var path = `/photos?caption=${this.caption}`;
				let response = await this.$axios.post(path, formData, {
					Headers: {
                        'Content-Type':'image/*',
                    }
				});

                if (response.status == 200 || response.status == 201) {
                    this.$emit('closeReload');
                }
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		},
	},
	mounted() {

	}
}
</script>

<template>
    <div class="modal" @click="close">
        <div class="modal-inner" @click.stop="" style="max-width: 80vw;">
            <div class="mt-2 d-flex">
                <div class="left-section col-lg-9">
                    <div v-if="selectedImage" style="max-width: 40vw; ">
                        <img :src="selectedImage" alt="Selected Image" class="img-fluid">
                    </div>
                    <div v-else class="image-placeholder">
                        <div class="placeholder-content">
                            <p>Choose Image</p>
                        </div>
                    </div>
                </div>

                <div class="right-section col-lg-3 d-flex flex-column" style="max-width: 20vw;">
                    <div class="d-flex justify-content-between">
                        <div>
                            <h4>{{ username }}</h4>
                        </div>
                        <svg class="feather" role="button" @click="close"><use href="/feather-sprite-v4.29.0.svg#x"/></svg>
                    </div>
                    <div class="flex-grow-1">
                        <div>
                            <label for="caption">Caption:</label>
                            <textarea type="text" id="caption" v-model="caption" class="form-control" />
                        </div>
                        <div class="mt-2">
                            <input class="form-control" type="file" @change="handleImageChange" accept="image/*" />
                        </div>
                    </div>
                    <div class="mt-2 d-flex justify-content-end">
                        <button style="font-size: 1.3rem;" class="profile-buttons profile-buttons-primary" @click="postPhoto">Post!</button>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<style>

.left-section {
    flex: 1;
    padding-right: 10px;
    border-right: 1px solid #ccc;
}

.right-section {
    flex: 1;
    padding-left: 10px;
}

.image-placeholder {
    width: 20vw;
    height: 30vh;
    display: flex;
    justify-content: center;
    align-items: center;
    border: 1px dashed #000000;
    background-color: rgb(233, 233, 233);
}

.placeholder-content {
    text-align: center;
}
</style>