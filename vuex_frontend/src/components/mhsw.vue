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
              <v-layout row justify-center>
                <v-dialog v-model="dialog" persistent max-width="500px">
                  <v-btn color="primary" light slot="activator" @click="loadField">Search</v-btn>
                  <v-card>
                    <v-card-title>
                      <span class="headline">Student Search</span>
                    </v-card-title>
                    <v-card-text>
                      <v-container grid-list-md>
                        <v-layout wrap>
                          <v-flex xs12 sm6>
                            <v-select label="Filter" v-bind:items="items" v-model="el_items" required></v-select>
                          </v-flex>

                          <v-flex xs12>
                            <v-text-field label="Enter Text" v-model="value_field" required></v-text-field>
                          </v-flex>
                          </v-flex>
                        </v-layout>
                      </v-container>
                      <small>*indicates required field</small>
                    </v-card-text>
                    <v-card-actions>
                      <v-spacer></v-spacer>
                      <v-btn color="blue darken-1" flat @click.native="close">Close</v-btn>
                      <v-btn color="blue darken-1" flat @click.native="search">Find</v-btn>
                    </v-card-actions>
                  </v-card>
                </v-dialog>
              </v-layout>
            </td>
          </tr>
          <tr></tr>
          <tr v-if="edit_mhsw">
            <template v-for="(mhsw,index) in pages_mhsw" v-bind:pagination.sync="pagination">
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
            <div v-if="isPagination" class="text-xs-center pt-4">
                <v-pagination v-model="page" :length="pages"></v-pagination>
              </div>
              <div v-if="isFilterPagination" class="text-xs-center pt-4">
                <v-pagination v-model="pageFilter" :length="pagesFilter"></v-pagination>
              </div>
          </tr>
          <tr v-else>
            <template v-for="(mhsw,index) in pages_mhsw">
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
            <div v-if="isPagination" class="text-xs-center pt-4">
                <v-pagination v-model="page" :length="pages"></v-pagination>
              </div>
              <div v-if="isFilterPagination" class="text-xs-center pt-4">
                <v-pagination v-model="page" :length="pages"></v-pagination>
              </div>
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
  data() {
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
      edit_mhsw: false,
      dialog: false,
      el_items: null,
      items:[],
      value_field: '',
      pagination: {rowsPerPage:5},
      isPagination: false,
      page:1,
      pages_mhsw: [],
      isFilterPagination: false
    }
  },
  computed: //mapGetters(['pages']),
  {
     pages() {
       return this.pagination.rowsPerPage ? Math.ceil(this.mhsws.length / this.pagination.rowsPerPage) : 0
     }
  },
  watch: {
    page() {
      // console.log(this.page)
      this.pages_mhsw = this.mhsws.slice((this.page-1)*this.pagination.rowsPerPage, this.page*this.pagination.rowsPerPage)
    }
  },
  methods: {
    show_all: function() {
      this.fetchmhsws()
    },
    fetchmhsws: function() {
      HTTP.get('/api/mhsws').then(response => {
        this.mhsws = response.data.data
        this.pages_mhsw = this.mhsws.slice(0, this.pagination.rowsPerPage)
        this.isPagination = true
        this.isFilterPagination = false
        this.page = 1
      })
        .catch(err => {
          console.log(err)
        })
    },
    add: function() {
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
        .catch(function(err) {
          console.log(err.response)
        })
      this.name = ''
      this.address = ''
      this.old = ''
      // console.log('when add edit is: ', this.edit_mhsw)
    },
    edit: function(index) {
      this.edit_mhsw = true
      this.key_id = this.pages_mhsw[index].id
      this.name_in = this.pages_mhsw[index].name
      this.address_in = this.pages_mhsw[index].address
      this.old_in = this.pages_mhsw[index].old
      // console.log('when edit edit is: ', this.name_in, this.address_in, this.old_in)
    },
    remove: function(_id) {
      HTTP.delete('/api/mhsws/' + _id).then(response => {
        console.log('succesfully delete _id: ' + _id)
        this.fetchmhsws()
      })
        .catch(err => {
          console.log(err.response)
        })
      // console.log('when delete edit is: ', this.edit_mhsw)
    },
    update: function(_id, _mhsw) {
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
    valid() { },
    cancel: function() {
      this.fetchmhsws()
      this.edit_mhsw = false
      // console.log('when cancel edit is: ', this.edit_mhsw)
    },
    loadField: function(){
      HTTP.get('/api/mhsws').then(response => {
        let myJson = response.data.data
        let valueJson = Object.values(myJson)[0]
        Object.keys(valueJson).forEach(function(element) {
           this.items.push(element)
        }, this);
      }).catch(err => {
        console.log(err)
      })
    },
    search: function(){
      let field = this.el_items
      let value = this.value_field
      let searchJson = {
        data: {
          field,
          value
        }
      }
      HTTP.post('/api/search_mhsw',searchJson).then(response => {
        console.log(response.data)
        this.mhsws = response.data.data
        this.pages_mhsw = this.mhsws.slice(0,this.pagination.rowsPerPage)
        this.close()
        this.isPagination = false
        this.isFilterPagination = true
      }).catch(err => {
        console.log(err)
      })
    },
    close: function(){
      this.items = []
      this.value_field = ''
      this.dialog = false
    }
  },
  event: {
    click() { }
  },

}
</script>