<template>
  <div class="home">
    <Navigation />
    <ImageGallery />
    <Portfolio />
  </div>
</template>

<script>
import Navigation from '@/components/Navigation.vue';
import ImageGallery from '@/components/ImageGallery.vue';
import Portfolio from '@/components/Portfolio.vue';

export default {
  name: 'Home',
  components: {
    Navigation,
    ImageGallery,
    Portfolio,
  },
};
</script>
