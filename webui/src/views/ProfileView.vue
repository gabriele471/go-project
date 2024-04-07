<script>
export default {
	data: function() {
		return {
			errormsg: null,
			username: null,
			profile : null,
			notfound:null,
			CommentBoxShowing:null,
			token:null,
			loading: false,
			username: null,
			isLoggedIn:false,
			profile:null,

			CommentBoxShowing:false,
			selectedPost:null,
			commentContent:null
		}
	},
	methods: {
		formatDate(time){
			const date = new Date(time);
			
			return date.toLocaleDateString('en-US', { weekday: 'long' }) + " | " + date.toLocaleTimeString('en-US', { timeStyle: 'short' })
		},
		getUsername(){
			
			return sessionStorage.getItem(localStorage.getItem('token'));
		},
		async commentPost(){
			
			this.loading = true;
			this.errormsg = false;
			
			const commentMessage = this.commentContent;
			const queryParams = {
				postId:this.selectedPost["PostId"]
			};
			const token = localStorage.getItem('token');
			const headers = {
				'Content-Type': 'application/json',
                'Authorization': 'Bearer ' + token
            };
			try {
				let response = await this.$axios.post("/post/comment", commentMessage, {params: queryParams, headers:headers});
				

			} catch (e) {
				this.errormsg = e.toString()
			}
			
			
			await this.searchUser();
			
			this.profile?.Post.forEach(post => {
				if (post["PostId"] == this.selectedPost.PostId) {
					this.selectedPost = post;
				}
			});
			this.loading = false
			this.follow = ''
			this.commentContent = ''
			
			
		},
		async deleteComment(_commentId){
			this.loading = true;
			this.errormsg = false;
			
			
			const queryParams = {
				postId:this.selectedPost["PostId"],
				commentId: _commentId
			};
			const token = localStorage.getItem('token');
			const headers = {
                'Authorization': 'Bearer ' + token
            };
			try {
				let response = await this.$axios.delete("/post/comment", {params: queryParams, headers:headers});
				

			} catch (e) {
				this.errormsg = e.toString()
			}
			
			
			await this.searchUser();
			
			this.profile?.Post.forEach(post => {
				if (post["PostId"] == this.selectedPost.PostId) {
					this.selectedPost = post;
				}
			});
			this.loading = false
			this.follow = ''
			this.commentContent = ''
			
			
		
		},
		toggleCommentBox(postId){
			this.profile?.Post.forEach(post => {
				if (post["PostId"] == postId) {
					this.selectedPost = post;
				}
			});
		
			this.CommentBoxShowing = true;
		},
		async unlikePhoto(postId){
			this.loading = true;
			this.errormsg = false;

			const token = localStorage.getItem('token');
			const config = {
				headers: {
					'Authorization': 'Bearer ' + token
				},
				params: {
					postId:postId
				}
			};
			
			try {
				let response = await this.$axios.delete("/post/like", config);
				
			} catch (e) {
				this.errormsg = e.toString()
			}
			
			this.loading = false;
			this.follow = '';
			this.searchUser()
		},
		getPostLikes(index){
			if (this.profile.Post[index].Likes == null) {
				return 0
			}
			return this.profile.Post[index].Likes.length;
			
			
		},
		async likePhoto(postId){
			this.loading = true;
			this.errormsg = false;

			const queryParams=  {
					postId:postId
			};
			const token = localStorage.getItem('token');
			const headers = {
                'Authorization': 'Bearer ' + token
            };
			try {
				let response = await this.$axios.post("/post/like", null, {params: queryParams, headers:headers});
				
			} catch (e) {
				this.errormsg = e.toString()
			}
			
			this.loading = false;
			this.follow = '';
			this.searchUser()
		},
		getImagePath(photoPath) {
			
			const basePath = "images/"+ photoPath + ".jpg";

			return basePath;
		},
		async searchUser() {
			this.loading = true;
			this.errormsg = null;
            
            const queryParams = {
                username: this.username
            };
            this.token = localStorage.getItem('token');
            
            const headers = {
                'Authorization': 'Bearer ' + this.token
            };
            
			try {
				let response = await this.$axios.get("/users/profile",  {params:queryParams, headers: headers});
				this.profile = response.data.message;
				
                
			} catch (e) {
				this.errormsg = "User not found"
				this.profile = '';
			}
			
			this.notfound = false;
			this.loading = false;
			
		},
	}
}
</script>

