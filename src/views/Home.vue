<template>
  <div class="home">
    <Navigation />
    <Layout />
    <ContactDetail />
  </div>
</template>

<script>
import Navigation from '@/components/Common/Navigation.vue';
import Layout from '@/components/Common/Layout.vue';
import ContactDetail from '@/components/Common/ContactDetail.vue';

export default {
  name: 'Home',
  components: {
    Navigation,
    Layout,
    ContactDetail,
  },
};
</script>

<style scoped lang="scss">
#home {
  position: relative;
  height: 100%;
}
</style>
