<template>
  <div id="content">
    <transition name="fade">
      <Projects :apps="appsProps" />
    </transition>
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
  top: 15%;
  left: 0;
  width: 80%;
  height: auto;
  min-height: 100% !important;
  padding: 20px;
}
</style>