<template>
	<div>
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">Profile page</h1>
			<div class="btn-toolbar mb-2 mb-md-0">
				<div class="btn-group me-2">
					<button type="button" class="btn btn-sm btn-outline-secondary" @click="refresh">
						Refresh
					</button>
					<button type="button" class="btn btn-sm btn-outline-secondary" @click="exportList">
						Export
					</button>
				</div>
				<div class="btn-group me-2">
					<button type="button" class="btn btn-sm btn-outline-primary" @click="newItem">
						New
					</button>
				</div>
			</div>
		</div>

		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
	</div>
    <div> 
		<form @submit.prevent="searchUser"> 
			<input type="text" id="username" v-model="username" required> <br>
		</form>
		<button @click="searchUser">Search User Profile</button>
		
		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
		<div v-if = "this.profile" > <br>
			<p> {{ profile?.Username }} </p>
			<div class="scroll-container"> 
				<ul>
					<li v-for="(image, index) in this.profile?.Post" :key="image.PostId"> 
						<img :id="image.PostId" :src="getImagePath(image.PostId)" style="width: 200px; height: auto;">
					
						<button class="buttonLike-scroll" @click="likePhoto(image.PostId)">Like</button> 
						<button class="buttonUnlike-scroll" @click="unlikePhoto(image.PostId)">unlike</button> 
						<button class="comments-scroll" @click="toggleCommentBox(image.PostId)">commments</button> 
						<p style="color:white;">Number of likes: {{getPostLikes(index)}}</p>
						<p style="color:white;">Upload: {{formatDate(image.Time)}}</p>
						

						
					</li>
				</ul>
				<div class="comment-section" v-if="this.CommentBoxShowing"> 
				<button class="exit-button" @click="this.CommentBoxShowing=false">X</button> 
				<ul class="comment-container">
					<li class="comment" v-for="comment in this.selectedPost?.Comments" :key="comment.CommentId"> 
						<p>{{ comment.Username }} : {{ comment.Message }} </p> 
						<button class="deleteComment-button" @click="deleteComment(comment.CommentId)" v-if="comment.Username==getUsername()">delete</button>
					</li>
				</ul>
			</div>
			<div v-if="this.CommentBoxShowing">
				<form class="comment-form" @submit.prevent="commentPost" > 
					<input type="text" id="commentPost" v-model="commentContent" required> <br>
					<input type="submit" value="Submit">
				</form>
			</div>
			</div>
		</div>
		
        
	</div>
</template>

<style>


.comment-form {
	position: absolute;  
	bottom:160px;
	left:638px;
}

.comment {
	margin-top: 10px;
	
	top:30px;
	left:10px;
	color:aliceblue;
}
.exit-button {
	
	left:210px;
	bottom: 270px;
}
.comment-section {
    position: absolute;  
	top:450px;
	left:630px;
	width: 250px;
	height: 300px;
	background-color: #3f3f3f;
	overflow-y: scroll;
	overflow-x: scroll;
}
 
.comments-scroll {
	position:relative;
	top: 70px;
	right: 90px;
}
.buttonUnlike-scroll {
	position:relative;
	bottom:0px;
	right: -5px;
}
.buttonDelete-scroll-container {
	position:relative;
	bottom:120px;
	color:red;
	right: -10px;
}
.buttonLike-scroll {
	position:relative;
	bottom:0px;
	right: -5px;
	
}
.scroll-container {
  overflow-x: hidden;
  overflow-y: scroll;
  background-color: #333;
  
  white-space: nowrap;
  padding: 10px;
  height: 500px;
  width: 400px;
}
</style>
