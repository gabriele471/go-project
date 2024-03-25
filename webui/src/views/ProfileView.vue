<script>
export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			username: null,
            requestSent: false
		}
	},
	methods: {
		async searchUser() {
			this.loading = true;
			this.errormsg = null;
            this.requestSent = false
            const queryParams = {
                username: this.username
            }
            const token = localStorage.getItem('token')
            
            const headers = {
                'Authorization': 'Bearer ' + token
            }
            
			try {
				let response = await this.$axios.get("/users/profile",  {params:queryParams, headers: headers});
				this.profile_data = response.data.message;
                
			} catch (e) {
				this.errormsg = e.toString();
			}
            this.requestSent = true
			this.loading = false;
		},
	},
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
		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
		<button @click="searchUser">Search User Profile</button>

        <div v-if="this.requestSent"> <br>
            <p>Username: {{profile_data.Username }}</p>
            <p>Number of posts: {{profile_data.Npost }}</p>
        </div>
	</div>
</template>

<style>
</style>
