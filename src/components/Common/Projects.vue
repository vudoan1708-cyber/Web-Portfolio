<template>
  <div id="project_wrapper" v-for="(app, appKey) in appsProp" :key="appKey">
      <div class="grid_display">
        <ProjectDescription :description="app.description" :githubURL="app.github"
          :liveURL="app.live_url" :programmingLangs="app.programming_languages"
          :projectType="app.project_type" :role:="app.role" :title="app.title" />
        <ProjectDisplay />
      </div>
  </div>
</template>

<script>
import { ref } from 'vue';

import ProjectDescription from '@/components/Reusable/ProjectDescription.vue';
import ProjectDisplay from '@/components/Reusable/ProjectDisplay.vue';

export default {
  name: 'Projects',
  props: {
    apps: {
      type: Array,
    },
  },
  components: {
    ProjectDescription,
    ProjectDisplay,
  },
  setup(props) {
    const appsProp = ref(props.apps);
    return {
      appsProp,
    };
  },
};
</script>

<style lang="scss" scoped>
#project_wrapper {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  background-color: rgba(0, 0, 0, 0.801);
  position: relative;
  width: 100%;
  height: 50%;
  border-radius: 20px;
  margin-bottom: 20px;

  .grid_display {
    display: grid;
    grid-template-columns: 1fr 1fr;
  }
}
</style>
