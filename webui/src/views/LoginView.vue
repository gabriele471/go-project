<script>
export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			token: null
		}
	},
	methods: {
		async submitForm() {
			this.loading = true;
			this.errormsg = null;

			const queryParams = {
				username: this.username
			}
			try {
				let response = await this.$axios.post("/session", null, {params: queryParams});
				this.token = response.data.message
				localStorage.setItem('token', this.token);

			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		},
	},
	
}
</script>

<template>
	<div>
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">Login page</h1>
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
		<form @submit.prevent="submitForm"> 
			<label for="username"> Username: </label><br>
			<input type="text" id="username" v-model="username" required> <br>
		</form>
		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
		<button @click="submitForm">Submit Form</button>
	</div>
</template>

<style>
</style>
