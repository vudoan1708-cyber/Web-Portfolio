<template>
  <div id="content">
    <!-- <transition name="fade"> -->
      <Projects :emitter="emitterObj" :projects="projectsProps" />
    <!-- </transition> -->
  </div>
</template>

<script>
import { ref, watch } from 'vue';

import Projects from '@/components/Common/Projects.vue';

export default {
  name: 'Content',
  props: {
    emitter: {
      type: Object,
    },
    projects: {
      type: Array,
    },
  },
  components: {
    Projects,
  },
  setup(props) {
    // Props
    const emitterObj = ref(props.emitter);
    const projectsProps = ref(props.projects);

    watch(() => props.projects, (data) => {
      projectsProps.value = data;
    });

    return {
      emitterObj,
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
