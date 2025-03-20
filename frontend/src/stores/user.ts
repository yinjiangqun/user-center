import { defineStore } from 'pinia'

export const useUserStore = defineStore('user', {
  state: () => ({
    user: null,
    token: null,
    isAuthenticated: false
  }),
  actions: {
    login(userData) {
      this.user = userData
      this.isAuthenticated = true
    },
    logout() {
      this.user = null
      this.isAuthenticated = false
    }
  }
})