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
			selectedFile:null
			
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
		async uploadPhoto(){
			this.loading = true;
			this.errormsg = false;
			
			const token = localStorage.getItem('token');

			const formData = new FormData();
			formData.append('image', this.selectedFile);
			console.log(this.selectedFile)
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
				sessionStorage.setItem(this.token, this.username)

			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
			this.isLoggedIn = true;
			this.searchUser()
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
			console.log(this.username);
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
			<h1 v-else class="h2">Profile page</h1>
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
				<li v-for="(user) in profile.Follower">
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
		
	</div>
	
</template>

<style>
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