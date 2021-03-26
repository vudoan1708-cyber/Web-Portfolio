<template>
  <div class="home">
    <Navigation />
    <Layout :emitter="emitter" />
    <ContactDetail :emitter="emitter" />
  </div>
</template>

<script>
import { getCurrentInstance } from 'vue';

// Common
import Navigation from '@/components/Common/Navigation.vue';
import Layout from '@/components/Common/Layout.vue';
import ContactDetail from '@/components/Common/ContactDetail.vue';

// Plugin
import { provideStore } from '@/components/Utils/plugin';

export default {
  name: 'Home',
  components: {
    Navigation,
    Layout,
    ContactDetail,
  },
  setup() {
    // instantiate the app's current instance to get global properties
    // registered in the main.js file
    const app = getCurrentInstance();
    const emitter = app.appContext.config.globalProperties.$emitter;

    provideStore('emitter', emitter);
    return {
      emitter,
    };
  },
};
</script>

<style scoped lang="scss">
#home {
  position: relative;
  height: 100%;
}
</style>
