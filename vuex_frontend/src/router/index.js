import Vue from 'vue'
import Router from 'vue-router'
import store from '../store/index'
import axios from 'axios'
import Mhsw from '@/components/mhsw'
import Regist from '@/components/register'
import Login from '@/components/Login'

Vue.use(Router)

let router = new Router({
  routes: [
    {
      path: '/mhsw',
      name: 'MHSW',
      component: Mhsw,
      meta: {requiresAuth: true},
      // meta: {disabledAction: true}
    },
    {
      path: '/regist',
      name: 'Register',
      component: Regist,
      meta: {disabledAction: true},

    },
    {
      path: '/login',
      name: 'Login',
      component: Login,
      meta: {disabledAction: true}
    }
  ],
  mode: 'history'
})

router.beforeEach((to, from, next) => {
  if (to.matched.some(record => record.meta.requiresAuth)) {
    if (!store.state.auth) {
      next({
        path: '/login',
        query: { redirect: to.fullPath }
      })
    } else {
      next()
    }
  } else {
    if (store.state.auth) {
      next({
        path: '/mhsw',
        query: {redirect: to.fullPath}
      })
    } else {
      next()
    } // make sure to always call next()!
  }
})
export default router

export const HTTP = axios.create({
  baseURL: 'http://192.168.1.8:8002',
  timeout: 1000,
  withCredentials: true,
  headers: {'Content-Type': 'application/json, text/plain, */*',
             'Authorization': 'Bearer '+store.state.token }
})
   HTTP.interceptors.request.use(async(config) =>{
        config.headers.Authorization ='Bearer '+store.state.token
        return config
      }, (error) => {
          return Promise.reject(error)
      })
