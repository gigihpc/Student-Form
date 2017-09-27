<template>
  <v-form ref="form">
      <v-layout justify-center>
    <v-flex xs12 sm10 md8 lg6>
      <v-alert error :value="isfail"
      transition="scale-transition">{{isfail}}</v-alert>
      <!-- <li>{{ isfail }}</li> -->
    <v-text-field
      label="Email"
      v-model="email"
      :rules="[rules.required, rules.email]"
      required
    ></v-text-field>
    <v-text-field
      label="Password"
      v-model="password"
      :append-icon="isShow ? 'visibility' : 'visibility_off'"
              :append-icon-cb="() => (isShow = !isShow)"
              :type="isShow ? 'text' : 'password'"
              counter
      :rules="[rules.required]"
      required
    ></v-text-field>
    
    <v-btn @click="Login">LOGIN</v-btn>
    <td><router-link to="regist">Create User</router-link></td>
    <router-view></router-view>
    </v-flex>
    </v-layout>
  </v-form>
</template>

<script>
import {HTTP} from '@/router/index'
export default{
  data () {
    return {
      email: null,
      password: null,
      user_list: [],
      isShow: false,
      isfail: '',
      rules: {
        required: (value) => !!value || 'Required.',
        email: (value) => {
          const pattern = /^(([^<>()\]\\.,;:\s@"]+(\.[^<>()\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
          return pattern.test(value) || 'Invalid e-mail.'
        }
      }
    }
  },
  methods: {
    Login: function () {
      let email = this.email
      let pass = this.password
    //   console.log('var: ', email, pass)
      HTTP.get('/api/user').then(response => {
        this.user_list = response.data.data
        this.user_list.forEach(function (element) {
          // console.log('from db: ', element.id, element.name, element.email, element.password, element.country)
          if (email === element.email && pass === element.password) {
            // console.log('Login Success!')
            localStorage.setItem('auth', '1234')
            this.$router.push('mhsw')
            // this.$router.replace({path: '/'})
          } else {
            this.isfail = 'Login Failed !'
            if (email === element.email && pass !== element.password) {
              this.isfail = this.isfail + ' Password wrong'
            } else if (email !== element.email && pass === element.password) {
              this.isfail = this.isfail + ' Email didn`t create'
            } else {
              this.isfail = this.isfail + ' User didn`t create'
            }
            // console.log('Login Failed!!')
          }
        }, this)
      })
      .catch(err => {
        console.log(err.response)
      })
    }
  }
}
</script>