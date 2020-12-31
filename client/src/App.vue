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

<style lang="scss">
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
