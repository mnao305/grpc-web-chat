<template>
  <div>
    <input type="text" v-model="state.name" placeholder="name">
    <input type="text" v-model="state.text" placeholder="text">
    <button @click="post">post</button>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive } from 'vue'
import { client } from '@/messenger'
import { MessageRequest } from '@/messenger/messenger_pb'

export default defineComponent({
  setup () {
    const state = reactive({
      text: '',
      name: ''
    })
    const post = () => {
      if (state.text === '') {
        alert('Please enter the text')
        return
      }
      const req = new MessageRequest()
      req.setMessage(state.text)
      req.setName(state.name === '' ? 'no name' : state.name)
      client.createMessage(req, null)
      state.text = ''
    }
    return { state, post }
  }
})
</script>
