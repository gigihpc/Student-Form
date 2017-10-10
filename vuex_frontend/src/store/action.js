 import {HTTP} from '../router/index'

 export const Authenticate = ({commit}, payload) => {
     return new Promise((resolve, reject) => {
         let credential = payload || {}
         HTTP.post('/api/user_auth', credential).then(response => {
            commit('loading', true)
            commit('authenticated', response.data)
            commit('authguard', true)
   
            setTimeout(()=> {
                commit('authenticated', response.data)
                commit('authguard', true)
                commit('loading', false)
                resolve(response.data)
            }, 1000)
             
         })
         .catch(err => {
             commit('statusLogin', err.response.statusText)
             console.log('error: ', err.response)
            //  reject(err)
         })
    })
}

export const Authenticate1 = ({commit}, payload) => {
    return  HTTP.post('/api/user_auth', payload).then(response => {
       commit('authenticated', response.data)
       commit('authguard', 'auth')
       // commit('loading', true)

       setTimeout(()=> {
           commit('authenticated', response.data)
           commit('authguard', 'auth')
        //   commit('loading', false)
       }, 1000)
        
    })
    .catch(err => {
       console.log('error: ', err.response)
    }) 
}