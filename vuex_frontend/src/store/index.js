import Vue from 'vue'
import Vuex from 'vuex'
import * as getters from '../store/getters'
import * as actions from '../store/action'
import * as mutations from '../store/mutations'

Vue.use(Vuex)

const state = {
    auth: localStorage.getItem('auth'),
    token: localStorage.getItem('token'),
    status: '',
    loading: false
  }
  const store = new Vuex.Store({
    state,
    getters,
    actions,
    mutations
  })
  
  if (module.hot) {
    module.hot.accept([
      '../store/getters.js',
      '../store/action.js',
      '../store/mutations.js'
    ], () => {
      store.hotUpdate({
        getters: require('../store/getters'),
        actions: require('../store/action'),
        mutations: require('../store/mutations')
      })
    })
  }
  export default store