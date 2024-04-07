<script>
export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			isLoggedIn:false,
			feed:null,
			commentContentFeed: null,
			username: null

			
		}
	},
	mounted(){
		const token = localStorage.getItem('token');
		if (token && token !== 'null') {
			this.isLoggedIn = true; // Set isLoggedIn to true if token exists
			

		}
		this.username = sessionStorage.getItem(token);
		this.getFeed()

		
	},
	methods: {
		
		formatDate(time){
			const date = new Date(time);
			
			return date.toLocaleDateString('en-US', { weekday: 'long' }) + " | " + date.toLocaleTimeString('en-US', { timeStyle: 'short' })
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
			this.getFeed();
		},
		async deleteComment(_postId, _commentId){
			this.loading = true;
			this.errormsg = false;
			
			
			const queryParams = {
				postId:_postId,
				commentId: _commentId
			};
			const token = localStorage.getItem('token');
			const headers = {
                'Authorization': 'Bearer ' + token
            };
			try {
				let response = await this.$axios.delete("/post/comments", {params: queryParams, headers:headers});
				

			} catch (e) {
				this.errormsg = e.toString()
			}
			
			
			await this.getFeed();
			
			
			this.loading = false
			this.follow = ''
			this.commentContent = ''
			
			
		
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
		
			this.getFeed()
			
		},
        getImagePath(photoPath) {
			
			const basePath = "images/"+ photoPath + ".jpg";

			return basePath;
		},
		async commentPost(_postId){
			
			this.loading = true;
			this.errormsg = false;
			
			const commentMessage = this.commentContentFeed;
			const queryParams = {
				postId:_postId
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
			
			
			await this.getFeed();
			
			this.loading = false
			this.commentContentFeed = ''
			
			
		},
		async getFeed(){

			this.loading = true;
			this.errormsg = false;
			
			
			
			const token = localStorage.getItem('token');
			const headers = {
                'Authorization': 'Bearer ' + token
            };
			try {
				let response = await this.$axios.get("/users/profile/feed",  {headers});
				this.feed = response.data.message;
                

			} catch (e) {
				this.errormsg = e.toString()
			}
			
        },
    }
}
</script>

<template>
	<div>
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 v-if="!isLoggedIn" class="h2">Log in first</h1>
			<h1 v-else class="h2">Your Feed</h1>
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
        <ul>
            <li class="post-container" v-for="(post, index) in this.feed" :key="post.PostId"> <br>
                <h6>{{ post.OwnerUsername }}</h6>
				<p>Number of comments: {{ post && post.Comments ? post.Comments.length : 0 }}</p>
				<p>Number of likes: {{ post && post.Likes ? post.Likes.length : 0 }}</p>
				
                <img :id="post.PostId" :src="getImagePath(post.PostId)" style="width: 200px; height: auto;">
				<p>Upload: {{formatDate(post.Time)}}</p>
				
				<div class="like-container"> 
					<button  @click="likePhoto(post.PostId)">Like</button> 
					<button  @click="unlikePhoto(post.PostId)">unlike</button> 

					<form @submit.prevent="commentPost(post.PostId)" > 
						<input type="text" id="commentPost" v-model="commentContentFeed" required> <br>
						<input type="submit" value="Submit">
				</form>
				</div>
				
				
                

                
				<ul class="comment-container-feed">
					<li class="comment" v-for="comment in post?.Comments" :key="comment['CommentId']"> 
						<p>{{ comment["Username"] }} : {{ comment["Message"] }} </p> 
						<button  @click="deleteComment(post.PostId, comment['CommentId'])" v-if="comment['Username'] ==this.username">delete</button>
					</li>
				</ul>
			
				
			

            </li>
        </ul>
		
    </div>
	
	
</template>

<style>
.post-container {
	
	height: 500px;
}
.like-container {
	position: relative;
	
}

.comment-container-feed{
	position: relative;
	bottom:385px;
	left:250px;
	width: 250px;
	height: 280px;
	background-color: #3f3f3f;
	overflow-y: scroll;
	overflow-x: scroll;
}
</style>