<script>
export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			token: null,
			username: null,
			isLoggedIn:false,
			profile:null,
			newUsername:null,
			responseCode:null,
			banned:null,
			unfollow:null,
			follow:null,
			unban:null,
			selectedFile:null,
			CommentBoxShowing:false,
			selectedPost:null,
			commentContent:null

			
		}
	},
	mounted(){
		const token = localStorage.getItem('token');
		if (token && token !== 'null') {
			this.isLoggedIn = true; // Set isLoggedIn to true if token exists
			

		}
		if (this.isLoggedIn){
			this.searchUser()
		}

		
	},
	methods: {
		formatDate(time){
			const date = new Date(time);
			
			
			return date.toLocaleDateString('en-US', { weekday: 'long' }) + " | " +  date.toLocaleTimeString('en-US', { timeStyle: 'short' })
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
				let response = await this.$axios.post("/post/comments", commentMessage, {params: queryParams, headers:headers});
				

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
				let response = await this.$axios.delete("/post/likes", config);
				
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
				let response = await this.$axios.post("/post/likes", null, {params: queryParams, headers:headers});
				
			} catch (e) {
				this.errormsg = e.toString()
			}
			
			this.loading = false;
			this.follow = '';
			this.searchUser()
		},
		async deletePhoto(postId){
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
				let response = await this.$axios.delete("/users/profile/post", config);
			} catch (e) {
				this.errormsg = e.toString()
			}
			this.loading = false;
			this.unban = '';
			this.searchUser()
		},
		getImagePath(photoPath) {
			
			const basePath = "images/"+ photoPath + ".jpg";

			return basePath;
		},
		async uploadPhoto(){
			this.loading = true;
			this.errormsg = false;
			
			const token = localStorage.getItem('token');

			const formData = new FormData();
			formData.append('image', this.selectedFile);
			
			const headers = {
                'Authorization': 'Bearer ' + token
            };
			try {
				let response = await this.$axios.post("/users/profile/post", formData, {headers:headers});
			} catch (e) {
				this.errormsg = e.toString()
			}
			this.loading = false;
			this.unban = '';
			this.searchUser();
			
		},
		selectPhoto(event){
			const file = event.target.files[0];
			this.selectedFile = file;
		},
		async unBan(){
			this.loading = true;
			this.errormsg = false;
			const queryParams = {
				username:this.unban
			};

			const token = localStorage.getItem('token');
			const headers = {
                'Authorization': 'Bearer ' + token
            };
			try {
				let response = await this.$axios.delete("/users/profile/ban", {params: queryParams, headers:headers});
			} catch (e) {
				this.errormsg = e.toString()
			}
			this.loading = false;
			this.unban = '';
		},
		async unfollowUser(){
			this.loading = true;
			this.errormsg = false;
			const queryParams = {
				username:this.unfollow
			};
			const token = localStorage.getItem('token');
			const headers = {
                'Authorization': 'Bearer ' + token
            };
			try {
				let response = await this.$axios.delete("/users/profile/followers", {params: queryParams, headers:headers});

			} catch (e) {
				this.errormsg = e.toString()
			}
			this.loading = false;
			this.unfollow = '';
			this.searchUser()
		},
		async followUser(){
			this.loading = true;
			this.errormsg = false;
			const queryParams = {
				username:this.follow
			};
			const token = localStorage.getItem('token');
			const headers = {
                'Authorization': 'Bearer ' + token
            };
			try {
				let response = await this.$axios.post("/users/profile/followers", null, {params: queryParams, headers:headers});

			} catch (e) {
				this.errormsg = e.toString()
			}
			this.loading = false;
			this.follow = '';
			this.searchUser()
		},
		async banUser(){
			this.loading = true;
			this.errormsg = false;
			const queryParams = {
				username:this.banned
			};
			const token = localStorage.getItem('token');
			const headers = {
                'Authorization': 'Bearer ' + token
            };
			try {
				let response = await this.$axios.post("/users/profile/ban", null, {params: queryParams, headers:headers});

			} catch (e) {
				this.errormsg = e.toString()
			}
			this.loading = false;
			this.banned = '';
			
		},
		async changeUsername(){
			this.loading = true;
			this.errormsg = false;
			
			const queryParams = {
				username:this.newUsername
			};
			const token = localStorage.getItem('token');
			const headers = {
                'Authorization': 'Bearer ' + token
            };
			try {
				let response = await this.$axios.post("/users/profile/username", null, {params: queryParams, headers:headers});

			} catch (e) {
				this.errormsg = e.toString()
			}
			sessionStorage.setItem(token, this.newUsername)
			this.newUsername = '';
			this.searchUser()
		},
		async login() {
			this.loading = true;
			this.errormsg = null;
			

			const queryParams = {
				username: this.username
			};
			try {
				let response = await this.$axios.post("/session",null,{params: queryParams});
				this.token = response.data.message;
				localStorage.setItem('token', this.token);
				sessionStorage.setItem(this.token, this.username);

			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
			this.isLoggedIn = true;
			this.searchUser();
		},
		logout() {
			// Clear token from local storage and update isLoggedIn
			localStorage.removeItem('token'); 
			this.isLoggedIn = false;
		},
		async searchUser() {
			this.loading = true;
			this.errormsg = null;
            
            
            const token = localStorage.getItem('token');
            this.username = sessionStorage.getItem(token);
			const queryParams = {
                username: this.username
            };
            const headers = {
                'Authorization': 'Bearer ' + token
            };
            
			try {
				let response = await this.$axios.get("/users/profile",  {params:queryParams, headers: headers});
				this.profile = response.data.message; 
				
                
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
			
			
			
			
		}
	},
	
}
</script>

