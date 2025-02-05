<template>
  <div
    v-for="(title, titleKey) in titles.portfolio"
    class="headers"
    :class="selected === title.toLowerCase() ? 'selected' : 'not_selected'"
    :key="titleKey"
    @click="changeView(title)">
    <a href="#content">
      <h2>{{ title }}</h2>
    </a>
    <hr />
  </div>
</template>

<script>
import { onBeforeMount, reactive, ref } from 'vue';

// JSON
import Headers from '@/components/JSON/headers.json';

// Plugin
import { useStore } from '@/components/Utils/plugin';

export default {
  name: 'Headers',
  setup() {
    const titles = reactive({
      portfolio: [],
      blog: [],
    });
    const selected = ref('code');

    // Emitter
    const emitter = ref(useStore('emitter'));

    // Code or Design
    function changeView(title) {
      // New selection
      selected.value = title.toLowerCase();
      if (!emitter.value.err) {
        emitter.value.store.emit('portfolio_view', title);
      }
    }

    onBeforeMount(() => {
      Headers.portfolio.forEach((header) => {
        titles.portfolio.push(header);
      });
      Headers.blog.forEach((header) => {
        titles.blog.push(header);
      });
    });

    return {
      titles,
      selected,
      changeView,
    };
  },
};
</script>

<style lang="scss" scoped>
.headers {
  position: relative;
  top: 15%;
  float: right;
  clear: both;
  width: 30%;
  margin-top: 30px;
  color: white;
  transition: .2s width;
  cursor: pointer;

  a {
    text-decoration: none;
    color: white !important;
  }

  &:hover {
    width: 20%;
  }

  &.selected {
    width: 20%;
  }
}

/* Responsive Text */
@media screen and (max-width: 770px) {
  .headers {
    font-size: 10px;
  }
}
@media screen and (max-width: 600px) {
  .headers {
    font-size: 10px;
  }
}
@media screen and (max-width: 460px) {
  .headers {
    font-size: 6px;
  }
}
</style>
