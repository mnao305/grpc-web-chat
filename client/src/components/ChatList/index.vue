<template>
  <div>
    <ul>
      <li  v-for="(a, i) of state.list"  :key="`chat-${i}`">{{a}} - {{i}}</li>
    </ul>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive } from 'vue'
import { Empty } from 'google-protobuf/google/protobuf/empty_pb'
import { client } from '@/messenger'

export default defineComponent({
  setup () {
    const state = reactive<{list: string[]}>({
      list: []
    })
    const stream = client.getMessages(new Empty())
    stream.on('data', (m: any) => {
      state.list.push(m.getMessage())
    })
    return { state }
  }
})
</script>
