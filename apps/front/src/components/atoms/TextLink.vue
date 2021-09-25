<template>
    <li><a :href="link.url" target="_blank" rel="noopener">{{ link.text }}</a></li>
</template>

<script lang="ts">
import { defineComponent, PropType } from 'vue'

export interface Link {
  url: string;
  text: string;
}

export default defineComponent({
  name: 'TextLink',
  props: {
    link: {
      type: Object as PropType<Link>,
      required: true
    }
  }
})
</script>

<style scoped lang="scss">
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
