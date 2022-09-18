<template>
  <div id="admin">
    <form @submit.prevent v-if="!reveal">
      <label v-for="(input, inputKey) in randomInputs" :key="inputKey">
        <span>{{ input.order }}</span>
        <input
          :id="`inputRefs-${inputKey}`"
          type="text"
          minlength="1"
          v-model="inputtedCode[inputKey]"
          @input="onInput(inputKey)" />
      </label>
    </form>

    <div class="json_wrapper" v-else>
      <h2>JSON file content</h2>
      <textarea id="json" v-html="JSON.stringify(Portfolio, null, 2)"></textarea>
      <button @click="saveJson">Save</button>
    </div>

    <div v-if="!!error.message || !!error.detail">
      <h3>{{ error.message }}</h3>
      <p>{{ error.detail }}</p>
    </div>
  </div>
</template>

<script>
import {
  onBeforeMount, onMounted, reactive, ref,
} from 'vue';

import ADMIN_CODE from '@/components/Utils/constants';

import Portfolio from '@/components/JSON/portfolio.json';

export default {
  name: 'Admin',
  setup() {
    const codes = ref(ADMIN_CODE);
    const randomInputs = ref([]);
    const inputtedCode = ref([]);

    const reveal = ref(false);
    const error = reactive({
      message: '',
      detail: '',
    });

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

    // Event Handlers
    const onInput = (idx) => {
      const input = document.querySelector(`#inputRefs-${idx + 1}`);
      if (!input) {
        for (let i = 0; i < randomInputs.value.length; i += 1) {
          if (randomInputs.value[i].code !== inputtedCode.value[i]) {
            error.message = 'Mismatch admin code';
            error.detail = `Potential mismatch at order ${randomInputs.value[i].order}`;
            break;
          }
          if (i === randomInputs.value.length - 1) {
            error.message = '';
            error.detail = '';
            reveal.value = true;
          }
        }
        return;
      }
      input.select();
    };

    const saveJson = () => {

    };

    // Life Cycles
    onBeforeMount(() => {
      randomPick(1);
    });

    onMounted(() => {
      const firstInput = document.querySelector('#inputRefs-0');
      firstInput.focus();
    });

    return {
      randomInputs,
      inputtedCode,
      error,
      reveal,

      Portfolio,

      onInput,
      saveJson,
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
    flex-direction: column;
    justify-content: center;
    align-items: center;
    gap: calc(8px * 3);
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

    .json_wrapper {
      position: relative;
      width: 100%;
      height: 100%;

      #json {
        position: relative;
        width: 100%;
        height: 80%;
        box-sizing: border-box;
        margin: calc(8px * 3) 0;
        padding: calc(8px * 3);
        color: rgb(226, 91, 91) ;
        background-color: rgba(0, 0, 0, .75);
        line-height: 120%;
        font-family: monospace, monospace;
        font-size: 100%;
        border: 0;
        overflow-y: scroll;
        overflow-x: hidden;
        outline: none;
        resize: none;
        word-wrap: break-word;
        white-space: pre-wrap;
        font-feature-settings: "liga" 0;
      }
    }
  }

  /* width of the entire scrollbar */
  ::-webkit-scrollbar,
  ::-webkit-scrollbar { /* Chrome/Opera/Safari */
      width: 4px;
  }
  ::-moz-scrollbar,
  ::-webkit-scrollbar { /* Firefox 19+ */
      width: 4px;
  }

  /* Track */
  ::-webkit-scrollbar-track {
    background-color: rgb(214, 105, 105);
  }

  /* Handle */
  ::-webkit-scrollbar-thumb {
    background-color: rgb(204, 46, 46);
    border-radius: 2px;
  }

  /* Handle on hover */
  ::-webkit-scrollbar-thumb:hover {
    background-color: rgb(204, 46, 46);;
  }
</style>
