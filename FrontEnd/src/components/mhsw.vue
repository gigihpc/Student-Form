<template>
  <v-form>
    <title>VueJs</title>
    <div class="container">
      <h1>VueJs Tutorial</h1>

      <table class="table">
        <thead>
          <tr>
            <th>Name</th>
            <th>Address</th>
            <th>Old</th>
            <th>Action</th>
          </tr>
          <tr>
            <td>
              <v-text-field v-model="name"></v-text-field>
            </td>
            <td>
              <v-text-field v-model="address"></v-text-field>
            </td>
            <td>
              <v-text-field v-model="old"></v-text-field>
            </td>
            <td>
              <v-btn v-on:click.stop.prevent="add">Add</v-btn>
              <v-btn @click="show_all">SHOW ALL</v-btn>
            </td>
          </tr>
          <tr></tr>
          <tr v-if="edit_mhsw">
            <template v-for="(mhsw,index) in mhsws">
              <tr :key="mhsw.id">
                <td v-if="mhsw.id==key_id"><input v-model="name_in"></input>
                </td>
                <td v-else>{{mhsw.name}}</td>
                <td v-if="mhsw.id==key_id"><input v-model="address_in"></input>
                </td>
                <td v-else>{{mhsw.address}}</td>
                <td v-if="mhsw.id==key_id"><input v-model="old_in"></input>
                </td>
                <td v-else>{{mhsw.old}}</td>

                <td>
                  <v-btn v-if="mhsw.id==key_id" @click="update(mhsw.id,mhsw)">Update</v-btn>
                  <v-btn v-else @click="edit(index)">Edit</v-btn>
                  <v-btn v-if="mhsw.id==key_id" @click="cancel">Cancel</v-btn>
                  <v-btn v-else @click="remove(mhsw.id)">Delete</v-btn>
                </td>
              </tr>
            </template>
          </tr>
          <tr v-else>
            <template v-for="(mhsw,index) in mhsws">
              <tr :key="mhsw.id">
                <td>{{mhsw.name}}</td>
                <td>{{mhsw.address}}</td>
                <td>{{mhsw.old}}</td>
                <td>
                  <v-btn @click="edit(index)">Edit</v-btn>
                  <v-btn @click="remove(mhsw.id)">Delete</v-btn>
                </td>
              </tr>
            </template>
          </tr>
        </thead>
      </table>
    </div>
  </v-form>
</template>

<script>
import { HTTP } from '@/router/index'
// import textinput from './text_input'
// import axios from 'axios'
export default {
  // components: {
  //   'text-input': textinput
  // },
  data () {
    return {
      name: '',
      address: '',
      old: '',
      name_in: '',
      address_in: '',
      old_in: '',
      mhsw: [],
      mhsws: [],
      key_id: '',
      edit_mhsw: false
    }
  },
  ready: function () {
    this.fetchmhsws()
  },
  methods: {
    show_all: function () {
      this.fetchmhsws()
      // console.log('when showall edit is: ', this.edit_mhsw)
    },
    fetchmhsws: function () {
      HTTP.get('/api/mhsws').then(response => {
        this.mhsws = response.data.data
      })
        .catch(err => {
          console.log(err)
        })
    },
    add: function () {
      var name = this.name
      var address = this.address
      var old = this.old
      var mhswJson = {
        data: {
          name,
          address,
          old
        }
      }
      console.log(mhswJson)
      HTTP.post('/api/mhsws', mhswJson)
        .then(response => {
          console.log('succesfully create mahasiswa')
          this.fetchmhsws()
        })
        .catch(function (err) {
          console.log(err.response)
        })
      this.name = ''
      this.address = ''
      this.old = ''
      // console.log('when add edit is: ', this.edit_mhsw)
    },
    edit: function (index) {
      this.edit_mhsw = true
      this.key_id = this.mhsws[index].id
      this.name_in = this.mhsws[index].name
      this.address_in = this.mhsws[index].address
      this.old_in = this.mhsws[index].old
      // console.log('when edit edit is: ', this.name_in, this.address_in, this.old_in)
    },
    remove: function (_id) {
      HTTP.delete('/api/mhsws/' + _id).then(response => {
        console.log('succesfully delete _id: ' + _id)
        this.fetchmhsws()
      })
        .catch(err => {
          console.log(err.response)
        })
      // console.log('when delete edit is: ', this.edit_mhsw)
    },
    update: function (_id, _mhsw) {
      // console.log('in: ', this.name_in, this.address_in, this.old_in)
      var name = this.name_in
      var address = this.address_in
      var old = this.old_in
      // console.log('curr-mhsw: ', name, address, old)
      var updateJson = {
        data: {
          name,
          address,
          old
        }
      }
      // console.log('json: ', updateJson)
      HTTP.put('/api/mhsws/' + _id, updateJson).then(response => {
        console.log('success update _id:' + _id)
        this.fetchmhsws()
      })
        .catch(err => {
          console.log(err.response)
        })
      this.edit_mhsw = false
      // console.log('when update edit is: ', this.edit_mhsw)
    },
    valid () { },
    cancel: function () {
      this.fetchmhsws()
      this.edit_mhsw = false
      // console.log('when cancel edit is: ', this.edit_mhsw)
    }
  },
  event: {
    click () { }
  }

}
</script>