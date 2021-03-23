<template>
  <div id="descriptions">
    <h2>{{ projectDetails.title }}</h2>
  </div>
</template>

<script>
import { reactive } from 'vue';

export default {
  name: 'ProjectDescription',
  props: {
    description: {
      type: String,
    },
    githubURL: {
      type: String,
    },
    liveURL: {
      type: String,
    },
    programmingLangs: {
      type: Array,
    },
    projectType: {
      type: String,
    },
    role: {
      type: String,
    },
    title: {
      type: String,
    },
  },
  setup(props) {
    const projectDetails = reactive({
      description: props.description,
      githubURL: props.githubURL,
      liveURL: props.liveURL,
      programmingLangs: props.programmingLangs,
      projectType: props.projectType,
      role: props.role,
      title: props.title,
    });

    return {
      projectDetails,
    };
  },
};
</script>

<style lang="scss" scoped>

</style>
