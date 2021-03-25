<template>
  <div id="content">
    <!-- <transition name="fade"> -->
      <Projects :projects="projectsProps" />
    <!-- </transition> -->
  </div>
</template>

<script>
import { ref, watch } from 'vue';

import Projects from '@/components/Common/Projects.vue';

export default {
  name: 'Content',
  props: {
    projects: {
      type: Array,
    },
  },
  components: {
    Projects,
  },
  setup(props) {
    const projectsProps = ref(props.projects);

    watch(() => props.projects, (data) => {
      projectsProps.value = data;
    });

    return {
      projectsProps,
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
