import {createRouter, createWebHashHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from '../views/LoginView.vue'
import ProfileView from '../views/ProfileView.vue'
import SearchView from '../views/SearchView.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', component: LoginView, name: 'Login'},
		{path: '/home', component: HomeView, name: 'Home'},
		{path: '/user/:username', component: ProfileView, name: 'Profile'},
		{path: '/search', component: SearchView, name: 'Search'},
	]
})

export default router
