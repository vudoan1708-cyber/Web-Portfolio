<template>
  <div id="portfolio">
    <ProjectTabs />
    <Projects />
  </div>
</template>

<script>
import ProjectTabs from '@/components/ProjectTabs.vue';
import Projects from '@/components/Projects.vue';

export default {
  name: 'Portfolio',
  components: {
    ProjectTabs,
    Projects,
  },
};
</script>

<style scoped lang="scss">
@import '@/sass/Unique/_portfolio.scss';
</style>
