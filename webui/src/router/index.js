import {createRouter, createWebHashHistory} from 'vue-router'
import FeedView from '../views/FeedView.vue'
import LoginView from '../views/LoginView.vue'
import ProfileView from '../views/ProfileView.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/session', component: LoginView}, 
		{path: '/users/profile', component: ProfileView},
		{path: '/users/profile/feed', component: FeedView},
	]
})

export default router
