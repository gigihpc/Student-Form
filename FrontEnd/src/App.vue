<template>
  <v-app light>
    <v-navigation-drawer v-if="!$route.meta.disabledAction" persistent :mini-variant="miniVariant" :clipped="clipped" v-model="drawer" enable-resize-watcher>
      <v-list>
        <v-list-tile value="true" v-for="(item, i) in items" :key="i">
          <v-list-tile-action>
            <v-icon light v-html="item.icon"></v-icon>
          </v-list-tile-action>
          <v-list-tile-content>
            <v-list-tile-title v-text="item.title"></v-list-tile-title>
          </v-list-tile-content>
        </v-list-tile>
      </v-list>
    </v-navigation-drawer>
    <v-toolbar fixed>
      <v-toolbar-side-icon v-if="!$route.meta.disabledAction" @click.stop="drawer = !drawer" light></v-toolbar-side-icon>
      <!-- <v-btn icon light @click.stop="miniVariant = !miniVariant">
        <v-icon v-html="miniVariant ? 'chevron_right' : 'chevron_left'"></v-icon>
      </v-btn> -->
      <!-- <v-btn icon light @click.stop="clipped = !clipped">
        <v-icon></v-icon>
      </v-btn>
      <v-btn icon light @click.stop="fixed = !fixed">
        <v-icon></v-icon>
      </v-btn> -->
      <v-layout justify-center>
        <v-toolbar-title v-text="title"></v-toolbar-title>
      </v-layout>
      <v-spacer></v-spacer>
      <v-btn v-if="!$route.meta.disabledAction" icon light @click.stop="rightDrawer = !rightDrawer">
        <v-icon>settings</v-icon>
      </v-btn>
    </v-toolbar>
    <main>
      <v-container fluid>
        <v-slide-y-transition mode="out-in">
          <router-view></router-view>
        </v-slide-y-transition>
      </v-container>
    </main>
    <v-navigation-drawer temporary :right="right" v-model="rightDrawer">
      <v-list>
        <!-- <v-list-tile @click="right = !right">
            <v-list-tile-action>
              <v-icon light>compare_arrows</v-icon>
            </v-list-tile-action>
            <v-list-tile-title>Switch drawer (click me)</v-list-tile-title>
          </v-list-tile> -->
        <v-list-tile @click="logOut">
          <v-list-tile-action>
            <v-icon light>lock_open</v-icon>
          </v-list-tile-action>
          <v-list-tile-content>
            <v-list-tile-title>Log Out</v-list-tile-title>
          </v-list-tile-content>
        </v-list-tile>
      </v-list>
    </v-navigation-drawer>
    <v-footer :fixed="fixed">
      <span>&copy; 2017</span>
    </v-footer>
  </v-app>
</template>

<script>
export default {
  data () {
    return {
      clipped: false,
      drawer: true,
      fixed: false,
      items: [
        { icon: 'home', title: 'Home' },
        { icon: 'account_circle', title: 'User' },
        { icon: 'business', title: 'Management' }
      ],
      miniVariant: false,
      right: true,
      rightDrawer: false,
      title: 'Welcome To My Page'
      // isLogin: !!localStorage.getItem('auth')
    }
  },
  methods: {
    logOut: function () {
      localStorage.removeItem('auth')
      this.$router.push('login')
    }
  }
}
</script>

<style lang="stylus">
  @import './stylus/main'
</style>
