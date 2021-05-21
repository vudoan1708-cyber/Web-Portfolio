<template>
  <div class="layout" v-if="emailDetail === ''">
    <Introduction />
    <Content :emitter="emitterObj" :projects="projects" />
    <Blogs />
    <Resume />
  </div>

  <div class="layout" v-else-if="emailDetail !== ''">
    <Message v-if="emailDetail !== 'vudoan1708@gmail.com'" :emailDetail="emailDetail" />
    <Email />
  </div>

  <div class="layout" v-if="furtherDetail !== null">
    <FurtherDetails :emitter="emitterObj" :furtherDetail="furtherDetail" />
  </div>
</template>

<script>
import { onBeforeMount, ref } from 'vue';

// Common
import Introduction from '@/components/Common/Introduction.vue';
import Content from '@/components/Common/Content.vue';
import Blogs from '@/components/Common/Blogs.vue';
import Resume from '@/components/Common/Resume.vue';
import Email from '@/components/Common/Email.vue';

// Reusable
import Message from '@/components/Reusable/Message.vue';
import FurtherDetails from '@/components/Reusable/FurtherDetails.vue';

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
    Email,
    Message,
    FurtherDetails,
  },
  setup(props) {
    // Props
    const emitterObj = ref(props.emitter);

    // Apps
    const projects = ref([]);

    // Email Section
    const emailDetail = ref('');

    // Further Detail Section
    const furtherDetail = ref(null);

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

    // Open up Email Screen
    emitterObj.value.on('email_screen', (email) => {
      emailDetail.value = email;
    });

    // Open up Further Detail Screen
    emitterObj.value.on('project_card', (detail) => {
      furtherDetail.value = detail;
    });

    onBeforeMount(() => {
      Portfolio.coding.apps.forEach((app) => {
        projects.value.push(app);
      });
    });

    return {
      emitterObj,
      projects,
      emailDetail,
      furtherDetail,
    };
  },
};
</script>

<style lang="scss" scoped>
.layout {
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
