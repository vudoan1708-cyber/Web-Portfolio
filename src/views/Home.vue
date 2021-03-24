<template>
  <div class="home">
    <Headers />
    <Navigation />
    <Layout />
    <ContactDetail />
  </div>
</template>

<script>
// Common
import Navigation from '@/components/Common/Navigation.vue';
import Layout from '@/components/Common/Layout.vue';
import ContactDetail from '@/components/Common/ContactDetail.vue';

// Reusable
import Headers from '@/components/Reusable/Headers.vue';

export default {
  name: 'Home',
  components: {
    Navigation,
    Layout,
    ContactDetail,
    Headers,
  },
};
</script>

<style scoped lang="scss">
#home {
  position: relative;
  height: 100%;
}
</style>
