import { createRouter, createWebHistory } from 'vue-router'
// import { useUserStore } from '@/stores/user'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      name: 'Home',
      component: () => import('@/components/HelloWorld.vue'),
    },
    {
      path: '/login',
      name: 'Login',
      component: () => import('@/views/Login.vue'),
    },
  ],
})

export default router

// router.beforeEach((to: RouteLocationNormalizedLoaded, from, next) => {
//   const userStore = useUserStore()
//   if (to.meta.requiresAuth && !userStore.isAuthenticated) {
//     next('/login')
//   } else {
//     next()
//   }
// })