<template>
	<div>
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 v-if="!isLoggedIn" class="h2">Login page</h1>
			<h1 v-else class="h2">Personal page</h1>
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
	
	<div v-if="!isLoggedIn"> 
		<form @submit.prevent="login"> 
			<input type="text" id="username" v-model="username" required> <br>
		</form>
		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
		<button @click="login">Login</button>

	</div>
	<div v-else>

		<button @click="logout">Log out</button> <br> <br>
		<p>Hello user {{ profile?.Username }}</p>
		<p>Number of posts:{{ profile?.Npost }}</p>
		<div class="following-container">
			<p>Following:</p>
			<ul>
				<li v-for="(user) in profile?.Following">
					{{ user.Username }}
				</li>
			</ul>
		</div>
		<div class="follower-container">
			<p>Followers:</p>
			<ul>
				<li v-for="(user) in profile?.Follower" :key="user.Id">
					{{ user.Username }}
				</li>
			</ul>
		</div>
		<div class="uploadPhoto-container">
			<input type="file" @change="selectPhoto">
			<button @click="uploadPhoto">Upload photo</button>
		</div>


		
		<div class="newUsername-container">
			<form @submit.prevent="changeUsername"> 
				<label for="newUsername">New Username:</label> <br>
				<input type="text" id="newUsername" v-model="newUsername" required> <br>
				<input type="submit" value="Change">
			</form>
		</div>

		<div class="ban-container">
			<form @submit.prevent="banUser"> 
				<label for="banned">Ban User:</label> <br>
				<input type="text" id="banned" v-model="banned" required> <br>
				<input type="submit" value="Submit">
			</form>
		</div>

		<div class="follow-container">
			<form @submit.prevent="followUser"> 
				<label for="follow">Follow User:</label> <br>
				<input type="text" id="follow" v-model="follow" required> <br>
				<input type="submit" value="Submit">
			</form>
		</div>
		<div class="unfollow-container">
			<form @submit.prevent="unfollowUser"> 
				<label for="unfollow">Unfollow User:</label> <br>
				<input type="text" id="unfollow" v-model="unfollow" required> <br>
				<input type="submit" value="Submit">
			</form>
		</div>

		<div class="unban-container">
			<form @submit.prevent="unBan"> 
				<label for="unban">Unban User:</label> <br>
				<input type="text" id="unban" v-model="unban" required> <br>
				<input type="submit" value="Submit">
			</form>
		</div>
		<div class="scroll-container"> 
			
			<ul>
				<li v-for="(image, index) in this.profile?.Post" :key="image.PostId"> 
					<img :id="image.PostId" :src="getImagePath(image.PostId)" style="width: 200px; height: auto;">
					<button class="buttonDelete-scroll-container" @click="deletePhoto(image.PostId)">Delete</button> 
					<button class="buttonLike-scroll-container" @click="likePhoto(image.PostId)">Like</button> 
					<button class="buttonUnlike-scroll-container" @click="unlikePhoto(image.PostId)">unlike</button> 
					<button class="comments-scroll-container" @click="toggleCommentBox(image.PostId)">commments</button> 
					<p style="color:white;">Number of likes: {{getPostLikes(index)}}</p>
					<p style="color:white;">Upload: {{ formatDate(image.Time) }}</p>
					

					
				</li>
			</ul>
			<div class="comment-section" v-if="this.CommentBoxShowing"> 
				<button class="exit-button" @click="this.CommentBoxShowing=false">X</button> 
				<ul class="comment-container">
					<li class="comment" v-for="comment in this.selectedPost?.Comments" :key="comment.CommentId"> 
						<p>{{ comment.Username }} : {{ comment.Message }} </p> 
						<button class="deleteComment-button" @click="deleteComment(comment.CommentId)" v-if="comment.Username==this.profile.Username">delete</button>
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
 
.comments-scroll-container {
	position:relative;
	top: 70px;
	right: 145px;
}
.buttonUnlike-scroll-container {
	position:relative;
	bottom:0px;
	right: 50px;
}
.buttonDelete-scroll-container {
	position:relative;
	bottom:120px;
	color:red;
	right: -10px;
}
.buttonLike-scroll-container {
	position:relative;
	bottom:0px;
	right: 50px;
	
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
.uploadPhoto-container {
  position: absolute; 
  top: 130px; 
  right: 220px;
  margin-left: 10000px;
}
.follower-container{
  position: absolute; 
  top: 500px; 
  right: 50px;
 
}
.following-container{
  position: absolute; 
  top: 500px; 
  right: 200px;
 
}
.ban-container {
  position: absolute;
  top: 200px; 
  right: 250px;
  padding: 20px; 
  box-sizing: border-box;
}
.follow-container {
  position: absolute;
  top: 300px; 
  right: 250px;
  padding: 20px; 
  box-sizing: border-box;
}
.unfollow-container {
  position: absolute;
  top: 300px; 
  right: 10px;
  padding: 20px;
  box-sizing: border-box;
}
.newUsername-container {
  position: absolute;
  top: 100px;
  right: 10px;
  padding: 20px; 
  box-sizing: border-box;
}
.unban-container {
  position: absolute;
  top: 200px; 
  right: 10px;
  padding: 20px;
  box-sizing: border-box;
}
</style>