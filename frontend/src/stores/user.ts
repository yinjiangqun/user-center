import { defineStore } from 'pinia'

export const useUserStore = defineStore('user', {
  state: () => ({
    user: null,
    token: null,
    isAuthenticated: false,
  }),
  actions: {
    login() {
      this.user = null
      this.isAuthenticated = true
    },
    logout() {
      this.user = null
      this.isAuthenticated = false
    },
  },
})
