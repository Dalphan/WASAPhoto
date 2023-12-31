<script>
  export default {
    name: 'Modal',
    props: {
        photo: Object,
    },
    data: function() {
        return {
            errormsg: null,
            loading: false,
            comments: [],
            likes: [],
            currentUser: null,
            commentText: null,
        }
    },
    watch:{
        photo: {
            immediate: true,
            handler(newVal, oldVal){
                if (newVal != oldVal) {
                    // this.$timestampToDate(newVal.timestamp);
                    this.getComments();
                    this.getLikes();
                }
            }
        }
    },	
    methods: {
        close() {
            this.$emit('close');
        },
      
        async getComments(){
            this.loading = true;
            this.errormsg = null;
            var path = `/photos/${this.photo.id}/comments`;
            try {
            let response = await this.$axios.get(path);
            if (response.status == 200) {
                // Adapt the height to that of the image
                this.$refs.detailsContainer.style.maxHeight = this.$refs.image.offsetHeight + 'px';

                this.currentUser = this.$getCurrentId();

                if (response.data !== null) {
                this.comments = response.data;
                this.comments.forEach( e => {
                    e.date = this.$timestampToDate(e.date);
                });
                }
            }
            } catch (e) {
            this.errormsg = e.toString();				
            }
            this.loading = false;
        },

        async getLikes(){
            this.loading = true;
            this.errormsg = null;
            var path = `/photos/${this.photo.id}/likes`;
            try {
            let response = await this.$axios.get(path);
            if (response.status == 200) {
                // Adapt the height to that of the image
                //this.$refs.detailsContainer.style.maxHeight = this.$refs.image.offsetHeight + 'px';

                this.currentUser = this.$getCurrentId();

                if (response.data !== null) {
                this.likes = response.data;
                console.log(this.likes);
                }
            }
            } catch (e) {
            this.errormsg = e.toString();				
            }
            this.loading = false;
        },

        async deleteComment(id, index){
            this.loading = true;
            this.errormsg = null;
            var path = `/photos/${this.photo.id}/comments/${id}`;
            try {
            let response = await this.$axios.delete(path);
            if (response.status == 200) {
                this.comments.splice(index, 1);
                this.photo.commentCount--;
            }
            } catch (e) {
            this.errormsg = e.toString();				
            }
            this.loading = false;
        },

        async sendComment(){
            this.loading = true;
            this.errormsg = null;
            var path = `/photos/${this.photo.id}/comments`;
            this.currentUser = this.$getCurrentId();
            try {
            let response = await this.$axios.post(path, {
                text: this.commentText,
            });

            if (response.status == 200 || response.status == 201) {
                // Add comment to list of comments
                this.commentText = null;
                let comment = response.data;
                comment.date = this.$timestampToDate(comment.date);
                comment.username = this.$getCurrentUsername();
                if (this.comments.length > 0) this.comments.splice(0, 0, comment);
                else this.comments.push(comment);
                this.photo.commentCount++;
            }
            } catch (e) {
            this.errormsg = e.toString();				
            }
            this.loading = false;
        },

        async putLike(){
            this.loading = true;
            this.errormsg = null;
            this.currentUser = this.$getCurrentId();
            var path = `/photos/${this.photo.id}/likes/${this.currentUser}`;
            try {
            let response = await this.$axios.put(path);
            if (response.status == 200 || response.status == 201) {
                this.photo.liked = true;
                this.photo.likeCount++;

                let like = response.data;
                like.username = this.$getCurrentUsername();

                // Add like to list of likes
                if (this.likes.length > 0) this.likes.splice(0, 0, like);
                else this.likes.push(like);
            }
            } catch (e) {
            this.errormsg = e.toString();				
            }
            this.loading = false;
        },

        async removeLike(){
            this.loading = true;
            this.errormsg = null;
            var path = `/photos/${this.photo.id}/likes/${this.currentUser}`;
            try {
            let response = await this.$axios.delete(path);
            if (response.status == 200) {
                this.photo.liked = false;
                this.photo.likeCount--;

                // Remove like from the list
                this.likes = this.likes.filter(like => !(like.user === response.data.user && like.photo === response.data.photo));
            }
            } catch (e) {
            this.errormsg = e.toString();				
            }
            this.loading = false;
        },

        async putRemoveLike(){
            if (this.photo.liked) this.removeLike();
            else this.putLike();
        }
    },
};
</script>

