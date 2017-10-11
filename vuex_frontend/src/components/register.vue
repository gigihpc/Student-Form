<template>
  <v-container fluid>
    <v-layout row wrap justify-center>
      <v-flex xs12 sm4>
        <v-card>
          <v-toolbar color="teal">
            <v-toolbar-title>Create User</v-toolbar-title>
          </v-toolbar>
          <v-container fluid>
            <v-layout row wrap>
              <v-flex xs12>
                <v-alert warning :value="isError" transition="scale-transition">{{err}}</v-alert>
              </v-flex>
              <v-flex xs12>
                <v-text-field label="Full Name" v-model="name" required></v-text-field>
              </v-flex>
              <v-flex xs12>
                <v-text-field label="Email" v-model="email" :rules="[rules.required, rules.email]" placeholder="michel@gmail.com" required></v-text-field>
              </v-flex>
              <v-text-field label="Password" v-model="password" required :append-icon="isShow ? 'visibility' : 'visibility_off'" :append-icon-cb="() => (isShow = !isShow)" :type="isShow ? 'text' : 'password'" counter></v-text-field>
              <v-flex xs12>
                <v-select autocomplete label="Country" placeholder="Select..." :items="countries" v-model="country" ref="country" required></v-select>
              </v-flex>
              <v-btn v-on:click="submit">SUBMIT</v-btn>
              <v-spacer></v-spacer>
              <v-btn v-on:click="cancel">CANCEL</v-btn>
            </v-layout>
          </v-container>
        </v-card>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>
import { HTTP } from '@/router/index'
export default {
  data() {
    return {
      countries: ['Afghanistan', 'Albania', 'Algeria', 'Andorra', 'Angola', 'Anguilla', 'Antigua &amp; Barbuda', 'Argentina', 'Armenia', 'Aruba', 'Australia', 'Austria', 'Azerbaijan', 'Bahamas', 'Bahrain', 'Bangladesh', 'Barbados', 'Belarus', 'Belgium', 'Belize', 'Benin', 'Bermuda', 'Bhutan', 'Bolivia', 'Bosnia &amp; Herzegovina', 'Botswana', 'Brazil', 'British Virgin Islands', 'Brunei', 'Bulgaria', 'Burkina Faso', 'Burundi', 'Cambodia', 'Cameroon', 'Cape Verde', 'Cayman Islands', 'Chad', 'Chile', 'China', 'Colombia', 'Congo', 'Cook Islands', 'Costa Rica', 'Cote D Ivoire', 'Croatia', 'Cruise Ship', 'Cuba', 'Cyprus', 'Czech Republic', 'Denmark', 'Djibouti', 'Dominica', 'Dominican Republic', 'Ecuador', 'Egypt', 'El Salvador', 'Equatorial Guinea', 'Estonia', 'Ethiopia', 'Falkland Islands', 'Faroe Islands', 'Fiji', 'Finland', 'France', 'French Polynesia', 'French West Indies', 'Gabon', 'Gambia', 'Georgia', 'Germany', 'Ghana', 'Gibraltar', 'Greece', 'Greenland', 'Grenada', 'Guam', 'Guatemala', 'Guernsey', 'Guinea', 'Guinea Bissau', 'Guyana', 'Haiti', 'Honduras', 'Hong Kong', 'Hungary', 'Iceland', 'India', 'Indonesia', 'Iran', 'Iraq', 'Ireland', 'Isle of Man', 'Israel', 'Italy', 'Jamaica', 'Japan', 'Jersey', 'Jordan', 'Kazakhstan', 'Kenya', 'Kuwait', 'Kyrgyz Republic', 'Laos', 'Latvia', 'Lebanon', 'Lesotho', 'Liberia', 'Libya', 'Liechtenstein', 'Lithuania', 'Luxembourg', 'Macau', 'Macedonia', 'Madagascar', 'Malawi', 'Malaysia', 'Maldives', 'Mali', 'Malta', 'Mauritania', 'Mauritius', 'Mexico', 'Moldova', 'Monaco', 'Mongolia', 'Montenegro', 'Montserrat', 'Morocco', 'Mozambique', 'Namibia', 'Nepal', 'Netherlands', 'Netherlands Antilles', 'New Caledonia', 'New Zealand', 'Nicaragua', 'Niger', 'Nigeria', 'Norway', 'Oman', 'Pakistan', 'Palestine', 'Panama', 'Papua New Guinea', 'Paraguay', 'Peru', 'Philippines', 'Poland', 'Portugal', 'Puerto Rico', 'Qatar', 'Reunion', 'Romania', 'Russia', 'Rwanda', 'Saint Pierre &amp; Miquelon', 'Samoa', 'San Marino', 'Satellite', 'Saudi Arabia', 'Senegal', 'Serbia', 'Seychelles', 'Sierra Leone', 'Singapore', 'Slovakia', 'Slovenia', 'South Africa', 'South Korea', 'Spain', 'Sri Lanka', 'St Kitts &amp; Nevis', 'St Lucia', 'St Vincent', 'St. Lucia', 'Sudan', 'Suriname', 'Swaziland', 'Sweden', 'Switzerland', 'Syria', 'Taiwan', 'Tajikistan', 'Tanzania', 'Thailand', 'Timor Leste', 'Togo', 'Tonga', 'Trinidad &amp; Tobago', 'Tunisia', 'Turkey', 'Turkmenistan', 'Turks &amp; Caicos', 'Uganda', 'Ukraine', 'United Arab Emirates', 'United Kingdom', 'United States', 'Uruguay', 'Uzbekistan', 'Venezuela', 'Vietnam', 'Virgin Islands (US)', 'Yemen', 'Zambia', 'Zimbabwe'],
      name: null,
      country: null,
      email: null,
      password: null,
      isShow: false,
      isError: false,
      err: '',
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
    submit: function() {
      // console.log('submit')
      this.err = ''
      let name = this.name
      let email = this.email
      let password = this.password
      let country = this.country
      var userJson = {
        data: {
          name,
          email,
          password,
          country
        }
      }
      // console.log(name, email, password, country)
      HTTP.post('/api/user', userJson)
        .then(response => {
          if (response.data.data.id) {
            console.log('success create user with id: ' + response.data.data.id)
            this.rule()
            this.$router.push('login')
          } else {
            // let err
            this.isError = true
            if (response.data.data.name === '') {
              this.err = 'Name '
            }
            if (response.data.data.email === '') {
              this.err = this.err + 'Email '
            }
            if (response.data.data.password === '') {
              this.err = this.err + 'Password '
            }
            if (response.data.data.country === '') {
              this.err = this.err + 'Country'
            }
            this.err = this.err + ' can`t empty !!!'
            // console.log('error is: ' + this.err)
            console.log('User can`t saved')
          }
          // console.log('response from db: ' + response.data.data.name, response.data.data.email, response.data.data.password, response.data.data.country)
        })
        .catch(error => {
          console.log(error.response)
        })
    },
    cancel: function() {
      this.name = ''
      this.email = ''
      this.password = ''
      this.country = ''
      this.$router.push('login')
    },
    rule: function() {
      this.name = ''
      this.email = ''
      this.password = ''
      this.country = ''
    }
  }
}
</script>
