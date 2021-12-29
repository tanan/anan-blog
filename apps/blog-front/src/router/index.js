import { createRouter, createWebHistory } from 'vue-router'
import Home from '../components/pages/Home.vue'
import Article from '../components/pages/Article.vue'
import About from '../components/pages/About.vue'
import Profile from '../components/pages/Profile.vue'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/article/:id',
    name: 'Article',
    component: Article
  },
  {
    path: '/about/:id',
    name: 'About',
    component: About
  },
  {
    path: '/profile',
    name: 'Profile',
    component: Profile
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
