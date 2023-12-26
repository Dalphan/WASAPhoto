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
        currentUser: null,
        commentText: null,
      }
	  },
    watch:{
      photo: {
          immediate: true,
          handler(newVal, oldVal){
            if (newVal != oldVal) {
              this.$timestampToDate(newVal.timestamp);
              this.getComments();
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
      <div class="modal">
        <transition name="modal-animation-inner">
          <div class="modal-inner">      
            <div class="container">
              <div class="row">
                <div class="col-md-6">
                  <img :src="photo.image" alt="Photo" class="img-fluid" ref="image">
                </div>

                <div class="col-md-6">   
                  <div class="photo-details-container" ref="detailsContainer">     
                    <div class="d-flex justify-content-between">
                      <h4>{{ photo.username }}</h4>
                      <svg class="feather" role="button" @click="close"><use href="/feather-sprite-v4.29.0.svg#x"/></svg>
                    </div>
                    <p v-if="photo.caption">{{ photo.caption }}</p>
                    <div class="d-flex justify-content-between align-items-center">
                      <div @click="putRemoveLike" class="d-flex justify-content-between align-items-center" role="button">
                        <svg xmlns="http://www.w3.org/2000/svg" width="22" height="22" fill="currentColor" class="bi bi-heart-fill text-danger" viewBox="0 0 16 16">
                          <path v-if="photo.liked" fill-rule="evenodd" d="M8 1.314C12.438-3.248 23.534 4.735 8 15-7.534 4.736 3.562-3.248 8 1.314"/>
                          <path v-else d="m8 2.748-.717-.737C5.6.281 2.514.878 1.4 3.053c-.523 1.023-.641 2.5.314 4.385.92 1.815 2.834 3.989 6.286 6.357 3.452-2.368 5.365-4.542 6.286-6.357.955-1.886.838-3.362.314-4.385C13.486.878 10.4.28 8.717 2.01zM8 15C-7.333 4.868 3.279-3.04 7.824 1.143c.06.055.119.112.176.171a3.12 3.12 0 0 1 .176-.17C12.72-3.042 23.333 4.867 8 15"/>
                        </svg>
                        &times;{{ photo.likeCount }}
                      </div>
                      <small> {{ photo.timestamp }} </small>
                    </div>

                    <div>Comments: {{ photo.commentCount }}</div>
                    <div class="input-group">
                      <input v-model="commentText" type="text" class="form-control input-comment" placeholder="Write a comment..." aria-label="Write a comment..." aria-describedby="send">
                      <div class="input-group-append">
                        <span @click="sendComment" class="input-group-text text-primary send-comment" id="send">
                          <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#send"/></svg>
                        </span>
                      </div>
                    </div>
                    <ul class="list-group comment-list">
                      <li class="list-group-item" v-for="(c, index) in comments" :key="c.id"> 
                        <div class="d-flex justify-content-between align-items-center">
                          <strong>{{ c.username }}</strong><br>
                          <svg v-if="c.user == currentUser" class="feather text-danger" role="button" @click="deleteComment(c.id, index)"><use href="/feather-sprite-v4.29.0.svg#trash-2"/></svg>
                        </div>
                        {{ c.text }}<br>
                        <small class="form-text font-italic"> {{ c.date }}</small>
                      </li>
                    </ul>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </transition>
      </div>
    </transition>
</template>


<style lang="css" scoped>
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

.modal {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  width: 100vw;
  position: fixed;
  top: 0;
  left: 0;
  background-color: rgba(255, 255, 255, 0.7);
}

.modal-inner {
  position: relative;
  max-width: auto;
  box-shadow: 4px 4px 6px 4px rgba(0, 0, 0, 0.1), 4px 4px 4px 4px rgba(0, 0, 0, 0.06);
  background-color: #fff;
  padding: 30px 24px;
}

.photo-details-container {
  display: flex;
  flex-direction: column;
  height: 100%;
}

.comment-list {
  overflow-y: auto; /* Add a scrollbar if comment list exceeds the max height */
}

.input-comment {
  font-size: 0.9rem;
}

.send-comment {
  height: 100%;
  cursor: pointer;
  border-top-left-radius: 0;
  border-bottom-left-radius: 0;
  background-color: white;
}

.send-comment:hover {
  background-color: rgba(13,110,253, 1) !important;
  color:white !important;
}
</style>