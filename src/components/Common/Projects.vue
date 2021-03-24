<template>
  <div class="project_wrapper" v-for="(app, appKey) in appsProp" :key="appKey">
      <div class="grid_display" @click="openWindow(app.live_url)">
        <ProjectDescription :description="app.description" :githubURL="app.github"
          :liveURL="app.live_url" :programmingLangs="app.programming_languages"
          :projectType="app.project_type" :role="app.role" :title="app.title"
          :videos="app.extra_links.videos" :presentations="app.extra_links.presentations"
          :orgName="app.organisation.name" :orgExtraLink="app.organisation.extra_links" />
        <ProjectDisplay :image="app.images.thumbnail" />
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

    function openWindow(url) {
      // eslint-disable-next-line no-unused-expressions
      url !== '#' ? window.open(url) : undefined;
    }
    return {
      appsProp,
      openWindow,
    };
  },
};
</script>

<style lang="scss" scoped>
.project_wrapper {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  position: relative;
  left: 0;
  width: 100%;
  height: 50%;
  border-radius: 20px;
  margin-bottom: 20px;
  cursor: pointer;
  transition: .2s all;

  &:hover {
    left: 20px;
    box-shadow: 2px 2px 15px black;
  }

  .grid_display {
    display: grid;
    grid-template-columns: 1fr 1fr;
  }
}
</style>
