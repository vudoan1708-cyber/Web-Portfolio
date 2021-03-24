<template>
  <div class="display">
    <img :src="require(`@/assets/${images}.png`)" />
  </div>
</template>

<script>
import { ref } from 'vue';

// Images
// import Cerberus from '@/assets/cerberus.png';

export default {
  name: 'ProjectDisplay',
  props: {
    image: {
      type: String,
    },
  },
  setup(props) {
    const images = ref((props.image));
    return {
      images,
    };
  },
};
</script>

<style lang="scss" scoped>
.display {
  position: absolute;
  top: 0;
  right: 0;
  background-color: black;
  width: 50%;
  height: 100%;
  border-top-right-radius: 20px;
  border-bottom-right-radius: 20px;

  img {
    position: relative;
    top: 50%;
    transform: translateY(-50%);
    width: 100%;
  }
}
</style>
