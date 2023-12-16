import {createRouter, createWebHashHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue'
import ProfileView from '../views/ProfileView.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', component: HomeView},
		{path: '/user', component: ProfileView},
		{path: '/link2', component: HomeView},
		{path: '/some/:id/link', component: HomeView},
	]
})

export default router
