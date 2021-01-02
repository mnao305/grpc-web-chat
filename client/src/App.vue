<template>
  <ChatPost />
  <ChatList />
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import ChatPost from './components/ChatPost/index.vue'
import ChatList from './components/ChatList/index.vue'

export default defineComponent({
  name: 'App',
  components: {
    ChatPost,
    ChatList
  }
})
</script>
