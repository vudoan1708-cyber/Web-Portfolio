<template>
  <div id="layout">
    <Introduction />
    <Content :apps="apps" />
  </div>
</template>

<script>
import { onBeforeMount, ref } from 'vue';

import Introduction from '@/components/Common/Introduction.vue';
import Content from '@/components/Common/Content.vue';

// JSON
import Portfolio from '@/components/JSON/portfolio.json';

export default {
  name: 'Layout',
  components: {
    Introduction,
    Content,
  },
  setup() {
    // Apps
    const apps = ref([]);

    onBeforeMount(() => {
      Portfolio.coding.apps.forEach((app) => {
        apps.value.push(app);
      });
    });

    return {
      apps,
    };
  },
};
</script>

<style lang="scss" scoped>
#layout {
  position: absolute;
  top: 0;
  left: 20%;
  width: 100%;
  height: 100vh;
  z-index: 0;
  background-color: rgb(255, 140, 99);
  text-align: center;
  overflow-y: scroll;
}
</style>
