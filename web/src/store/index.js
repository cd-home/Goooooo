import { createStore } from 'vuex'


const store = createStore({
    state () {
      return {
        user: null,
      }
    },
    getters: {
      
    },
    mutations: {
      login(state, user) {
        state.user = user
      },
      logout (state) {
        state.user = null;
        Cookie.remove("SESSIONID");
      }
    },
    actions: {}
})

export default store