<template>
  <div class="project_wrapper" v-for="(project, projectKey) in projectsProp" :key="projectKey">
    <div class="grid_display" @click="openWindow(project)">
      <ProjectDescription :description="project.description" :githubURL="project.github"
        :liveURL="project.live_url" :technologies="project.technologies"
        :projectType="project.project_type" :role="project.role" :title="project.title"
        :videos="project.extra_links.videos" :presentations="project.extra_links.presentations"
        :orgName="project.organisation.name" :orgExtraLink="project.organisation.extra_links"
        :orgNDA="project.organisation.nda" :timeline="project.timeline"
        :mobile="isMobile" />
      <ProjectDisplay :image="project.images.thumbnail" />
    </div>
  </div>
</template>

<script>
import { ref, watch } from 'vue';

// Reusable Components
import ProjectDescription from '@/components/Reusable/ProjectDescription.vue';
import ProjectDisplay from '@/components/Reusable/ProjectDisplay.vue';

export default {
  name: 'Projects',
  props: {
    emitter: {
      type: Object,
    },
    projects: {
      type: Array,
    },
    mobile: {
      type: Boolean,
    },
  },
  components: {
    ProjectDescription,
    ProjectDisplay,
  },
  setup(props) {
    // Props
    const emitterObj = ref(props.emitter);
    const projectsProp = ref(props.projects);
    const isMobile = ref(props.mobile);

    function openWindow(project) {
      // Emit an event of 'project_card' to show further detail of a clicked card
      emitterObj.value.emit('project_card', project);
    }

    watch(() => (props.projects), (data) => {
      projectsProp.value = data;
    });

    return {
      projectsProp,
      isMobile,
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
