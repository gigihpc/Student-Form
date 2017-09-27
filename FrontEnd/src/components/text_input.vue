<template>
      <!-- <td><input v-bind:value="value" @input="updateValue($event.target.value)"></input></td> -->
      <td><input v-model="value"></input></td>
       <!-- <td><v-text-field v-model="value"></v-text-field></td>  -->
</template>

<script>

export default {
  props: {
    value: {
      type: String,
      default: ''
    }
  },
  methods: {
    updateValue: function (value) {
      // console.log(value)
      this.$emit('input', value)
    }
  }
}
</script>