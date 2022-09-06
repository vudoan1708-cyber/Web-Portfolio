<template>
  <div class="display">
    <img :src="require(`@/assets/imgs/${images}.png`)" />
  </div>
</template>

<script>
import { ref, watch } from 'vue';

export default {
  name: 'ProjectDisplay',
  props: {
    image: {
      type: String,
    },
  },
  setup(props) {
    const images = ref((props.image));

    watch(() => (props.image), (data) => {
      images.value = data;
    });

    return {
      images,
    };
  },
};
</script>

<style lang="scss" scoped>
.display {
  position: inherit;
  top: 0;
  right: 0;
  background-color: black;
  width: 100%;
  height: 100%;
  border-top-right-radius: 20px;
  border-bottom-right-radius: 20px;
  overflow: hidden;

  img {
    position: relative;
    top: 50%;
    transform: translateY(-50%);
    width: 100%;
  }
}
</style>
