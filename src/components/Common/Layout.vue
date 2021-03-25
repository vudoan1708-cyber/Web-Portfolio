<template>
  <div id="layout">
    <Introduction />
    <Content :projects="projects" />
    <Blogs />
    <Resume />
  </div>
</template>

<script>
import { onBeforeMount, ref } from 'vue';

import Introduction from '@/components/Common/Introduction.vue';
import Content from '@/components/Common/Content.vue';
import Blogs from '@/components/Common/Blogs.vue';
import Resume from '@/components/Common/Resume.vue';

// JSON
import Portfolio from '@/components/JSON/portfolio.json';

// Plugin
// import { provideStore } from '@/components/Utils/plugin';

export default {
  name: 'Layout',
  props: {
    emitter: {
      type: Object,
    },
  },
  components: {
    Introduction,
    Content,
    Blogs,
    Resume,
  },
  setup(props) {
    // Apps
    const projects = ref([]);

    // Emitter
    const emitterObj = ref(props.emitter);

    // Code or Design
    emitterObj.value.on('portfolio_view', (view) => {
      projects.value = [];
      if (view === 'Code') {
        Portfolio.coding.apps.forEach((app) => {
          projects.value.push(app);
        });
      } else if (view === 'Design') {
        Portfolio.design.forEach((app) => {
          projects.value.push(app);
        });
      }
    });
    // provideStore('portfolio_view', projects);

    onBeforeMount(() => {
      Portfolio.coding.apps.forEach((app) => {
        projects.value.push(app);
      });
    });

    return {
      projects,
    };
  },
};
</script>

<style lang="scss" scoped>
#layout {
  position: absolute;
  display: block;
  top: 0;
  left: 20%;
  width: 100%;
  height: 100vh;
  z-index: 0;
  background: linear-gradient(to top right, rgb(226, 91, 91) 20%, rgb(255, 140, 99) 50%);
  text-align: center;
  overflow-y: scroll;
  scroll-behavior: smooth;
}
</style>
