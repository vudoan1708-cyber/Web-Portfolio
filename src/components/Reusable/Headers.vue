<template>
  <div class="headers" v-for="(title, titleKey) in titles.portfolio" :key="titleKey">
    <h2>{{ title }}</h2>
    <hr />
  </div>
</template>

<script>
import { onBeforeMount, reactive } from 'vue';

// JSON
import Headers from '@/components/JSON/headers.json';

export default {
  name: 'Headers',
  setup() {
    const titles = reactive({
      portfolio: [],
      blog: [],
    });

    onBeforeMount(() => {
      Headers.portfolio.forEach((header) => {
        titles.portfolio.push(header);
      });
      Headers.blog.forEach((header) => {
        titles.blog.push(header);
      });
    });

    return {
      titles,
    };
  },
};
</script>

<style lang="scss" scoped>
.headers {
  position: relative;
  top: 15%;
  float: right;
  clear: both;
  width: 30%;
  margin-top: 30px;
  color: white;
  transition: .2s width;

  &:hover {
    width: 20%;
  }
}
</style>
