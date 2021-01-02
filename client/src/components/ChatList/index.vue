<template>
  <div>
    <div v-for="(a, i) of state.list" :key="`chat-${i}`" class="chat">
      <span class="name">{{ a.name }}</span>
      <br />
      <span class="text">{{ a.text }}</span>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive } from 'vue'
import { Empty } from 'google-protobuf/google/protobuf/empty_pb'
import { client } from '@/messenger'
import { MessageResponse } from '@/messenger/messenger_pb'

export type chat = {
  name: string;
  text: string;
  date: string;
};

export default defineComponent({
  setup () {
    const state = reactive<{ list: chat[] }>({
      list: []
    })
    const stream = client.getMessages(new Empty())
    stream.on('data', m => {
      const c = m as MessageResponse
      state.list.push({
        text: c.getMessage(),
        name: c.getName(),
        date: c.getDate()
      })
    })
    return { state }
  }
})
</script>

<style lang="scss" scoped>
.chat {
  border: 1px solid #ccc;
  padding: 5px;
  margin: 5px 0;
}
</style>
