<template>
  <div id="admin">
    <form @submit.prevent>
      <label v-for="(input, inputKey) in randomInputs" :key="inputKey">
        <span>{{ input.order }}</span>
        <input
          type="text"
          minlength="1"
          ref="labelRefs"
          v-model="inputtedCode[inputKey]" />
      </label>
    </form>
  </div>
</template>

<script>
import {
  onBeforeMount, onMounted, ref,
} from 'vue';

import ADMIN_CODE from '@/components/Utils/constants';

export default {
  name: 'Admin',
  setup() {
    const codes = ref(ADMIN_CODE);
    const randomInputs = ref([]);
    const inputtedCode = ref([]);
    // DOM Ref
    const labelRefs = ref([]);

    const randomPick = (count) => {
      if (count > 4) return;
      const randomIdx = Math.floor(Math.random() * codes.value.length);
      // Add to the array
      randomInputs.value = [
        ...randomInputs.value, codes.value[randomIdx],
      ];
      // Remove the item at that picked index from the codes
      codes.value.splice(randomIdx, 1);

      randomPick(count + 1);
    };

    onBeforeMount(() => {
      randomPick(1);
    });

    onMounted(() => {
      console.log(labelRefs.value);
      // labelRefs.value[0].focus();
    });

    return {
      randomInputs,
      inputtedCode,
      labelRefs,
    };
  },
};
</script>

<style scoped lang="scss">
  @import '@/sass/Shared/main';

  #admin {
    position: relative;
    width: 100%;
    height: 100vh;
    display: flex;
    justify-content: center;
    align-items: center;
    background: linear-gradient(to top right, rgb(226, 91, 91) 20%, rgb(255, 140, 99) 50%);
    text-align: center;
    scroll-behavior: smooth;

    label {
      font-size: 20px;

      span {
        background-color: white;
        margin-bottom: 8px;
      }

      input {
      width: 75px;
      height: 100px;
      font-size: 75px;
      border-radius: 5px;
      text-align: center;
    }
    }
  }
</style>
