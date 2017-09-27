import Vue from 'vue'
import Router from 'vue-router'
import axios from 'axios'
import Hello from '@/components/Hello'
import Form from '@/components/Form'
import Mhsw from '@/components/mhsw'
import Regist from '@/components/register'
import Login from '@/components/Login'

Vue.use(Router)

let router = new Router({
  routes: [
    {
      path: '/hello',
      name: 'Hello',
      component: Hello
    },
    {
      path: '/form',
      name: 'Form',
      component: Form
    },
    {
      path: '/mhsw',
      name: 'MHSW',
      component: Mhsw,
      meta: {requiresAuth: true}
    },
    {
      path: '/regist',
      name: 'Register',
      component: Regist,
      meta: {disabledAction: true}
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
    if (!localStorage.getItem('auth')) {
      next({
        path: '/login',
        query: { redirect: to.fullPath }
      })
    } else {
      next()
    }
  } else {
    if (localStorage.getItem('auth')) {
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
  baseURL: 'http://192.168.1.8:8001',
  timeout: 1000,
  withCredentials: true,
  headers: {'Content-Type': 'application/json, text/plain, */*'}
})