<template>
  <transition name="modal-animation">
    <div class="modal" @click="close">
      <transition name="modal-animation-inner">
        <div class="modal-inner" @click.stop="">      
          <div class="container">
            <div class="row">
              <div class="col-md-6">
                <img :src="photo.image" alt="Photo" class="img-fluid" ref="image">
              </div>

              <div class="col-md-6">   
                <div class="photo-details-container" ref="detailsContainer">     
                  <div class="d-flex justify-content-between">
                      <RouterLink :to="this.$pathToProfile(photo.username)" class="user-link">
                          <h4>{{ photo.username }}</h4>
                      </RouterLink>
                      <svg class="feather" role="button" @click="close"><use href="/feather-sprite-v4.29.0.svg#x"/></svg>
                  </div>
                  <p v-if="photo.caption">{{ photo.caption }}</p>
                  <div class="d-flex justify-content-between align-items-center" style="margin-bottom: 5px;">
                    <div class="d-flex justify-content-between align-items-center" role="button">
                      <svg @click="putRemoveLike" xmlns="http://www.w3.org/2000/svg" width="22" height="22" fill="currentColor" class="bi bi-heart-fill text-danger" viewBox="0 0 16 16">
                        <path v-if="photo.liked" fill-rule="evenodd" d="M8 1.314C12.438-3.248 23.534 4.735 8 15-7.534 4.736 3.562-3.248 8 1.314"/>
                        <path v-else d="m8 2.748-.717-.737C5.6.281 2.514.878 1.4 3.053c-.523 1.023-.641 2.5.314 4.385.92 1.815 2.834 3.989 6.286 6.357 3.452-2.368 5.365-4.542 6.286-6.357.955-1.886.838-3.362.314-4.385C13.486.878 10.4.28 8.717 2.01zM8 15C-7.333 4.868 3.279-3.04 7.824 1.143c.06.055.119.112.176.171a3.12 3.12 0 0 1 .176-.17C12.72-3.042 23.333 4.867 8 15"/>
                      </svg>
                    </div>
                    <small> {{ photo.timestamp }} </small>
                  </div>
                  <div>
                    <ul class="nav nav-tabs nav-justified">
                      <li class="nav-item" role="presentation">
                        <button class="nav-link active" id="comment-tab" data-bs-toggle="tab" data-bs-target="#comments" type="button" role="tab" aria-controls="home" aria-selected="true"> 
                          <svg class="feather" style="width: 23px; height: 23px;"><use href="/feather-sprite-v4.29.0.svg#message-circle"/></svg>
                          {{ photo.commentCount }} comments 
                        </button>
                      </li>
                      <li class="nav-item" role="presentation">
                        <button class="nav-link" id="comment-tab" data-bs-toggle="tab" data-bs-target="#likes" type="button" role="tab" aria-controls="home" aria-selected="true">  
                          <svg class="feather" style="width: 23px; height: 23px;"><use href="/feather-sprite-v4.29.0.svg#heart"/></svg>
                          {{ photo.likeCount }} Likes
                        </button>
                      </li>
                    </ul>

                  </div>
                  <div class="tab-content" style="overflow-y: auto;">
                    <div id="comments" class="tab-pane fade show active" role="tabpanel" >
                      <div class="input-group">
                        <input v-model="commentText" type="text" class="form-control input-comment" placeholder="Write a comment..." aria-label="Write a comment..." aria-describedby="send">
                        <div class="input-group-append">
                          <span @click="sendComment" class="input-group-text text-primary send-button" id="send">
                            <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#send"/></svg>
                          </span>
                        </div>
                      </div>
                      <ul class="list-group">
                        <li class="list-group-item" v-for="(c, index) in comments" :key="c.id">
                          <div class="comment-container"> 
                            <div class="d-flex justify-content-between align-items-center">
                              <RouterLink :to="this.$pathToProfile(c.username)" class="user-link">
                                  <strong>{{ c.username }}</strong><br>
                              </RouterLink>
                              <svg v-if="c.user == currentUser" class="feather text-danger" role="button" @click="deleteComment(c.id, index)"><use href="/feather-sprite-v4.29.0.svg#trash-2"/></svg>
                            </div>
                            {{ c.text }}<br>
                            <small class="form-text font-italic"> {{ c.date }}</small>
                          </div>
                        </li>
                      </ul>
                    </div>
                    <div id="likes" class="tab-pane fade" role="tabpanel">
                      <ul class="list-group">
                        <li class="list-group-item" v-for="(l, index) in likes" :key="l.user">
                          <div class="comment-container"> 
                            <div class="d-flex">
                              <RouterLink :to="this.$pathToProfile(l.username)" class="user-link">
                                  <strong>{{ l.username }}</strong><br>
                              </RouterLink>
                            </div>
                          </div>
                        </li>
                      </ul>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </transition>
    </div>
  </transition>
</template>


<style>
/* .modal-animation-enter-active,
.modal-animation-leave-active {
  transition: opacity 0.3s cubic-bezier(0.52, 0.02, 0.19, 1.02);
}

.modal-animation-enter-from,
.modal-animation-leave-to {
  opacity: 0;
}

.modal-animation-inner-enter-active {
  transition: all 0.3s cubic-bezier(0.52, 0.02, 0.19, 1.02) 0.15s;
}

.modal-animation-inner-leave-active {
  transition: all 0.3s cubic-bezier(0.52, 0.02, 0.19, 1.02);
}

.modal-animation-inner-enter-from {
  opacity: 0;
  transform: scale(0.8);
}

.modal-animation-inner-leave-to {
  transform: scale(0.8);
} */

.photo-details-container {
  display: flex;
  flex-direction: column;
  height: 100%;
}

.input-comment {
  font-size: 0.85rem;
}

.comment-container {
  max-width: 100%; /* Adjust as needed */
  word-wrap: break-word; /* or word-wrap: normal; depending on your preference */
}

.likes {
  margin-left: 5px;
}

.likes:hover {
  text-decoration: underline;
}
</style>