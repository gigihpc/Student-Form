<template>
  <v-container align-start fluid>
    <v-layout row wrap justify-center>
      <v-flex xs10 offset-xs1 sm4>
        <v-card>
          <v-app top-toolbar style="min-height: 0;" standalone>
          <v-toolbar color="teal" app>
          </v-toolbar>
          <main>
          <v-container fluid>
            <v-layout row wrap>
              <v-flex xs10>
                <v-alert error :value=getStatusLogin transition="scale-transition">{{getStatusLogin}}</v-alert>
                <v-progress-circular v-if="loading" indeterminate v-bind:size="50" class="primary--text"></v-progress-circular>
              </v-flex>
              <v-flex xs10>
                <v-text-field label="Email" v-model="email" :rules="[rules.required, rules.email]" required></v-text-field>
              </v-flex>
              <v-flex xs10>
                <v-text-field label="Password" v-model="password" :append-icon="isShow ? 'visibility' : 'visibility_off'" :append-icon-cb="() => (isShow = !isShow)" :type="isShow ? 'text' : 'password'" counter :rules="[rules.required]" required></v-text-field>
              </v-flex>
                <v-btn @click="Login">LOGIN</v-btn>
                <v-spacer></v-spacer>
                <v-btn @click="SignUp">SIGNUP</v-btn>
            </v-layout>
          </v-container>
          </main>
          </v-app>
        </v-card>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>
// import {HTTP} from '@/router/index'
import { mapGetters } from 'vuex'
export default {
  data() {
    return {
      email: null,
      password: null,
      user_list: [],
      isShow: false,
      // loading: false,
      rules: {
        required: (value) => !!value || 'Required.',
        email: (value) => {
          const pattern = /^(([^<>()\]\\.,;:\s@"]+(\.[^<>()\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
          return pattern.test(value) || 'Invalid e-mail.'
        }
      }
    }
  },

  computed: mapGetters(['getStatusLogin', 'loading']),

  methods: {
    Login: function() {
      this.$store.dispatch('Authenticate', {
        data: {
          user: this.email,
          password: this.password
        }
      }).then(() => {
        localStorage.setItem('auth', this.$store.state.auth)
        localStorage.setItem('token', this.$store.state.token)
        this.$router.push('mhsw')
        this.$store.commit('loading', false)
        this.$store.commit('statusLogin', '')
      })
    },
    SignUp: function() {
      this.$router.push('regist')
    }
  }
}
</script>