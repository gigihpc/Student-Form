<template>
  <v-form v-model="valid">
    <v-text-field label="Name" v-model="name" ></v-text-field>
    <v-text-field label="Address" v-model="address"></v-text-field>
    <v-text-field label="Old" v-model="old"></v-text-field>

    <v-btn @click="add">Add</v-btn>
    <v-btn @click="edit">Edit</v-btn>
    <v-btn @click="remove">Delete</v-btn>
    <v-btn @click="showAll">Show All</v-btn>
  </v-form>
</template>

<script>
import {HTTP} from '@/router/index'
export default {
  $validates: true,
  data () {
    return {
      name: '',
      address: '',
      old: '',
      result: [],
      error: []
    }
  },
  methods: {
    showAll () {
      HTTP.get('/api/mhsws').then(response => {
        this.result = response.data
        this.result.data.forEach(function (element) {
          console.log(element.id + ';' + element.name + ';' + element.address + ';' + element.old)
        }, this)
      })
      .catch(e => {
        console.log(e)
      })
    },
    add () {},
    edit () {},
    remove () {}
  }
}
</script>
