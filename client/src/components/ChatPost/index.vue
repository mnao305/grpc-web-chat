<template>
  <div>
    <input type="text" v-model="state.text">
    <button @click="post">投稿</button>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive } from 'vue'
import { client } from '@/messenger'
import { MessageRequest } from '@/messenger/messenger_pb'

export default defineComponent({
  setup () {
    const state = reactive({
      text: ''
    })
    const post = () => {
      const req = new MessageRequest()
      req.setMessage(state.text)
      client.createMessage(req, null, res => console.log(res))
      state.text = ''
    }
    return { state, post }
  }
})
</script>
