<template>
  <div id="content">
    <!-- <transition name="fade"> -->
      <Projects :emitter="emitterObj" :projects="projectsProps" :mobile="isMobile" />
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
      type: Object,
    },
    mobile: {
      type: Boolean,
    },
  },
  components: {
    Projects,
  },
  setup(props) {
    // Props
    const emitterObj = ref(props.emitter);
    const projectsProps = ref(props.projects);
    const isMobile = ref(props.mobile);

    watch(() => props.projects, (data) => {
      projectsProps.value = data;
    });

    return {
      emitterObj,
      projectsProps,
      isMobile,
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
