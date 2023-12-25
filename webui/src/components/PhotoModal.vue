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
        comments: null,
        currentUser: null,
      }
	  },
    watch:{
      photo: {
          immediate: true,
          handler(newVal, oldVal){
            if (newVal != oldVal) {
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

            this.comments = response.data;
            this.comments.forEach( e => {
              e.date = this.$timestampToDate(e.date);
            });
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
            console.log(response);
            this.comments.splice(index, 1);
          }
        } catch (e) {
          this.errormsg = e.toString();				
        }
        this.loading = false;
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
                      <p><strong>Likes:</strong> {{ photo.likeCount }}</p>
                      <p> {{ photo.timestamp }} </p>
                    </div>

                    <h6>Comments: {{ photo.commentCount }}</h6>
                    <ul class="list-group comment-list">
                      <li class="list-group-item" v-for="(c, index) in comments" :key="c.id"> 
                        <div class="d-flex justify-content-between align-items-center">
                          <strong>{{ c.username }}</strong><br>
                          <svg v-if="c.user == currentUser" class="feather" role="button" @click="deleteComment(c.id, index)"><use href="/feather-sprite-v4.29.0.svg#trash-2"/></svg>
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
</style>