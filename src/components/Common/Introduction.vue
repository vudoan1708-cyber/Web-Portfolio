<template>
  <div id="intro">
    <h1>{{ Intro.name }}</h1>
    <div class="intro_contents" v-for="(content, contentKey) in Intro.content" :key="contentKey">
      <p>{{ content }}</p>
    </div>
  </div>
</template>

<script>
// JSON
import Intro from '@/components/JSON/introduction.json';

export default {
  name: 'Introduction',
  setup() {
    return {
      Intro,
    };
  },
};
</script>

<style lang="scss" scoped>
#intro {
  display: block;
  position: relative;
  top: 15%;
  width: 80%;
  padding: 20px;
  color: snow;

  .intro_contents {
    text-align: left;
    margin: 10px;
    padding: 5px;
  }
}
</style>
