<template>
  <div id="content">
    <Projects :apps="appsProps" />
  </div>
</template>

<script>
import { ref } from 'vue';

import Projects from '@/components/Common/Projects.vue';

export default {
  name: 'Content',
  props: {
    apps: {
      type: Array,
    },
  },
  components: {
    Projects,
  },
  setup(props) {
    const appsProps = ref(props.apps);

    return {
      appsProps,
    };
  },
};
</script>

<style lang="scss" scoped>
#content {
  display: block;
  position: relative;
  top: 25%;
  width: 80%;
  height: 100%;
  padding: 20px;
}
</style>
